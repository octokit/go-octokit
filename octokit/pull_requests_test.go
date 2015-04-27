package octokit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPullRequestService_One(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octokit/go-octokit/pulls/1", "pull_request", nil)

	url, err := PullRequestsURL.Expand(M{"owner": "octokit", "repo": "go-octokit", "number": 1})
	assert.NoError(t, err)

	pr, result := client.PullRequests(url).One()

	assert.False(t, result.HasError())
	assert.EqualValues(t, 1, pr.ChangedFiles)
	assert.EqualValues(t, 1, pr.Deletions)
	assert.EqualValues(t, 1, pr.Additions)
	assert.EqualValues(t, "aafce5dfc44270f35270b24677abbb859b20addf", pr.MergeCommitSha)
	assert.EqualValues(t, "2013-06-09 00:53:38 +0000 UTC", pr.MergedAt.String())
	assert.EqualValues(t, "2013-06-09 00:53:38 +0000 UTC", pr.ClosedAt.String())
	assert.EqualValues(t, "2013-06-19 00:35:24 +0000 UTC", pr.UpdatedAt.String())
	assert.EqualValues(t, "2013-06-09 00:52:12 +0000 UTC", pr.CreatedAt.String())
	assert.EqualValues(t, "typo", pr.Body)
	assert.EqualValues(t, "Update README.md", pr.Title)
	assert.EqualValues(t, "https://api.github.com/repos/jingweno/octokat/pulls/1", pr.URL)
	assert.EqualValues(t, 6206442, pr.ID)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1", pr.HTMLURL)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1.diff", pr.DiffURL)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1.patch", pr.PatchURL)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1", pr.IssueURL)
	assert.EqualValues(t, 1, pr.Number)
	assert.EqualValues(t, "closed", pr.State)

	assert.EqualValues(t, "octocat", pr.Assignee.Login)
	assert.EqualValues(t, 1, pr.Assignee.ID)
	assert.EqualValues(t, "https://github.com/images/error/octocat_happy.gif", pr.Assignee.AvatarURL)
	assert.EqualValues(t, "somehexcode", pr.Assignee.GravatarID)
	assert.EqualValues(t, "https://api.github.com/users/octocat", pr.Assignee.URL)

	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1/commits", pr.CommitsURL)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1/comments", pr.ReviewCommentsURL)
	assert.EqualValues(t, "/repos/jingweno/octokat/pulls/comments/{number}", pr.ReviewCommentURL)
	assert.EqualValues(t, "https://api.github.com/repos/jingweno/octokat/issues/1/comments", pr.CommentsURL)

	assert.Nil(t, pr.Mergeable)
}

func TestPullRequestService_Post(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/go-octokit/pulls", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r,
			"{\"base\":\"base\",\"head\":\"head\",\"title\":\"title\",\"body\":\"body\",\"assignee\":\"assignee\"}\n")
		respondWithJSON(w, loadFixture("pull_request.json"))
	})

	url, err := PullRequestsURL.Expand(M{"owner": "octokit", "repo": "go-octokit"})
	assert.NoError(t, err)

	params := PullRequestParams{
		Base:     "base",
		Head:     "head",
		Title:    "title",
		Body:     "body",
		Assignee: "assignee",
	}
	pr, result := client.PullRequests(url).Create(params)

	assert.False(t, result.HasError())
	assert.EqualValues(t, 1, pr.ChangedFiles)
	assert.EqualValues(t, 1, pr.Deletions)
	assert.EqualValues(t, 1, pr.Additions)
	assert.EqualValues(t, "aafce5dfc44270f35270b24677abbb859b20addf", pr.MergeCommitSha)
	assert.EqualValues(t, "2013-06-09 00:53:38 +0000 UTC", pr.MergedAt.String())
	assert.EqualValues(t, "2013-06-09 00:53:38 +0000 UTC", pr.ClosedAt.String())
	assert.EqualValues(t, "2013-06-19 00:35:24 +0000 UTC", pr.UpdatedAt.String())
	assert.EqualValues(t, "2013-06-09 00:52:12 +0000 UTC", pr.CreatedAt.String())
	assert.EqualValues(t, "typo", pr.Body)
	assert.EqualValues(t, "Update README.md", pr.Title)
	assert.EqualValues(t, "https://api.github.com/repos/jingweno/octokat/pulls/1", pr.URL)
	assert.EqualValues(t, 6206442, pr.ID)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1", pr.HTMLURL)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1.diff", pr.DiffURL)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1.patch", pr.PatchURL)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1", pr.IssueURL)
	assert.EqualValues(t, 1, pr.Number)
	assert.EqualValues(t, "closed", pr.State)

	assert.EqualValues(t, "octocat", pr.Assignee.Login)
	assert.EqualValues(t, 1, pr.Assignee.ID)
	assert.EqualValues(t, "https://github.com/images/error/octocat_happy.gif", pr.Assignee.AvatarURL)
	assert.EqualValues(t, "somehexcode", pr.Assignee.GravatarID)
	assert.EqualValues(t, "https://api.github.com/users/octocat", pr.Assignee.URL)

	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1/commits", pr.CommitsURL)
	assert.EqualValues(t, "https://github.com/jingweno/octokat/pull/1/comments", pr.ReviewCommentsURL)
	assert.EqualValues(t, "/repos/jingweno/octokat/pulls/comments/{number}", pr.ReviewCommentURL)
	assert.EqualValues(t, "https://api.github.com/repos/jingweno/octokat/issues/1/comments", pr.CommentsURL)
}

