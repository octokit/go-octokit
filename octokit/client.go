package octokit

import (
	"io"
	"net/http"
	"net/url"

	"github.com/lostisland/go-sawyer"
	"github.com/lostisland/go-sawyer/hypermedia"
)

func NewClient(authMethod AuthMethod) *Client {
	return NewClientWith(gitHubAPIURL, userAgent, authMethod, nil)
}

func NewClientWith(baseURL string, userAgent string, authMethod AuthMethod, httpClient *http.Client) *Client {
	client, _ := sawyer.NewFromString(baseURL, httpClient)
	return &Client{sawyerClient: client, UserAgent: userAgent, AuthMethod: authMethod}
}

type Client struct {
	UserAgent    string
	AuthMethod   AuthMethod
	sawyerClient *sawyer.Client
	rootRels     hypermedia.Relations
	middlewares  []Middleware
}

func (c *Client) Use(m Middleware) {
	c.middlewares = append(c.middlewares, m)
}

func (c *Client) NewRequest(urlStr string) (req *Request, err error) {
	req, err = newRequest(c, urlStr)
	if err != nil {
		return
	}

	c.applyRequestHeaders(req)

	return
}

func (c *Client) head(url *url.URL, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Head(output)
	})
}

func (c *Client) get(url *url.URL, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Get(output)
	})
}

func (c *Client) post(url *url.URL, input interface{}, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Post(input, output)
	})
}

func (c *Client) put(url *url.URL, input interface{}, output interface{}) *Result {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Put(input, output)
	})
}

func (c *Client) delete(url *url.URL, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Delete(output)
	})
}

func (c *Client) patch(url *url.URL, input interface{}, output interface{}) (result *Result) {
	return sendRequest(c, url, func(req *Request) (*Response, error) {
		return req.Patch(input, output)
	})
}

func (c *Client) upload(uploadUrl *url.URL, asset io.ReadCloser, contentType string, contentLength int64) (result *Result) {
	req, err := c.NewRequest(uploadUrl.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	req.Header.Set("Content-Type", contentType)
	req.ContentLength = contentLength

	req.Body = asset
	sawyerResp := req.Request.Post()

	resp, err := NewResponse(sawyerResp)
	return newResult(resp, err)
}

func (c *Client) applyRequestHeaders(req *Request) {
	req.Header.Set("Accept", defaultMediaType)
	req.Header.Set("User-Agent", c.UserAgent)

	if c.AuthMethod != nil {
		req.Header.Set("Authorization", c.AuthMethod.String())
	}

	if basicAuth, ok := c.AuthMethod.(BasicAuth); ok && basicAuth.OneTimePassword != "" {
		req.Header.Set("X-GitHub-OTP", basicAuth.OneTimePassword)
	}

	return
}

func sendRequest(c *Client, url *url.URL, fn func(r *Request) (*Response, error)) (result *Result) {
	req, err := c.NewRequest(url.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	resp, err := fn(req)
	result = newResult(resp, err)

	return
}

func (c *Client) applyRequestMiddlewares(req *Request) error {
	for _, m := range c.middlewares {
		if err := m.PrepareRequest(req); err != nil {
			return err
		}
	}

	return nil
}

func (c *Client) applyResponseMiddlewares(resp *Response) error {
	for _, m := range c.middlewares {
		if err := m.PrepareResponse(resp); err != nil {
			return err
		}
	}

	return nil
}
