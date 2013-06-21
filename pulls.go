package octokat

import (
	"fmt"
	"time"
)

type PullRequestParams struct {
	Base  string `json:"base"`
	Head  string `json:"head"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

type PullRequestForIssueParams struct {
	Base  string `json:"base"`
	Head  string `json:"head"`
	Issue string `json:"issue"`
}

type PullRequest struct {
	URL               string    `json:"url"`
	Id                int       `json:"id"`
	HTMLURL           string    `json:"html_url"`
	DiffURL           string    `json:"diff_url"`
	PatchURL          string    `json:"patch_url"`
	IssueURL          string    `json:"issue_url"`
	Number            int       `json:"number"`
	State             string    `json:"state"`
	Title             string    `json:"title"`
	User              User      `json:"user"`
	Body              string    `json:"body"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ClosedAt          time.Time `json:"closed_at"`
	MergedAt          time.Time `json:"merged_at"`
	MergedCommitSha   string    `json:"merged_commit_sha"`
	Assignee          *User     `json:"assignee"`
	CommitsUrl        string    `json:"commits_url"`
	ReviewCommentsUrl string    `json:"review_comments_url"`
	ReviewCommentUrl  string    `json:"review_comment_url"`
	CommentsUrl       string    `json:"comments_url"`
	Head              Commit    `json:"head"`
	Base              Commit    `json:"base"`
	Merged            bool      `json:"merged"`
	MergedBy          User      `json:"merged_by"`
	Comments          int       `json:"comments"`
	ReviewComments    int       `json:"review_comments"`
	Commits           int       `json:"commits"`
	Additions         int       `json:"additions"`
	Deletions         int       `json:"deletions"`
	ChangedFiles      int       `json:"changed_files"`
}

type Commit struct {
	Label string     `json:"label"`
	Ref   string     `json:"ref"`
	Sha   string     `json:"sha"`
	User  User       `json:"user"`
	Repo  Repository `json:"repo"`
}

func (c *Client) PullRequest(repo Repo, number string) (*PullRequest, error) {
	path := fmt.Sprintf("repos/%s/pulls/%s", repo, number)

	var pr PullRequest
	c.jsonGet(path, nil, &pr)

	return &pr, nil
}

func (c *Client) CreatePullRequest(repo Repo, params PullRequestParams) (*PullRequest, error) {
	return c.createPullRequest(repo, params)
}

func (c *Client) CreatePullRequestForIssue(repo Repo, params PullRequestForIssueParams) (*PullRequest, error) {
	return c.createPullRequest(repo, params)
}

func (c *Client) createPullRequest(repo Repo, params interface{}) (*PullRequest, error) {
	path := fmt.Sprintf("repos/%s/pulls", repo)
	var pr PullRequest
	err := c.jsonPost(path, nil, params, &pr)
	if err != nil {
		return nil, err
	}

	return &pr, nil
}
