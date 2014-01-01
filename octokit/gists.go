package octokit

import (
	"time"
)

type Gist struct {
	URL         string     `json:"url,omitempty"`
	ForksURL    string     `json:"forks_url,omitempty"`
	CommitsURL  string     `json:"commits_url,omitempty"`
	ID          int        `json:"id,omitempty"`
	Description string     `json:"description,omitempty"`
	Public      bool       `json:"public,omitempty"`
	User        *User      `json:user,omitempty"`
	Files       GistFiles  `json:files,omitempty"`
	Comments    int        `json:comments,omitempty"`
	CommentsURL string     `json:comments_url,omitempty"`
	HTMLURL     string     `json:html_url,omitempty"`
	GitPullURL  string     `json:git_pull_url,omitempty"`
	GitPushURL  string     `json:git_push_url,omitempty"`
	CreatedAt   *time.Time `json:created_at,omitempty"`
	UpdatedAt   *time.Time `json:updated_at,omitempty"`
}

type GistFile struct {
	Size     int    `json:"size,omitempty"`
	Filename string `json:"filename,omitempty"`
	RawURL   string `json:"raw_url,omitempty"`
	Type     string `json:"type,omitempty"`
	Language string `json:"language,omitempty"`
}

type GistFiles map[string]GistFile
