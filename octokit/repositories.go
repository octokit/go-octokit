package octokit

import (
	"net/url"
	"time"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// Hyperlinks to the various repository locations on github.
// RepositoryURL is a template for a particular repository for a particular owner.
// ForksURL is a template for the forks of a user's repository.
// UserRepositoriesURL is the address for all user repositories.
// OrgRepositoriesUrl is the template for repositories within a particular organization.
var (
	RepositoryURL       = Hyperlink("repos/{owner}/{repo}")
	ForksURL            = Hyperlink("repos/{owner}/{repo}/forks")
	UserRepositoriesURL = Hyperlink("user/repos")
	OrgRepositoriesURL  = Hyperlink("orgs/{org}/repos")
)

// Repositories creates a RepositoriesService with a base url
func (c *Client) Repositories(url *url.URL) (repos *RepositoriesService) {
	repos = &RepositoriesService{client: c, URL: url}
	return
}

// RepositoriesService is a service providing access to repositories from a particular url
type RepositoriesService struct {
	client *Client
	URL    *url.URL
}

// One gets a specific repository based on the url of the service
func (r *RepositoriesService) One() (repo *Repository, result *Result) {
	result = r.client.get(r.URL, &repo)
	return
}

// All gets a list of all repositories associated with the url of the service
func (r *RepositoriesService) All() (repos []Repository, result *Result) {
	result = r.client.get(r.URL, &repos)
	return
}

// Create posts a new repository based on parameters in a Repository struct to
// the respository service url
func (r *RepositoriesService) Create(params interface{}) (repo *Repository, result *Result) {
	result = r.client.post(r.URL, params, &repo)
	return
}

// Repository represents a respository on GitHub with all associated metadata with
// respect to the particular accessing url
type Repository struct {
	*hypermedia.HALResource

	ID              int           `json:"id,omitempty"`
	Owner           User          `json:"owner,omitempty"`
	Name            string        `json:"name,omitempty"`
	FullName        string        `json:"full_name,omitempty"`
	Description     string        `json:"description,omitempty"`
	Private         bool          `json:"private"`
	Fork            bool          `json:"fork,omitempty"`
	URL             string        `json:"url,omitempty"`
	HTMLURL         string        `json:"html_url,omitempty"`
	CloneURL        string        `json:"clone_url,omitempty"`
	GitURL          string        `json:"git_url,omitempty"`
	SSHURL          string        `json:"ssh_url,omitempty"`
	SVNURL          string        `json:"svn_url,omitempty"`
	MirrorURL       string        `json:"mirror_url,omitempty"`
	Homepage        string        `json:"homepage,omitempty"`
	Language        string        `json:"language,omitempty"`
	Forks           int           `json:"forks,omitempty"`
	ForksCount      int           `json:"forks_count,omitempty"`
	StargazersCount int           `json:"stargazers_count,omitempty"`
	Watchers        int           `json:"watchers,omitempty"`
	WatchersCount   int           `json:"watchers_count,omitempty"`
	Size            int           `json:"size,omitempty"`
	MasterBranch    string        `json:"master_branch,omitempty"`
	OpenIssues      int           `json:"open_issues,omitempty"`
	PushedAt        *time.Time    `json:"pushed_at,omitempty"`
	CreatedAt       *time.Time    `json:"created_at,omitempty"`
	UpdatedAt       *time.Time    `json:"updated_at,omitempty"`
	Permissions     Permissions   `json:"permissions,omitempty"`
	Organization    *Organization `json:"organization,omitempty"`
	Parent          *Repository   `json:"parent,omitempty"`
	Source          *Repository   `json:"source,omitempty"`
	HasIssues       bool          `json:"has_issues,omitempty"`
	HasWiki         bool          `json:"has_wiki,omitempty"`
	HasDownloads    bool          `json:"has_downloads,omitempty"`
}

// Permissions represent the permissions as they apply to the accessing url
type Permissions struct {
	Admin bool
	Push  bool
	Pull  bool
}
