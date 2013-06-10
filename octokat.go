package octokat

import (
	"net/http"
)

func NewClient() *Client {
	return &Client{&http.Client{}, "", "", ""}
}
