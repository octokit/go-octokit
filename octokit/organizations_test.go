package octokit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationService_Repos(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/orgs/rails/repos", "repository", nil)

	url, err := OrganizationReposURL.Expand(nil)
	assert.NoError(t, err)

	organizationResults, result := client.Organization(url).OrganizationRepos()

	assert.False(t, result.HasError())
	assert.False(t, organizationResults.IncompleteResults)
	assert.Equal(t, organizationResults.TotalCount, 2)
	assert.Equal(t, len(organizationResults.Items), 2)
	assert.Equal(t, organizationResults.Items[0].ID, 3338555)
	assert.Equal(t, organizationResults.Items[0].Login, "dhruvsinghal")
	assert.Equal(t, organizationResults.Items[1].ID, 9294878)
	assert.Equal(t, organizationResults.Items[1].Login, "dhruvsinghal5")
}

func TestOrganizationService_Info(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/orgs/{octokit}", "repository", nil)

	url, err := SearchURL.Expand()
	assert.NoError(t, err)

	organizationResults, result := client.OrganizationRepos(url).Users()

	assert.False(t, result.HasError())
	assert.False(t, organizationResults.IncompleteResults)
	assert.Equal(t, organizationResults.TotalCount, 2)
	assert.Equal(t, len(organizationResults.Items), 2)
	assert.Equal(t, organizationResults.Items[0].ID, 3338555)
	assert.Equal(t, organizationResults.Items[0].Login, "dhruvsinghal")
	assert.Equal(t, organizationResults.Items[1].ID, 9294878)
	assert.Equal(t, organizationResults.Items[1].Login, "dhruvsinghal5")
}

func TestFollowersService_Failure(t *testing.T) {
	setup()
	defer tearDown()

	url := Hyperlink("}")
	followers, result := client.Followers().All(&url, nil)
	assert.True(t, result.HasError())
	assert.Len(t, followers, 0)

	success, result := client.Followers().Check(&url, nil)
	assert.True(t, result.HasError())
	assert.False(t, success)

	success, result = client.Followers().Follow(&url, nil)
	assert.True(t, result.HasError())
	assert.False(t, success)

	success, result = client.Followers().Unfollow(&url, nil)
	assert.True(t, result.HasError())
	assert.False(t, success)
}
