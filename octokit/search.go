package octokit

import (
	"github.com/jingweno/go-sawyer/hypermedia"
	"net/url"
)

var SearchURL = Hyperlink("search{/type}?q={query}{&page,per_page,sort,order}")

func (c *Client) Searches(url *url.URL) (searches *SearchService) {
	searches = &SearchService{client: c, URL: url}
	return
}

// A service to return search records
type SearchService struct {
	client *Client
	URL    *url.URL
}

// Get a list of search results based on SearchService#URL
func (g *SearchService) UserSearch() (userSearchResults UserSearchResults,
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
