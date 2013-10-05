package octokat

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestClient_Root(t *testing.T) {
	setup()
	defer tearDown()

	root, _ := client.Root(nil)
	assert.Equal(t, Hyperlink(testURLOf("users/{user}")), root.UserURL)
}
