package octokat

import (
	"net/http"
)

const (
	GitHubURL    string = "https://github.com"
	GitHubAPIURL string = "https://api.github.com"
)

func NewClient() *Client {
	return &Client{baseURL: GitHubAPIURL, httpClient: &http.Client{}}
}
