package octokit

import (
	"github.com/jingweno/go-sawyer/hypermedia"
	"net/url"
)

var SearchURL = Hyperlink("search{/type}?q={query}{&page,per_page,sort,order}")

func (c *Client) Search(url *url.URL) (searches *SearchService) {
	searches = &SearchService{client: c, URL: url}
	return
}

// A service to return search records
type SearchService struct {
	client *Client
	URL    *url.URL
}

// Get the user search results based on SearchService#URL
func (g *SearchService) Users() (userSearchResults UserSearchResults,
	result *Result) {
	result = g.client.get(g.URL, &userSearchResults)
	return
}

// Get the issue search results based on SearchService#URL
func (g *SearchService) Issues() (issueSearchResults IssueSearchResults,
	result *Result) {
	result = g.client.get(g.URL, &issueSearchResults)
	return
}

// Get the repository search results based on SearchService#URL
func (g *SearchService) Repositories() (
	repositorySearchResults RepositorySearchResults, result *Result) {
	result = g.client.get(g.URL, &repositorySearchResults)
	return
}

// Get the code search results based on SearchService#URL
func (g *SearchService) Code() (
	codeSearchResults CodeSearchResults, result *Result) {
	result = g.client.get(g.URL, &codeSearchResults)
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
