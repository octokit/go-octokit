package octokit

import (
	"net/url"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// TagsURL is a template for accessing tags in a particular repository
// for a particular owner that can be expanded to a full address.
//
// https://developer.github.com/v3/git/tags/
var (
	TagsURL = Hyperlink("repos/{owner}/{repo}/git/refs/tags")
)

// Release is a representation of a release on GitHub. Published releases are
// available to everyone.
type Tag struct {
	*hypermedia.HALResource

	Ref    string `json:"ref,omitempty"`
	URL    string `json:"url,omitempty"`
	Commit Commit `json:"object,omitempty"`
}

// Tags creates a TagsService with a base url
//
// https://developer.github.com/v3/git/tags/
func (c *Client) Tags(url *url.URL) (tags *TagsService) {
	tags = &TagsService{client: c, URL: url}
	return
}

// TagsService is a service providing access to tags from a particular url
type TagsService struct {
	client *Client
	URL    *url.URL
}

// All gets all tags for a given repository based on the URL of the service
//
// https://developer.github.com/v3/git/tags/#get-a-tag
func (t *TagsService) All() (tags []Tag, result *Result) {
	result = t.client.get(t.URL, &tags)
	return
}
