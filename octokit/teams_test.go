package octokit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func assertTeam(t *testing.T, team Team, full bool) {
	assert.Equal(t, 1, team.ID)
	assert.Equal(t, "https://api.github.com/teams/1", string(team.URL))
	assert.Equal(t, "Justice League", team.Name)
	assert.Equal(t, "justice-league", team.Slug)
	assert.Equal(t, "A great team.", team.Description)
	assert.Equal(t, "closed", team.Privacy)
	assert.Equal(t, "admin", team.Permission)
	assert.Equal(t, "https://api.github.com/teams/1/members{/member}", string(team.MembersURL))
	assert.Equal(t, "https://api.github.com/teams/1/repos", string(team.RepositoriesURL))

	if full {
		assert.Equal(t, 3, team.MembersCount)
		assert.Equal(t, 10, team.ReposCount)
		assert.Equal(t, "github", team.Organization.Login)
		assert.Equal(t, "A great organization", team.Organization.Description)
	}
}

// lower-level Team APIs

func TestTeamsService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/orgs/org/teams", "teams", nil)

	teams, result := client.Teams().All(nil, M{"org": "org"})

	assert.False(t, result.HasError())
	assert.Len(t, teams, 1)
	assertTeam(t, teams[0], false)
}

func TestTeamsService_One(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/teams/1", "team", nil)

	team, result := client.Teams().One(nil, M{"id": 1})

	assert.False(t, result.HasError())
	assertTeam(t, team, true)
}

func TestTeamsService_Create(t *testing.T) {
	setup()
	defer tearDown()

	input := TeamParams{Name: "Justice League"}
	wantReqBody, _ := json.Marshal(input)

	stubPost(t, "/orgs/org/teams", "team", nil, string(wantReqBody)+"\n", nil)

	team, result := client.Teams().Create(nil, input, M{"org": "org"})

	assert.False(t, result.HasError())
	assertTeam(t, team, true)
}

func TestTeamsService_Update(t *testing.T) {
	setup()
	defer tearDown()

	input := TeamParams{Description: "A great team."}
	wantReqBody, _ := json.Marshal(input)

	stubPatch(t, "/teams/1", "team", nil, string(wantReqBody)+"\n", nil)

	team, result := client.Teams().Update(nil, input, M{"id": 1})

	assert.False(t, result.HasError())
	assertTeam(t, team, true)
}

func TestTeamsService_Delete(t *testing.T) {
	setup()
	defer tearDown()

	var respHeaderParams = map[string]string{"Content-Type": "application/json"}
	stubDeletewCode(t, "/teams/1", respHeaderParams, 204)

	success, result := client.Teams().Delete(nil, M{"id": 1})

	assert.False(t, result.HasError())
	assert.True(t, success)
}

// Higher-level convenience APIs

func TestTeamsService_GetTeams(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/orgs/org/teams", "teams", nil)

	teams, result := client.Organization().GetTeams(nil, M{"org": "org"})

	assert.False(t, result.HasError())
	assert.Len(t, teams, 1)
	assertTeam(t, teams[0], false)
}

func TestTeamsService_GetMembers(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/teams/1/members", "users", nil)

	users, result := client.Teams().GetMembers(nil, M{"id": 1})

	assert.False(t, result.HasError())
	assert.Len(t, users, 1)
}

func TestTeamsService_GetMembership(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/teams/1/memberships/SomeUser", "team_membership", nil)

	membership, result := client.Teams().GetMembership(nil, M{"id": 1, "username": "SomeUser"})

	assert.False(t, result.HasError())
	assert.Equal(t, "https://api.github.com/teams/1/memberships/SomeUser", membership.URL)
	assert.Equal(t, "member", membership.Role)
	assert.Equal(t, "active", membership.State)
}

func TestTeamsService_AddMembership(t *testing.T) {
	setup()
	defer tearDown()

	input := M{"role": "member"}
	wantReqBody, _ := json.Marshal(input)

	respHeaderParams := map[string]string{"Content-Type": "application/json"}
	stubPutwCode(t, "/teams/1/memberships/SomeUser", "team_membership", nil, string(wantReqBody)+"\n", respHeaderParams, 200)

	membership, result := client.Teams().AddMembership(nil, M{"id": 1, "username": "SomeUser"}, "member")

	assert.False(t, result.HasError())
	assert.Equal(t, "https://api.github.com/teams/1/memberships/SomeUser", membership.URL)
	assert.Equal(t, "member", membership.Role)
	assert.Equal(t, "active", membership.State)
}

func TestTeamsService_RemoveMembership(t *testing.T) {
	setup()
	defer tearDown()

	var respHeaderParams = map[string]string{"Content-Type": "application/json"}
	stubDeletewCode(t, "/teams/1/memberships/SomeUser", respHeaderParams, 204)

	success, result := client.Teams().RemoveMembership(nil, M{"id": 1, "username": "SomeUser"})

	assert.False(t, result.HasError())
	assert.True(t, success)
}

func TestTeamsService_GetRepositories(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/teams/1/repos", "repositories", nil)

	repos, result := client.Teams().GetRepositories(nil, M{"id": 1})

	assert.False(t, result.HasError())
	assert.Len(t, repos, 30)
}

func TestTeamsService_CheckRepository(t *testing.T) {
	setup()
	defer tearDown()

	var respHeaderParams = map[string]string{"Content-Type": "application/json"}
	stubGetwCode(t, "/teams/1/repos/SomeOrg/SomeRepo", "repository", respHeaderParams, 200)

	manages, _, result := client.Teams().CheckRepository(nil, M{"id": 1, "owner": "SomeOrg", "repo": "SomeRepo"})

	assert.False(t, result.HasError())
	assert.True(t, manages)
}

func TestTeamsService_UpdateRepository(t *testing.T) {
	setup()
	defer tearDown()

	input := M{"permission": "pull"}
	wantReqBody, _ := json.Marshal(input)

	var respHeaderParams = map[string]string{"Content-Type": "application/json"}
	stubPutwCode(t, "/teams/1/repos/SomeOrg/SomeRepo", "", nil, string(wantReqBody)+"\n", respHeaderParams, 204)

	success, result := client.Teams().UpdateRepository(nil, M{"id": 1, "owner": "SomeOrg", "repo": "SomeRepo"}, "pull")

	assert.False(t, result.HasError())
	assert.True(t, success)
}

func TestTeamsService_RemoveRepository(t *testing.T) {
	setup()
	defer tearDown()

	var respHeaderParams = map[string]string{"Content-Type": "application/json"}
	stubDeletewCode(t, "/teams/1/repos/SomeOrg/SomeRepo", respHeaderParams, 204)

	success, result := client.Teams().RemoveRepository(nil, M{"id": 1, "owner": "SomeOrg", "repo": "SomeRepo"})

	assert.False(t, result.HasError())
	assert.True(t, success)
}
