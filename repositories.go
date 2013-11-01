package octokit

import (
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
