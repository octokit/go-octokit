package octokat

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestClient_Root(t *testing.T) {
	setup()
	defer tearDown()

	var root Root
	requester := client.Root()
	requester.Request(&root)
	assert.Equal(t, Hyperlink(testURLStringOf("users/{user}")), root.UserURL)
}
