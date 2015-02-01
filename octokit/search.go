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
func (g *SearchService) UserSearch() (userSearchResults UserSearchResults,
	result *Result) {
	result = g.client.get(g.URL, &userSearchResults)
	return
}

// Get the issue search results based on SearchService#URL
func (g *SearchService) IssueSearch() (issueSearchResults IssueSearchResults,
	result *Result) {
	result = g.client.get(g.URL, &issueSearchResults)
	return
}

// Get the repository search results based on SearchService#URL
func (g *SearchService) RepositorySearch() (
	repositorySearchResults RepositorySearchResults, result *Result) {
	result = g.client.get(g.URL, &repositorySearchResults)
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
