package octokat

import (
	"github.com/bmizerany/assert"
	"github.com/octokit/octokat/hyper"
	"testing"
)

func TestClient_Root(t *testing.T) {
	setup()
	defer tearDown()

	root, _ := client.Root(nil)
	assert.Equal(t, hyper.Link(testURLOf("users/{user}")), root.UserURL)
}
