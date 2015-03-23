package octokit

import (
	"github.com/jingweno/go-sawyer/hypermedia"
)

var (
	CodeSearchURL       = Hyperlink("/search/code?q={query}{&page,per_page,sort,order}")
	IssueSearchURL      = Hyperlink("/search/issues?q={query}{&page,per_page,sort,order}")
	RepositorySearchURL = Hyperlink("/search/repositories?q={query}{&page,per_page,sort,order}")
	UserSearchURL       = Hyperlink("/search/users?q={query}{&page,per_page,sort,order}")
)

func (c *Client) Search() *SearchService {
	return &SearchService{client: c}
}

// A service to return search records
type SearchService struct {
	client *Client
}

// Get the user search results based on SearchService#URL
func (g *SearchService) Users(uri *Hyperlink, params M) (
	userSearchResults UserSearchResults, result *Result) {
	if uri == nil {
		uri = &UserSearchURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return UserSearchResults{}, &Result{Err: err}
	}
	result = g.client.get(url, &userSearchResults)
	return
}

// Get the issue search results based on SearchService#URL
func (g *SearchService) Issues(uri *Hyperlink, params M) (
	issueSearchResults IssueSearchResults, result *Result) {
	if uri == nil {
		uri = &IssueSearchURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return IssueSearchResults{}, &Result{Err: err}
	}
	result = g.client.get(url, &issueSearchResults)
	return
}

// Get the repository search results based on SearchService#URL
func (g *SearchService) Repositories(uri *Hyperlink, params M) (
	repositorySearchResults RepositorySearchResults, result *Result) {
	if uri == nil {
		uri = &RepositorySearchURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return RepositorySearchResults{}, &Result{Err: err}
	}
	result = g.client.get(url, &repositorySearchResults)
	return
}

// Get the code search results based on SearchService#URL
func (g *SearchService) Code(uri *Hyperlink, params M) (
	codeSearchResults CodeSearchResults, result *Result) {
	if uri == nil {
		uri = &CodeSearchURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return CodeSearchResults{}, &Result{Err: err}
	}
	result = g.client.get(url, &codeSearchResults)
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
