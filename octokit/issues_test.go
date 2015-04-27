package octokit

import (
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIssuesService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octocat/Hello-World/issues", "issues", nil)

	url, err := RepoIssuesURL.Expand(M{"owner": "octocat", "repo": "Hello-World"})
	assert.NoError(t, err)

	issues, result := client.Issues(url).All()
	assert.False(t, result.HasError())
	assert.Len(t, issues, 1)

	issue := issues[0]
	validateIssue(t, issue)
}

func TestIssuesService_One(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octocat/Hello-World/issues/1347", "issue", nil)

	url, err := RepoIssuesURL.Expand(M{"owner": "octocat", "repo": "Hello-World", "number": 1347})
	assert.NoError(t, err)

	issue, result := client.Issues(url).One()

	assert.False(t, result.HasError())
	validateIssue(t, *issue)
}

func TestIssuesService_Create(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octocat/Hello-World/issues", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, "{\"title\":\"title\",\"body\":\"body\"}\n")
		respondWithJSON(w, loadFixture("issue.json"))
	})

	url, err := RepoIssuesURL.Expand(M{"owner": "octocat", "repo": "Hello-World"})
	assert.NoError(t, err)

	params := IssueParams{
		Title: "title",
		Body:  "body",
	}
	issue, result := client.Issues(url).Create(params)

	assert.False(t, result.HasError())
	validateIssue(t, *issue)
}

func TestIssuesService_Update(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octocat/Hello-World/issues/1347", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testBody(t, r, "{\"title\":\"title\",\"body\":\"body\"}\n")
		respondWithJSON(w, loadFixture("issue.json"))
	})

	url, err := RepoIssuesURL.Expand(M{"owner": "octocat", "repo": "Hello-World", "number": 1347})
	assert.NoError(t, err)

	params := IssueParams{
		Title: "title",
		Body:  "body",
	}
	issue, result := client.Issues(url).Update(params)

	assert.False(t, result.HasError())
	validateIssue(t, *issue)
}

func validateIssue(t *testing.T, issue Issue) {

	assert.EqualValues(t, "https://api.github.com/repos/octocat/Hello-World/issues/1347", issue.URL)
	assert.EqualValues(t, "https://github.com/octocat/Hello-World/issues/1347", issue.HTMLURL)
	assert.EqualValues(t, 1347, issue.Number)
	assert.EqualValues(t, "open", issue.State)
	assert.EqualValues(t, "Found a bug", issue.Title)
	assert.EqualValues(t, "I'm having a problem with this.", issue.Body)

	assert.EqualValues(t, "octocat", issue.User.Login)
	assert.EqualValues(t, 1, issue.User.ID)
	assert.EqualValues(t, "https://github.com/images/error/octocat_happy.gif", issue.User.AvatarURL)
	assert.EqualValues(t, "somehexcode", issue.User.GravatarID)
	assert.EqualValues(t, "https://api.github.com/users/octocat", issue.User.URL)

	assert.Len(t, issue.Labels, 1)
	assert.EqualValues(t, "https://api.github.com/repos/octocat/Hello-World/labels/bug", issue.Labels[0].URL)
	assert.EqualValues(t, "bug", issue.Labels[0].Name)

	assert.EqualValues(t, "octocat", issue.Assignee.Login)
	assert.EqualValues(t, 1, issue.Assignee.ID)
	assert.EqualValues(t, "https://github.com/images/error/octocat_happy.gif", issue.Assignee.AvatarURL)
	assert.EqualValues(t, "somehexcode", issue.Assignee.GravatarID)
	assert.EqualValues(t, "https://api.github.com/users/octocat", issue.Assignee.URL)

	assert.EqualValues(t, "https://api.github.com/repos/octocat/Hello-World/milestones/1", issue.Milestone.URL)
	assert.EqualValues(t, 1, issue.Milestone.Number)
	assert.EqualValues(t, "open", issue.Milestone.State)
	assert.EqualValues(t, "v1.0", issue.Milestone.Title)
	assert.EqualValues(t, "", issue.Milestone.Description)

	assert.EqualValues(t, "octocat", issue.Milestone.Creator.Login)
	assert.EqualValues(t, 1, issue.Milestone.Creator.ID)
	assert.EqualValues(t, "https://github.com/images/error/octocat_happy.gif", issue.Milestone.Creator.AvatarURL)
	assert.EqualValues(t, "somehexcode", issue.Milestone.Creator.GravatarID)
	assert.EqualValues(t, "https://api.github.com/users/octocat", issue.Milestone.Creator.URL)

	assert.EqualValues(t, 4, issue.Milestone.OpenIssues)
	assert.EqualValues(t, 8, issue.Milestone.ClosedIssues)
	assert.EqualValues(t, "2011-04-10 20:09:31 +0000 UTC", issue.Milestone.CreatedAt.String())
	assert.EqualValues(t, (*time.Time)(nil), issue.Milestone.DueOn)

	assert.EqualValues(t, 0, issue.Comments)
	assert.EqualValues(t, "https://github.com/octocat/Hello-World/pull/1347", issue.PullRequest.HTMLURL)
	assert.EqualValues(t, "https://github.com/octocat/Hello-World/pull/1347.diff", issue.PullRequest.DiffURL)
	assert.EqualValues(t, "https://github.com/octocat/Hello-World/pull/1347.patch", issue.PullRequest.PatchURL)

	assert.EqualValues(t, (*time.Time)(nil), issue.ClosedAt)
	assert.EqualValues(t, "2011-04-22 13:33:48 +0000 UTC", issue.CreatedAt.String())
	assert.EqualValues(t, "2011-04-22 13:33:48 +0000 UTC", issue.UpdatedAt.String())
}
