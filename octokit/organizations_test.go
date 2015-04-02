package octokit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOrganizationService_Repos(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/orgs/rails/repos", "repositories", nil)

	organizationResults, result := client.Organization().OrganizationRepos(nil, M{"org": "rails"})

	assert.False(t, result.HasError())
	assert.Equal(t, 30, len(organizationResults))
	assert.Equal(t, 8514, organizationResults[0].ID)
	assert.Equal(t, "rails", organizationResults[0].Name)
	assert.Equal(t, 13992, organizationResults[1].ID)
	assert.Equal(t, "docrails", organizationResults[1].Name)
}

func TestOrganizationService_Info(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/orgs/octokit", "organization", nil)

	organizationResults, result := client.Organization().OrganizationInfo(nil, M{"org": "octokit"})

	assert.False(t, result.HasError())
	assert.Equal(t, "octokit", organizationResults.Login)
	assert.Equal(t, 3430433, organizationResults.ID)
}

func TestOrganizationService_Failure(t *testing.T) {
	setup()
	defer tearDown()
	url := Hyperlink("}")
	organizationResultsRepo, result := client.Organization().OrganizationRepos(&url, nil)
	assert.True(t, result.HasError())
	assert.Equal(t, make([]Repository, 0), organizationResultsRepo)

	organizationResult, result := client.Organization().OrganizationInfo(&url, nil)
	assert.True(t, result.HasError())
	assert.Equal(t, Organization{}, organizationResult)

	organizationResults, result := client.Organization().YourOrganizations(&url, nil)
	assert.True(t, result.HasError())
	assert.Equal(t, make([]Organization, 0), organizationResults)

	organizationResults, result = client.Organization().UserOrganizations(&url, nil)
	assert.True(t, result.HasError())
	assert.Equal(t, make([]Organization, 0), organizationResults)
}
