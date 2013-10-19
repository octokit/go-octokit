package octokit

import (
	"net/http"
)

const (
	GitHubURL          = "https://github.com"
	GitHubAPIURL       = "https://api.github.com"
	UserAgent          = "Octokit Go " + Version
	MediaType          = "application/vnd.github.beta+json"
	DefaultContentType = "application/json"
	Version            = "0.3.0"
)

func NewClient() *Client {
	return &Client{BaseURL: GitHubAPIURL, httpClient: &http.Client{}}
}
