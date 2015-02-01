package octokit

import (
	"net/url"
	"time"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// RepoIssuesURL is a template for accessing issues in a particular
// repository for a particular owner that can be expanded to a full address.
var RepoIssuesURL = Hyperlink("repos/{owner}/{repo}/issues{/number}")

// Issues creates an IssuesService with a base url
func (c *Client) Issues(url *url.URL) (issues *IssuesService) {
	issues = &IssuesService{client: c, URL: url}
	return
}

// IssuesService is a service providing access to issues from a particular url
type IssuesService struct {
	client *Client
	URL    *url.URL
}

// One gets a specific issue based on the url of the service
func (i *IssuesService) One() (issue *Issue, result *Result) {
	result = i.client.get(i.URL, &issue)
	return
}

// All gets a list of all issues associated with the url of the service
func (i *IssuesService) All() (issues []Issue, result *Result) {
	result = i.client.get(i.URL, &issues)
	return
}

// Create posts a new issue with particular parameters to the issues service url
func (i *IssuesService) Create(params interface{}) (issue *Issue, result *Result) {
	result = i.client.post(i.URL, params, &issue)
	return
}

// Update modifies a specific issue given parameters on the service url
func (i *IssuesService) Update(params interface{}) (issue *Issue, result *Result) {
	result = i.client.patch(i.URL, params, &issue)
	return
}

// Issue represents an issue on GitHub with all associated fields
type Issue struct {
	*hypermedia.HALResource

	URL     string `json:"url,omitempty,omitempty"`
	HTMLURL string `json:"html_url,omitempty,omitempty"`
	Number  int    `json:"number,omitempty"`
	State   string `json:"state,omitempty"`
	Title   string `json:"title,omitempty"`
	Body    string `json:"body,omitempty"`
	User    User   `json:"user,omitempty"`
	Labels  []struct {
		URL   string `json:"url,omitempty"`
		Name  string `json:"name,omitempty"`
		Color string `json:"color,omitempty"`
	}
	Assignee  User `json:"assignee,omitempty"`
	Milestone struct {
		URL          string     `json:"url,omitempty"`
		Number       int        `json:"number,omitempty"`
		State        string     `json:"state,omitempty"`
		Title        string     `json:"title,omitempty"`
		Description  string     `json:"description,omitempty"`
		Creator      User       `json:"creator,omitempty"`
		OpenIssues   int        `json:"open_issues,omitempty"`
		ClosedIssues int        `json:"closed_issues,omitempty"`
		CreatedAt    time.Time  `json:"created_at,omitempty"`
		DueOn        *time.Time `json:"due_on,omitempty"`
	}
	Comments    int `json:"comments,omitempty"`
	PullRequest struct {
		HTMLURL  string `json:"html_url,omitempty"`
		DiffURL  string `json:"diff_url,omitempty"`
		PatchURL string `json:"patch_url,omitempty"`
	} `json:"pull_request,omitempty"`
	CreatedAt time.Time  `json:"created_at,omitempty"`
	ClosedAt  *time.Time `json:"closed_at,omitempty"`
	UpdatedAt time.Time  `json:"updated_at,omitempty"`
}

// IssueParams represents the struture used to create or update an Issue
type IssueParams struct {
	Title     string   `json:"title,omitempty"`
	Body      string   `json:"body,omitempty"`
	Assignee  string   `json:"assignee,omitempty"`
	State     string   `json:"state,omitempty"`
	Milestone uint64   `json:"milestone,omitempty"`
	Labels    []string `json:"labels,omitempty"`
}
