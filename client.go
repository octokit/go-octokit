package octokit

import (
	"github.com/lostisland/go-sawyer"
	"net/http"
	"net/url"
)

func NewClient() *Client {
	return NewClientWith(gitHubAPIURL, nil)
}

func NewClientWith(baseURL string, httpClient *http.Client) *Client {
	client, _ := sawyer.NewFromString(baseURL, httpClient)
	return &Client{sawyerClient: client, UserAgent: userAgent}
}

type Client struct {
	UserAgent    string
	sawyerClient *sawyer.Client
}

func (c *Client) NewRequest(urlStr string) (req *Request, err error) {
	sawyerReq, err := c.sawyerClient.NewRequest(urlStr)
	if err != nil {
		return
	}

	sawyerReq.Header.Add("User-Agent", c.UserAgent)

	req = &Request{sawyerReq: sawyerReq}
	return
}

func (c *Client) Head(url *url.URL, output interface{}) (result *Result) {
	req, err := c.NewRequest(url.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	resp, err := req.Head(output)
	result = newResult(resp, err)

	return
}

func (c *Client) Get(url *url.URL, output interface{}) (result *Result) {
	req, err := c.NewRequest(url.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	resp, err := req.Get(output)
	result = newResult(resp, err)

	return
}

func (c *Client) Post(url *url.URL, input interface{}, output interface{}) (result *Result) {
	req, err := c.NewRequest(url.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	resp, err := req.Post(input, output)
	result = newResult(resp, err)

	return
}

func (c *Client) Put(url *url.URL, input interface{}, output interface{}) (result *Result) {
	req, err := c.NewRequest(url.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	resp, err := req.Put(input, output)
	result = newResult(resp, err)

	return
}

func (c *Client) Delete(url *url.URL, output interface{}) (result *Result) {
	req, err := c.NewRequest(url.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	resp, err := req.Delete(output)
	result = newResult(resp, err)

	return
}
