package http

import (
	"github.com/lostisland/go-sawyer"
	"net/http"
	"reflect"
)

func NewClient(httpClient *http.Client) *Client {
	client, _ := sawyer.NewFromString("https://api.github.com", httpClient)
	client.ErrorType = reflect.TypeOf(ResponseError{})
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
