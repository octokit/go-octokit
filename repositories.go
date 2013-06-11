package octokat

import (
	"fmt"
	"time"
)

type Repository struct {
	Id            int           `json:"id"`
	Owner         User          `json:"owner"`
	Name          string        `json:"name"`
	FullName      string        `json:"full_name"`
	Description   string        `json:"description"`
	Private       bool          `json:"private"`
	Fork          bool          `json:"fork"`
	Url           string        `json:"url"`
	HtmlUrl       string        `json:"html_url"`
	CloneUrl      string        `json:"clone_url"`
	GitUrl        string        `json:"git_url"`
	SshUrl        string        `json:"ssh_url"`
	SvnUrl        string        `json:"svn_url"`
	MirrorUrl     string        `json:"mirror_url"`
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

type Organization User

func (c *Client) Repository(repo Repo) (*Repository, error) {
	path := fmt.Sprintf("repos/%s", repo)
	var repository Repository
	err := c.jsonGet(path, nil, &repository)
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
