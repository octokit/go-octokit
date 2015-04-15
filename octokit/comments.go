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
