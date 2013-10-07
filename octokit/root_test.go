package octokit

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestClient_Root(t *testing.T) {
	setup()
	defer tearDown()

	root, _ := client.Root()
	assert.Equal(t, Hyperlink(testURLStringOf("users/{user}")), root.UserURL)
}
