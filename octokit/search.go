package octokit

import (
	"github.com/jingweno/go-sawyer/hypermedia"
)

var SearchURITemplate = "search{/type}?q={query}{&page,per_page,sort,order}"

func (c *Client) Search(uriTemplate string) *SearchService {
	return &SearchService{&GenericService{client: c, uriTemplate: uriTemplate}}
}

// A service to return search records
type SearchService struct {
	*GenericService
}

// Get the user search results based on SearchService#URL
func (g *SearchService) Users(params M) (userSearchResults UserSearchResults,
	result *Result) {
	url, e := g.getURL(params, M{"type": "users"})
	if e != nil {
		return UserSearchResults{}, &Result{Err: e}
	}
	result = g.GenericService.client.get(url, &userSearchResults)
	return
}

// Get the issue search results based on SearchService#URL
func (g *SearchService) Issues(params M) (issueSearchResults IssueSearchResults,
	result *Result) {
	url, e := g.getURL(params, M{"type": "issues"})
	if e != nil {
		return IssueSearchResults{}, &Result{Err: e}
	}
	result = g.GenericService.client.get(url, &issueSearchResults)
	return
}

// Get the repository search results based on SearchService#URL
func (g *SearchService) Repositories(params M) (
	repositorySearchResults RepositorySearchResults, result *Result) {
	url, e := g.getURL(params, M{"type": "repositories"})
	if e != nil {
		return RepositorySearchResults{}, &Result{Err: e}
	}
	result = g.GenericService.client.get(url, &repositorySearchResults)
	return
}

// Get the code search results based on SearchService#URL
func (g *SearchService) Code(params M) (codeSearchResults CodeSearchResults,
	result *Result) {
	url, e := g.getURL(params, M{"type": "code"})
	if e != nil {
		return CodeSearchResults{}, &Result{Err: e}
	}
	result = g.GenericService.client.get(url, &codeSearchResults)
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
