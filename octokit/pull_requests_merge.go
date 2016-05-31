package octokit

import (
	"net/url"
)

var PullRequestsMergeURL = Hyperlink("repos/{owner}/{repo}/pulls/{number}/merge")

func (c *Client) PullRequestsMerge(url *url.URL) (pullRequests *PullRequestsMergeService) {
	pullRequests = &PullRequestsMergeService{client: c, URL: url}
	return
}

type PullRequestsMergeService struct {
	client *Client
	URL    *url.URL
}

func (p *PullRequestsMergeService) Merge(input *PullRequestsMergeRequest) (*PullRequestsMergeResponse, *Result) {
	var resp PullRequestsMergeResponse
	result := p.client.put(p.URL, input, &resp)
	return &resp, result
}

type PullRequestsMergeRequest struct {
	CommitMessage string `json:"commit_message,omitempty"`
	Sha           string `json:"sha,omitempty"`
}

type PullRequestsMergeResponse struct {
	Sha     string `json:"sha,omitempty"`
	Merged  bool   `json:"merged,omitempty"`
	Message string `json:"message,omitempty"`
}
