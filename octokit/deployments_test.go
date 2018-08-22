package octokit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeploymentsService_All(t *testing.T) {
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

	validateDeployment(t, deployments[0])

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

func TestDeploymentsService_Create(t *testing.T) {
	setup()
	defer tearDown()

	params := DeploymentParams{
		Ref: "topic-branch",
		Payload: "{\"user\":\"atmos\",\"room_id\":123456}",
		Description: "Deploying my sweet branch",
	}
	wantReqBody, _ := json.Marshal(params)
	stubPost(t, "/repos/octocat/Hello-World/deployments", "deployment", nil, string(wantReqBody)+"\n", nil)

	deployment, result := client.Deployments().Create(nil, M{"owner": "octocat",
		"repo": "Hello-World"}, params)

	assert.False(t, result.HasError())
	validateDeployment(t, *deployment)
}

func validateDeployment(t *testing.T, deployment Deployment) {
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
}
