package octokit

import (
	"github.com/jingweno/go-sawyer/hypermedia"
	"net/url"
)

var SearchURL = Hyperlink("search{/type}")

func (c *Client) UserSearches(url *url.URL) (searches *UserSearchService) {
	searches = &UserSearchService{client: c, URL: url}
	return
}

// A service to return search records
type UserSearchService struct {
	client *Client
	URL    *url.URL
}

// Get a list of search results based on SearchService#URL
func (g *UserSearchService) All() (userSearchResults UserSearchResults,
	result *Result) {
	result = g.client.get(g.URL, &userSearchResults)
	return
}

type UserSearchResults struct {
	*hypermedia.HALResource

	TotalCount        int    `json:"total_count,omitempty"`
	IncompleteResults bool   `json:"incomplete_results,omitempty"`
	Items             []User `json:"items,omitempty"`
}
