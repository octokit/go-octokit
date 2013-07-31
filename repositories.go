package octokat

import (
	"fmt"
	"time"
)

type Repository struct {
	ID            int           `json:"id"`
	Owner         User          `json:"owner"`
	Name          string        `json:"name"`
	FullName      string        `json:"full_name"`
	Description   string        `json:"description"`
	Private       bool          `json:"private"`
	Fork          bool          `json:"fork"`
	URL           string        `json:"url"`
	HTMLURL       string        `json:"html_url"`
	CloneURL      string        `json:"clone_url"`
	GitURL        string        `json:"git_url"`
	SshURL        string        `json:"ssh_url"`
	SvnURL        string        `json:"svn_url"`
	MirrorURL     string        `json:"mirror_url"`
	Homepage      string        `json:"homepage"`
	Language      string        `json:"language"`
	Forks         int           `json:"forks"`
	ForksCount    int           `json:"forks_count"`
	Watchers      int           `json:"watchers"`
	WatchersCount int           `json:"watchers_count"`
	Size          int           `json:"size"`
	MasterBranch  string        `json:"master_branch"`
	OpenIssues    int           `json:"open_issues"`
	PushedAt      time.Time     `json:"pushed_at"`
	CreatedAt     time.Time     `json:"created_at"`
	UpdatedAt     time.Time     `json:"updated_at"`
	Organization  *Organization `json:"organization"`
	Parent        *Repository   `json:"parent"`
	Source        *Repository   `json:"source"`
	HasIssues     bool          `json:"has_issues"`
	HasWiki       bool          `json:"has_wiki"`
	HasDownloads  bool          `json:"has_downloads"`
}

func (c *Client) Repositories(username string, params *Params) ([]Repository, error) {
	var path string
	if username == "" {
		path = "user/repos"
	} else {
		path = fmt.Sprintf("users/%s/repos", username)
	}

	var repositories []Repository
	err := c.jsonGet(path, nil, &repositories)
	if err != nil {
		return nil, err
	}

	return repositories, nil
}

func (c *Client) Repository(repo Repo) (*Repository, error) {
	path := fmt.Sprintf("repos/%s", repo)
	var repository Repository
	err := c.jsonGet(path, nil, &repository)
	if err != nil {
		return nil, err
	}

	return &repository, nil
}

func (c *Client) CreateRepository(name string, params *Params) (*Repository, error) {
	organization := params.Delete("organization")
	params.Put("name", name)

	var path string
	if organization == nil {
		path = "user/repos"
	} else {
		path = fmt.Sprintf("orgs/%s/repos", organization)
	}

	var repository Repository
	err := c.jsonPost(path, nil, params, &repository)
	if err != nil {
		return nil, err
	}

	return &repository, nil
}

func (c *Client) Fork(repo Repo, params *Params) (*Repository, error) {
	path := fmt.Sprintf("repos/%s/forks", repo)
	var repository Repository
	err := c.jsonPost(path, nil, params, &repository)
	if err != nil {
		return nil, err
	}

	return &repository, nil
}
