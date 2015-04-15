package octokit

import (
	"time"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// RepoCommentsURL is a template for comments linked to a specific repository
// CommitCommentsURL is a template for comments linked to a specific commit
var (
	RepoCommentsURL   = Hyperlink("/repos/{owner}/{repo}/comments{/id}")
	CommitCommentsURL = Hyperlink("/repos/{owner}/{repo}/commits/{sha}/comments")
)

// Create a CommentsService
func (c *Client) Comments() (k *CommentsService) {
	k = &CommentsService{client: c}
	return
}

// A service to return comments
type CommentsService struct {
	client *Client
}

// Get a list of all comments
func (c *CommentsService) All(uri *Hyperlink, params M) (comments []Comment, result *Result) {
	if uri == nil {
		uri = &RepoCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = c.client.get(url, &comments)
	return
}

// Get a single comment by id
func (c *CommentsService) One(uri *Hyperlink, params M) (comment *Comment, result *Result) {
	if uri == nil {
		uri = &RepoCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = c.client.get(url, &comment)
	return
}

// Creates a comment on a commit
func (c *CommentsService) Create(uri *Hyperlink, params M, input interface{}) (comment *Comment, result *Result) {
	if uri == nil {
		uri = &CommitCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = c.client.post(url, input, &comment)
	return
}

// Updates a comment on a commit
func (c *CommentsService) Update(uri *Hyperlink, params M, input interface{}) (comment *Comment, result *Result) {
	if uri == nil {
		uri = &RepoCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}

	result = c.client.patch(url, input, &comment)
	return
}

// Deletes a comment on a commit
func (c *CommentsService) Delete(uri *Hyperlink, params M) (success bool, result *Result) {
	if uri == nil {
		uri = &RepoCommentsURL
	}

	url, err := uri.Expand(params)
	if err != nil {
		return false, &Result{Err: err}
	}

	result = c.client.delete(url, nil, nil)
	success = (result.Response.StatusCode == 204)
	return
}

type Comment struct {
	*hypermedia.HALResource

	ID        int        `json:"id,omitempty"`
	URL       string     `json:"url,omitempty"`
	HTMLURL   string     `json:"html_url,omitempty"`
	User      User       `json:"user,omitempty"`
	Position  int        `json:"position,omitempty"`
	Line      int        `json:"line,omitempty"`
	Path      string     `json:"path,omitempty"`
	CommitID  string     `json:"commit_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	Body      string     `json:"body,omitempty"`
}
