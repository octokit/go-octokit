package octokat

import (
	"fmt"
)

type PullRequestParams struct {
	Base  string `json:"base"`
	Head  string `json:"head"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PullRequestForIssueParams struct {
	Base  string `json:"base"`
	Head  string `json:"head"`
	Issue string `json:"issue"`
}

type PullRequest struct {
	URL      string `json:"url"`
	HTMLURL  string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
	IssueURL string `json:"issue_url"`
}

func (c *Client) CreatePullRequest(repo Repo, params PullRequestParams) (*PullRequest, error) {
	return c.createPullRequest(repo, params)
}

func (c *Client) CreatePullRequestForIssue(repo Repo, params PullRequestForIssueParams) (*PullRequest, error) {
	return c.createPullRequest(repo, params)
}

func (c *Client) createPullRequest(repo Repo, params interface{}) (*PullRequest, error) {
	path := fmt.Sprintf("repos/%s/pulls", repo)
	var pr PullRequest
	err := c.jsonPost(path, nil, params, &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}
