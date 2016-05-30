package octokit

import (
	"net/url"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// https://developer.github.com/v3/repos/statuses/#get-the-combined-status-for-a-specific-ref
var StatusURL = Hyperlink("repos/{owner}/{repo}/status/{ref}")

func (c *Client) Status(url *url.URL) (statuses *StatusService) {
	statuses = &StatusService{client: c, URL: url}
	return
}

type StatusService struct {
	client *Client
	URL    *url.URL
}

func (s *StatusService) Get() (status *StatusCombined, result *Result) {
	result = s.client.get(s.URL, &status)
	return
}

type StatusCombined struct {
	*hypermedia.HALResource

	State      string `json:"state,omitempty"`
	TotalCount int    `json:"total_count,omitempty"`
	Sha        string `json:"sha,omitempty"`
}
