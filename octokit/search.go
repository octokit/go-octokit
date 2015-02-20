package octokit

import (
	"github.com/jingweno/go-sawyer/hypermedia"
	"net/url"
)

var SearchURL = Hyperlink("search{/type}?q={query}&page={page}&per_page={per_page}&sort={sort}&order={order}")

func (c *Client) Search(uriTemplate *Hyperlink) (searches *SearchService) {
	searches = &SearchService{client: c, uriTemplate: uriTemplate}
	return
}

// A service to return search records
type SearchService struct {
	client      *Client
	uriTemplate *Hyperlink
}

// Get the user search results based on SearchService#URL
func (g *SearchService) Users(options M) (userSearchResults UserSearchResults,
	result *Result) {
	result = g.client.get(g.uriTemplate.Expand(m), &userSearchResults)
	return
}

// Get the issue search results based on SearchService#URL
func (g *SearchService) Issues(options M) (issueSearchResults IssueSearchResults,
	result *Result) {
	result = g.client.get(g.uriTemplate.Expand(m), &issueSearchResults)
	return
}

// Get the repository search results based on SearchService#URL
func (g *SearchService) Repositories(options M) (
	repositorySearchResults RepositorySearchResults, result *Result) {
	result = g.client.get(g.uriTemplate.Expand(m), &repositorySearchResults)
	return
}

// Get the code search results based on SearchService#URL
func (g *SearchService) Code(options M) (
	codeSearchResults CodeSearchResults, result *Result) {
	result = g.client.get(g.uriTemplate.Expand(m), &codeSearchResults)
	return
}

type UserSearchResults struct {
	*hypermedia.HALResource

	TotalCount        int    `json:"total_count,omitempty"`
	IncompleteResults bool   `json:"incomplete_results,omitempty"`
	Items             []User `json:"items,omitempty"`
}

type IssueSearchResults struct {
	*hypermedia.HALResource

	TotalCount        int     `json:"total_count,omitempty"`
	IncompleteResults bool    `json:"incomplete_results,omitempty"`
	Items             []Issue `json:"items,omitempty"`
}

type RepositorySearchResults struct {
	*hypermedia.HALResource

	TotalCount        int          `json:"total_count,omitempty"`
	IncompleteResults bool         `json:"incomplete_results,omitempty"`
	Items             []Repository `json:"items,omitempty"`
}

type CodeSearchResults struct {
	*hypermedia.HALResource

	TotalCount        int        `json:"total_count,omitempty"`
	IncompleteResults bool       `json:"incomplete_results,omitempty"`
	Items             []CodeFile `json:"items,omitempty"`
}

type CodeFile struct {
	*hypermedia.HALResource

	Name       string     `json:"name,omitempty"`
	Path       string     `json:"path,omitempty"`
	SHA        string     `json:"sha,omitempty"`
	URL        Hyperlink  `json:"url,omitempty"`
	GitURL     Hyperlink  `json:"git_url,omitempty"`
	HTMLURL    Hyperlink  `json:"html_url,omitempty"`
	Repository Repository `json:"repository,omitempty"`
}
