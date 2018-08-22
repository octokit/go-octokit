package octokit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploymentsService_all(t *testing.T) {
	setup()
	defer tearDown()

	//Username URL
	stubGet(t, "/repos/octocat/Hello-World/deployments", "deployments", nil)

	deployments, result := client.Deployments().All(&DeploymentsURL, M{
		"owner": "octocat",
		"repo":  "Hello-World",
	})

	assert.False(t, result.HasError())
	assert.Len(t, deployments, 1)

	deployment := deployments[0]
	assert.Equal(t, "https://api.github.com/repos/octocat/example/deployments/1", deployment.URL)
	assert.Equal(t, 1, deployment.ID)
	assert.Equal(t, "MDEwOkRlcGxveW1lbnQx", deployment.NodeID)
	assert.Equal(t, "a84d88e7554fc1fa21bcbc4efae3c782a70d2b9d", deployment.Sha)
	assert.Equal(t, "master", deployment.Ref)
	assert.Equal(t, "deploy", deployment.Task)
	assert.Equal(t, "deploy:migrate", deployment.Payload.Task)

	assert.Equal(t, "https://github.com/images/error/octocat_happy.gif", deployment.Creator.AvatarURL)
	assert.Equal(t, "https://api.github.com/repos/octocat/example/deployments/1/statuses", deployment.StatusesURL)
	assert.Equal(t, "https://api.github.com/repos/octocat/example", deployment.RepositoryURL)

	//Nil case
	deploymentsNil, resultNil := client.Deployments().All(nil, M{
		"owner": "octocat",
		"repo":  "Hello-World",
	})

	assert.False(t, resultNil.HasError())
	assert.Equal(t, deploymentsNil, deployments)

	//Error case
	var invalid = Hyperlink("{")
	deploymentsErr, resultErr := client.Deployments().All(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.Equal(t, deploymentsErr, make([]Deployment, 0))
}
