package main

import "time"

type CommitFile struct {
	Additions   int    `json:"additions"`
	BlobURL     string `json:"blob_url"`
	Changes     int    `json:"changes"`
	ContentsURL string `json:"contents_url"`
	Deletions   int    `json:"deletions"`
	Filename    string `json:"filename"`
	Patch       string `json:"patch"`
	RawURL      string `json:"raw_url"`
	Sha         string `json:"sha"`
	Status      string `json:"status"`
}

type CommitStats struct {
	Additions int `json:"additions"`
	Deletions int `json:"deletions"`
	Total     int `json:"total"`
}

type CommitCommit struct {
	Author struct {
		Date  *time.Time `json:"date"`
		Email string     `json:"email"`
		Name  string     `json:"name"`
	} `json:"author"`
	CommentCount int `json:"comment_count"`
	Committer    struct {
		Date  *time.Time `json:"date"`
		Email string     `json:"email"`
		Name  string     `json:"name"`
	} `json:"committer"`
	Message string `json:"message"`
	Tree    struct {
		Sha string `json:"sha"`
		URL string `json:"url"`
	} `json:"tree"`
	URL string `json:"url"`
}

type Commit struct {
	Author      *User         `json:"author"`
	CommentsURL string        `json:"comments_url"`
	Commit      *CommitCommit `json:"commit"`
	Committer   *User         `json:"committer"`
	Files       []CommitFile  `json:"files"`
	HtmlURL     string        `json:"html_url"`
	Parents     []Commit      `json:"parents"`
	Sha         string        `json:"sha"`
	Stats       CommitStats   `json:"stats"`
	URL         string        `json:"url"`
}
