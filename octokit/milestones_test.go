package octokit

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMilestonesService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octokit/go-octokit/milestones", "milestones", nil)

	milestones, result := client.Milestones().All(nil, M{"owner": "octokit", "repo": "go-octokit"})

	assert.False(t, result.HasError())

	assert.Equal(t, 1, len(milestones))

	assert.Equal(t, "https://api.github.com/repos/octocat/Hello-World/milestones/1", milestones[0].URL)
	assert.Equal(t, 1, milestones[0].Number)
	assert.Equal(t, "open", milestones[0].State)
	assert.Equal(t, "v1.0", milestones[0].Title)
	assert.Equal(t, "", milestones[0].Description)
	assert.Equal(t, 4, milestones[0].OpenIssues)
	assert.Equal(t, 8, milestones[0].ClosedIssues)
	assert.Equal(t, "2011-04-10 20:09:31 +0000 UTC", milestones[0].CreatedAt.String())
	assert.Equal(t, (*time.Time)(nil), milestones[0].DueOn)

	assert.NotNil(t, milestones[0].Creator)
	assert.Equal(t, "octocat", milestones[0].Creator.Login)
	assert.Equal(t, 1, milestones[0].Creator.ID)
	assert.Equal(t, "https://github.com/images/error/octocat_happy.gif", milestones[0].Creator.AvatarURL)
	assert.Equal(t, "somehexcode", milestones[0].Creator.GravatarID)
	assert.Equal(t, "https://api.github.com/users/octocat", milestones[0].Creator.URL)
}

func TestMilestonesService_One(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octokit/go-octokit/milestones", "milestone_created", nil)

	milestone, result := client.Milestones().One(nil, M{"owner": "octokit", "repo": "go-octokit"})

	assert.False(t, result.HasError())

	validateMilestone(t, *milestone)
}

func TestMilestonesService_Create(t *testing.T) {
	setup()
	defer tearDown()

	input := M{"title": "I am a title", "state": "open", "description": "I am a description."}
	wantReqBody, _ := json.Marshal(input)
	stubPost(t, "/repos/octokit/go-octokit/milestones", "milestone_created", nil, string(wantReqBody)+"\n", nil)

	milestone, result := client.Milestones().Create(nil, M{"owner": "octokit", "repo": "go-octokit"}, input)

	assert.False(t, result.HasError())

	validateMilestone(t, *milestone)
}

func TestMilestonesService_Delete(t *testing.T) {
	setup()
	defer tearDown()

	var respHeaderParams = map[string]string{"Content-Type": "application/json"}
	stubDeletewCode(t, "/repos/octokit/go-octokit/milestones", respHeaderParams, 204)

	success, result := client.Milestones().Delete(nil, M{"owner": "octokit", "repo": "go-octokit", "id": 19158753})
	assert.False(t, result.HasError())

	assert.True(t, success)
}

func validateMilestone(t *testing.T, milestone Milestone) {
	assert.Equal(t, "open", milestone.State)

	assert.Equal(t, "https://api.github.com/repos/octocat/Hello-World/milestones/1", milestone.URL)
	assert.Equal(t, "https://github.com/octocat/Hello-World/milestones/v1.0", milestone.HTMLURL)
	assert.Equal(t, "https://api.github.com/repos/octocat/Hello-World/milestones/1/labels", milestone.LabelsURL)
	assert.Equal(t, 1, milestone.Number)
	assert.Equal(t, 1002604, milestone.ID)
	assert.Equal(t, "open", milestone.State)
	assert.Equal(t, "v1.0", milestone.Title)
	assert.Equal(t, "Tracking milestone for version 1.0", milestone.Description)
	assert.Equal(t, 4, milestone.OpenIssues)
	assert.Equal(t, 8, milestone.ClosedIssues)
	assert.Equal(t, "2011-04-10 20:09:31 +0000 UTC", milestone.CreatedAt.String())
	assert.Equal(t, "2014-03-03 18:58:10 +0000 UTC", milestone.UpdatedAt.String())
	assert.Equal(t, "2013-02-12 13:22:01 +0000 UTC", milestone.ClosedAt.String())
	assert.Equal(t, "2012-10-09 23:39:01 +0000 UTC", milestone.DueOn.String())

	user := milestone.Creator

	assert.Equal(t, "octocat", user.Login)
	assert.Equal(t, 1, user.ID)
}
