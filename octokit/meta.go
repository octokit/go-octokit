package octokit

import (
	"net/url"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// https://developer.github.com/v3/meta/
var (
	MetaURL = Hyperlink("/meta")
)

// https://developer.github.com/v3/meta/
func (c *Client) Meta(url *url.URL) *MetaService {
	return &MetaService{client: c, URL: url}
}

// A service to return information about GitHub.com, the service.
type MetaService struct {
	client *Client
	URL    *url.URL
}

// Get the user search results based on MetaService#URL
//
// https://developer.github.com/v3/meta/#meta
func (m *MetaService) One() (meta Meta, result *Result) {
	result = m.client.get(m.URL, &meta)
	return
}

type Meta struct {
	*hypermedia.HALResource

	VerifiablePasswordAuthentication bool     `json:"verifiable_password_authentication,omitempty"`
	GithubServicesSha                string   `json:"github_services_sha,omitempty"`
	Hooks                            []string `json:"hooks,omitempty"`
	Git                              []string `json:"git,omitempty"`
	Pages                            []string `json:"pages,omitempty"`
}