func TestPullRequestService_All(t *testing.T) {
	setup()
	defer tearDown()

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("repositories/8514/pulls?page=2"), testURLOf("repositories/8514/pulls?page=14"))
	stubGet(t, "/repos/rails/rails/pulls", "pull_requests", map[string]string{"Link": link})

	url, err := PullRequestsURL.Expand(M{"owner": "rails", "repo": "rails"})
	assert.NoError(t, err)

	prs, result := client.PullRequests(url).All()
	assert.False(t, result.HasError())
	assert.Len(t, prs, 30)
	assert.EqualValues(t, testURLStringOf("repositories/8514/pulls?page=2"), string(*result.NextPage))
	assert.EqualValues(t, testURLStringOf("repositories/8514/pulls?page=14"), string(*result.LastPage))
}

func TestPullRequestService_Diff(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/go-octokit/pulls/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", diffMediaType)
		respondWith(w, "diff --git")
	})

	url, err := PullRequestsURL.Expand(M{"owner": "octokit", "repo": "go-octokit", "number": 1})
	assert.NoError(t, err)

	diff, result := client.PullRequests(url).Diff()

	assert.False(t, result.HasError())
	content, err := ioutil.ReadAll(diff)
	assert.NoError(t, err)
	assert.EqualValues(t, "diff --git", string(content))
}

func TestPullRequestService_Patch(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/go-octokit/pulls/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", patchMediaType)
		respondWith(w, "patches galore")
	})

	url, err := PullRequestsURL.Expand(M{"owner": "octokit", "repo": "go-octokit", "number": 1})
	assert.NoError(t, err)

	patch, result := client.PullRequests(url).Patch()

	assert.False(t, result.HasError())
	content, err := ioutil.ReadAll(patch)
	assert.NoError(t, err)
	assert.EqualValues(t, "patches galore", string(content))
}

func TestPullRequestService_Unmergeable(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octokit/go-octokit/pulls/1", "unmergeable_pull_request", nil)

	url, err := PullRequestsURL.Expand(M{"owner": "octokit", "repo": "go-octokit", "number": 1})
	assert.NoError(t, err)

	pr, result := client.PullRequests(url).One()

	assert.False(t, result.HasError())
	assert.NotNil(t, pr.Mergeable)
	assert.False(t, *pr.Mergeable)
	assert.EqualValues(t, MergeStateUnknown, pr.MergeableState)
}

func TestPullRequestService_Mergeable(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octokit/go-octokit/pulls/1", "mergeable_pull_request", nil)

	url, err := PullRequestsURL.Expand(M{"owner": "octokit", "repo": "go-octokit", "number": 1})
	assert.NoError(t, err)

	pr, result := client.PullRequests(url).One()

	assert.False(t, result.HasError())
	assert.NotNil(t, pr.Mergeable)
	assert.True(t, *pr.Mergeable)
	assert.EqualValues(t, MergeStateClean, pr.MergeableState)
}
