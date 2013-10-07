package octokat

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestPullRequest(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/jingweno/octokat/pulls/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("pull_request.json"))
	})

	repo := Repo{UserName: "jingweno", Name: "octokat"}
	pr, _ := client.PullRequest(repo, "1", nil)

	assert.Equal(t, 1, pr.ChangedFiles)
	assert.Equal(t, 1, pr.Deletions)
	assert.Equal(t, 1, pr.Additions)
	assert.Equal(t, "aafce5dfc44270f35270b24677abbb859b20addf", pr.MergeCommitSha)
	assert.Equal(t, "2013-06-09 00:53:38 +0000 UTC", pr.MergedAt.String())
	assert.Equal(t, "2013-06-09 00:53:38 +0000 UTC", pr.ClosedAt.String())
	assert.Equal(t, "2013-06-19 00:35:24 +0000 UTC", pr.UpdatedAt.String())
	assert.Equal(t, "2013-06-09 00:52:12 +0000 UTC", pr.CreatedAt.String())
	assert.Equal(t, "typo", pr.Body)
	assert.Equal(t, "Update README.md", pr.Title)
	assert.Equal(t, "https://api.github.com/repos/jingweno/octokat/pulls/1", pr.URL)
	assert.Equal(t, 6206442, pr.ID)
	assert.Equal(t, "https://github.com/jingweno/octokat/pull/1", pr.HTMLURL)
	assert.Equal(t, "https://github.com/jingweno/octokat/pull/1.diff", pr.DiffURL)
	assert.Equal(t, "https://github.com/jingweno/octokat/pull/1.patch", pr.PatchURL)
	assert.Equal(t, "https://github.com/jingweno/octokat/pull/1", pr.IssueURL)
	assert.Equal(t, 1, pr.Number)
	assert.Equal(t, "closed", pr.State)
	assert.T(t, nil == pr.Assignee)
	assert.Equal(t, "https://github.com/jingweno/octokat/pull/1/commits", pr.CommitsURL)
	assert.Equal(t, "https://github.com/jingweno/octokat/pull/1/comments", pr.ReviewCommentsURL)
	assert.Equal(t, "/repos/jingweno/octokat/pulls/comments/{number}", pr.ReviewCommentURL)
	assert.Equal(t, "https://api.github.com/repos/jingweno/octokat/issues/1/comments", pr.CommentsURL)
}
