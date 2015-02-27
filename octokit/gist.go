package octokit

import (
	"io"
	"net/url"
	"time"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// GistsURL is a template for accessing gists from GitHub possibly with
// a specific identification code that can be expanded to a full address.
var GistsURL = Hyperlink("gists{/gist_id}")

// Gists creates a GistsService with a base url
func (c *Client) Gists(url *url.URL) (gists *GistsService) {
	gists = &GistsService{client: c, URL: url}
	return
}

// GistsService is a service providing access to gists from a particular url
type GistsService struct {
	client *Client
	URL    *url.URL
}

// One gets a specific gist based on the url of the service
func (g *GistsService) One() (gist *Gist, result *Result) {
	result = g.client.get(g.URL, &gist)
	return
}

// Update modifies a specific gist based on the url of the service
func (g *GistsService) Update(params interface{}) (gist *Gist, result *Result) {
	result = g.client.put(g.URL, params, &gist)
	return
}

// All gets a list of all gists associated with the url of the service
func (g *GistsService) All() (gists []Gist, result *Result) {
	result = g.client.get(g.URL, &gists)
	return
}

// Raw gets the raw contents of first file in a specific gist
func (g *GistsService) Raw() (body io.ReadCloser, result *Result) {
	var gist *Gist
	var rawURL *url.URL

	gist, result = g.One()
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
