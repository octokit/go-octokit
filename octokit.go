package octokit

import (
	"net/http"
)

func NewClient() *Client {
	return &Client{&http.Client{}, "", "", ""}
}
