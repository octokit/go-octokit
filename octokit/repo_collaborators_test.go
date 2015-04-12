package octokit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCollaboratorsService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octokit/go-octokit/collaborators", "collaborators",
		nil)

	collabs, result := client.Collaborators().All(nil, M{"owner": "octokit",
		"repo": "go-octokit"})

	fmt.Println(result.Error())
	assert.False(t, result.HasError())
	assert.Len(t, collabs, 24)
}
