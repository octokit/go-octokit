package octokit

import (
	"github.com/lostisland/go-sawyer/hypermedia"
	"time"
)

var (
	PullRequestsURL = Hyperlink("/repos/{owner}/{repo}/pulls{/number}")
)

func (c *Client) CreatePullRequest(owner, repo string, params *PullRequestParams) (pull *PullRequest, result *Result) {
	url, err := c.Root().PullsURL.Expand(M{"owner": owner, "repo": repo})
	if err != nil {
		result = newResult(nil, err)
		return
	}

	result = c.post(url, params, &pull)
	return
}

type PullRequest struct {
	*hypermedia.HALResource

	URL               string     `json:"url,omitempty"`
	ID                int        `json:"id,omitempty"`
	HTMLURL           string     `json:"html_url,omitempty"`
	DiffURL           string     `json:"diff_url,omitempty"`
	PatchURL          string     `json:"patch_url,omitempty"`
	IssueURL          string     `json:"issue_url,omitempty"`
	Number            int        `json:"number,omitempty"`
	State             string     `json:"state,omitempty"`
	Title             string     `json:"title,omitempty"`
	User              User       `json:"user,omitempty"`
	Body              string     `json:"body,omitempty"`
	CreatedAt         time.Time  `json:"created_at,omitempty"`
	UpdatedAt         time.Time  `json:"updated_at,omitempty"`
	ClosedAt          *time.Time `json:"closed_at,omitempty"`
	MergedAt          *time.Time `json:"merged_at,omitempty"`
	MergeCommitSha    string     `json:"merge_commit_sha,omitempty"`
	Assignee          *User      `json:"assignee,omitempty"`
	CommitsURL        string     `json:"commits_url,omitempty"`
	ReviewCommentsURL string     `json:"review_comments_url,omitempty"`
	ReviewCommentURL  string     `json:"review_comment_url,omitempty"`
	CommentsURL       string     `json:"comments_url,omitempty"`
	Head              Commit     `json:"head,omitempty"`
	Base              Commit     `json:"base,omitempty"`
	Merged            bool       `json:"merged,omitempty"`
	MergedBy          User       `json:"merged_by,omitempty"`
	Comments          int        `json:"comments,omitempty"`
	ReviewComments    int        `json:"review_comments,omitempty"`
	Commits           int        `json:"commits,omitempty"`
	Additions         int        `json:"additions,omitempty"`
	Deletions         int        `json:"deletions,omitempty"`
	ChangedFiles      int        `json:"changed_files,omitempty"`
}

type Commit struct {
	Label string     `json:"label,omitempty"`
	Ref   string     `json:"ref,omitempty"`
	Sha   string     `json:"sha,omitempty"`
	User  User       `json:"user,omitempty"`
	Repo  Repository `json:"repo,omitempty"`
}

type PullRequestParams struct {
	Base  string `json:"base,omitempty"`
	Head  string `json:"head,omitempty"`
	Title string `json:"title,omitempty"`
	Body  string `json:"body,omitempty"`
}

type PullRequestForIssueParams struct {
	Base  string `json:"base,omitempty"`
	Head  string `json:"head,omitempty"`
	Issue string `json:"issue,omitempty"`
}
