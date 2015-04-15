package octokit

import (
	"io"
	"net/url"
	"time"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// GistsURL is a template for accessing gists from GitHub possibly with
// a specific identification code that can be expanded to a full address.
//
// https://developer.github.com/v3/gists
var GistsURL = Hyperlink("gists{/gist_id}")

// Gists creates a GistsService to be used with any proper URL
//
// https://developer.github.com/v3/gists/
func (c *Client) Gists() (gists *GistsService) {
	gists = &GistsService{client: c}
	return
}

// GistsService is a service providing access to gists from a particular url
type GistsService struct {
	client *Client
}

// One gets a specific gist based on the url of the service
//
// https://developer.github.com/v3/gists/#get-a-single-gist
func (g *GistsService) One(uri *Hyperlink, params M) (gist Gist, result *Result) {
	if uri == nil {
		uri = &GistsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return Gist{}, &Result{Err: err}
	}
	result = g.client.get(url, &gist)
	return
}

// Update modifies a specific gist based on the url of the service
//
// https://developer.github.com/v3/gists/#edit-a-gist
func (g *GistsService) Update(uri *Hyperlink, params M, edits interface{}) (gist Gist, result *Result) {
	if uri == nil {
		uri = &GistsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return Gist{}, &Result{Err: err}
	}
	result = g.client.patch(url, params, &gist)
	return
}

// All gets a list of all gists associated with the url of the service
//
// https://developer.github.com/v3/gists/#list-gists
func (g *GistsService) All(uri *Hyperlink, params M) (gists []Gist, result *Result) {
	if uri == nil {
		uri = &GistsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return make([]Gist, 0), &Result{Err: err}
	}
	result = g.client.get(url, &gists)
	return
}

// Raw gets the raw contents of first file in a specific gist
//
// https://developer.github.com/v3/gists/#get-a-single-gist
func (g *GistsService) Raw(uri *Hyperlink, params M) (body io.ReadCloser, result *Result) {
	var gist Gist
	var rawURL *url.URL

	gist, result = g.One(uri, params)
	for _, file := range gist.Files {
		rawURL, _ = url.Parse(file.RawURL)
		break
	}

	body, result = g.client.getBody(rawURL, textMediaType)
	return
}

// GistFile is a representation of the file stored in a gist
type GistFile struct {
	*hypermedia.HALResource

	FileName  string `json:"filename,omitempty"`
	Type      string `json:"type,omitempty"`
	Language  string `json:"language,omitempty"`
	RawURL    string `json:"raw_url,omitempty"`
	Size      int    `json:"size,omitempty"`
	Truncated bool   `json:"truncated,omitempty"`
	Content   string `json:"content,omitempty"`
}

// Gist is a representation of a gist on github, a standalone file that acts as a
// sole element of its own repository
type Gist struct {
	*hypermedia.HALResource

	ID          string               `json:"id,omitempty"`
	Comments    float64              `json:"comments,omitempty"`
	CommentsURL string               `json:"comments_url,omitempty"`
	CommitsURL  string               `json:"commits_url,omitempty"`
	CreatedAt   string               `json:"created_at,omitempty"`
	Description string               `json:"description,omitempty"`
	Files       map[string]*GistFile `json:"files,omitempty"`
	ForksURL    Hyperlink            `json:"forks_url,omitempty"`
	GitPullURL  Hyperlink            `json:"git_pull_url,omitempty"`
	GitPushURL  Hyperlink            `json:"git_push_url,omitempty"`
	HtmlURL     Hyperlink            `json:"html_url,omitempty"`
	Owner       *User                `json:"owner,omitempty"`
	Public      bool                 `json:"public,omitempty"`
	UpdatedAt   *time.Time           `json:"updated_at,omitempty"`
	URL         string               `json:"url,omitempty"`
	User        *User                `json:"user,omitempty"`
}
