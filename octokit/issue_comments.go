package octokit

import (
	"time"

	"github.com/jingweno/go-sawyer/hypermedia"
)

var IssueCommentsURL = Hyperlink("/repos/{owner}/{repo}/issues{/number}/comments{/id}")

// Create a IssueCommentsService
func (c *Client) IssueComments() (k *IssueCommentsService) {
	k = &IssueCommentsService{client: c}
	return
}

// A service to return comments for issues
type IssueCommentsService struct {
	client *Client
}

// Get a list of all issue comments
func (c *IssueCommentsService) All(uri *Hyperlink, params M) (comments []IssueComment, result *Result) {
	if uri == nil {
		uri = &IssueCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = c.client.get(url, &comments)
	return
}

// Get a single comment by id
func (c *IssueCommentsService) One(uri *Hyperlink, params M) (comment *IssueComment, result *Result) {
	if uri == nil {
		uri = &IssueCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = c.client.get(url, &comment)
	return
}

// Creates a comment on an issue
func (c *IssueCommentsService) Create(uri *Hyperlink, params M, input interface{}) (comment *IssueComment, result *Result) {
	if uri == nil {
		uri = &IssueCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = c.client.post(url, input, &comment)
	return
}

// Updates a comment on an issue
func (c *IssueCommentsService) Update(uri *Hyperlink, params M, input interface{}) (comment *IssueComment, result *Result) {
	if uri == nil {
		uri = &IssueCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = c.client.patch(url, input, &comment)
	return
}

// Deletes a comment on an issue
func (c *IssueCommentsService) Delete(uri *Hyperlink, params M) (success bool, result *Result) {
	if uri == nil {
		uri = &IssueCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return false, &Result{Err: err}
	}

	result = c.client.delete(url, nil, nil)
	success = (result.Response.StatusCode == 204)
	return
}

type IssueComment struct {
	*hypermedia.HALResource

	ID        int        `json:"id,omitempty"`
	URL       string     `json:"url,omitempty"`
	User      User       `json:"user,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Body      string     `json:"body,omitempty"`
	HTMLURL   string     `json:"html_url,omitempty"`
	IssueURL  string     `json:"issue_url",omitempty"`
}
