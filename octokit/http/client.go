package http

import (
	"github.com/lostisland/go-sawyer"
	"net/http"
)

func NewClient(baseURL string, httpClient *http.Client) *Client {
	client, _ := sawyer.NewFromString(baseURL, httpClient)
	return &Client{client}
}

type Client struct {
	sawyerClient *sawyer.Client
}

func (c *Client) NewRequest(urlStr string) (req *Request, err error) {
	sawyerReq, err := c.sawyerClient.NewRequest(urlStr)
	if err != nil {
		return
	}

	req = &Request{sawyerReq: sawyerReq}
	return
}
