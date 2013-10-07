package octokat

import (
	"fmt"
	"time"
)

type Repository struct {
	ID            int           `json:"id,omitempty"`
	Owner         User          `json:"owner,omitempty"`
	Name          string        `json:"name,omitempty"`
	FullName      string        `json:"full_name,omitempty"`
	Description   string        `json:"description,omitempty"`
	Private       bool          `json:"private,omitempty"`
	Fork          bool          `json:"fork,omitempty"`
	URL           string        `json:"url,omitempty"`
	HTMLURL       string        `json:"html_url,omitempty"`
	CloneURL      string        `json:"clone_url,omitempty"`
	GitURL        string        `json:"git_url,omitempty"`
	SSHURL        string        `json:"ssh_url,omitempty"`
	SVNURL        string        `json:"svn_url,omitempty"`
	MirrorURL     string        `json:"mirror_url,omitempty"`
	Homepage      string        `json:"homepage,omitempty"`
	Language      string        `json:"language,omitempty"`
	Forks         int           `json:"forks,omitempty"`
	ForksCount    int           `json:"forks_count,omitempty"`
	Watchers      int           `json:"watchers,omitempty"`
	WatchersCount int           `json:"watchers_count,omitempty"`
	Size          int           `json:"size,omitempty"`
	MasterBranch  string        `json:"master_branch,omitempty"`
	OpenIssues    int           `json:"open_issues,omitempty"`
	PushedAt      time.Time     `json:"pushed_at,omitempty"`
	CreatedAt     time.Time     `json:"created_at,omitempty"`
	UpdatedAt     time.Time     `json:"updated_at,omitempty"`
	Organization  *Organization `json:"organization,omitempty"`
	Parent        *Repository   `json:"parent,omitempty"`
	Source        *Repository   `json:"source,omitempty"`
	HasIssues     bool          `json:"has_issues,omitempty"`
	HasWiki       bool          `json:"has_wiki,omitempty"`
	HasDownloads  bool          `json:"has_downloads,omitempty"`
}

// List repositories
//
// If username is not supplied, repositories for the current
// authenticated user are returned
//
// See http://developer.github.com/v3/repos/#list-your-repositories
func (c *Client) Repositories(username string, options *Options) (repositories []Repository, err error) {
	var path string
	if username == "" {
		path = "user/repos"
	} else {
		path = fmt.Sprintf("users/%s/repos", username)
	}

	err = c.jsonGet(path, options, &repositories)
	return
}

// Get a single repository
//
// See http://developer.github.com/v3/repos/#get
func (c *Client) Repository(repo Repo, options *Options) (repository *Repository, err error) {
	path := fmt.Sprintf("repos/%s", repo)
	err = c.jsonGet(path, options, &repository)
	return
}

type RepositoryParams struct {
	Name              string `json:"name,omitempty"`
	Description       string `json:"description,omitempty"`
	Homepage          string `json:"homepage,omitempty"`
	Private           bool   `json:"private,omitempty"`
	HasIssues         bool   `json:"has_issues,omitempty"`
	HasWiki           bool   `json:"has_wiki,omitempty"`
	HasDownloads      bool   `json:"has_downloads,omitempty"`
	TeamID            int    `json:"team_id,omitempty"`
	AutoInit          bool   `json:"auto_init,omitempty"`
	GitignoreTemplate string `json:"gitignore_template,omitempty"`
}

// Create a repository for a user or organization
//
// If org is not specified, create a repository for current user.
//
// See http://developer.github.com/v3/repos/#create
func (c *Client) CreateRepository(org string, options *Options) (repository *Repository, err error) {
	var path string
	if org == "" {
		path = "user/repos"
	} else {
		path = fmt.Sprintf("orgs/%s/repos", org)
	}

	err = c.jsonPost(path, options, &repository)
	return
}

// Fork a repository
//
// See http://developer.github.com/v3/repos/forks/#create-a-fork
func (c *Client) Fork(repo Repo, options *Options) (repository *Repository, err error) {
	path := fmt.Sprintf("repos/%s/forks", repo)
	err = c.jsonPost(path, options, &repository)
	return
}
