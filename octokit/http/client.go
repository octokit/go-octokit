package http

import (
	"github.com/lostisland/go-sawyer"
	"net/http"
)

func NewClient(httpClient *http.Client) *Client {
	client, _ := sawyer.NewFromString("https://api.github.com", httpClient)
	return &Client{client}
}

type Client struct {
	sawyerClient *sawyer.Client
}

func (c *Client) NewRequest(urlStr string) *Request {
	return &Request{client: c, URL: urlStr}
}
