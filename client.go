package octokat

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	httpClient *http.Client
	BaseURL    string
	Login      string
	Password   string
	Token      string
}

func (c *Client) WithLogin(login, password string) *Client {
	c.Login = login
	c.Password = password
	return c
}

func (c *Client) WithToken(token string) *Client {
	c.Token = token
	return c
}

func (c *Client) get(path string, options *Options) ([]byte, error) {
	return c.request("GET", path, options, nil)
}

func (c *Client) post(path string, options *Options, content io.Reader) ([]byte, error) {
	return c.request("POST", path, options, content)
}

func (c *Client) jsonGet(path string, options *Options, v interface{}) error {
	body, err := c.get(path, options)
	if err != nil {
		return err
	}

	return jsonUnmarshal(body, v)
}

func (c *Client) jsonPost(path string, options *Options, v interface{}) error {
	var buffer *bytes.Buffer
	if options.Params != nil {
		b, err := jsonMarshal(options.Params)
		if err != nil {
			return err
		}

		buffer = bytes.NewBuffer(b)
	}

	body, err := c.post(path, options, buffer)
	if err != nil {
		return err
	}

	return jsonUnmarshal(body, v)
}

func (c *Client) request(method, path string, options *Options, content io.Reader) ([]byte, error) {
	url := concatPath(c.BaseURL, path)
	request, err := http.NewRequest(method, url, content)
	if err != nil {
		return nil, err
	}

	c.setDefaultHeaders(request)

	if options != nil {
		for h, v := range options.Headers {
			request.Header.Set(h, v)
		}
	}

	response, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 && response.StatusCode < 600 {
		return nil, handleErrors(body)
	}

	return body, nil
}

func (c *Client) setDefaultHeaders(request *http.Request) {
	request.Header.Set("Accept", "application/vnd.github.beta+json")
	request.Header.Set("Content-Type", "application/json")
	if c.Token != "" {
		request.Header.Set("Authorization", fmt.Sprintf("token %s", c.Token))
	}
	if c.Login != "" && c.Password != "" {
		request.Header.Set("Authorization", fmt.Sprintf("Basic %s", hashAuth(c.Login, c.Password)))
	}
}
