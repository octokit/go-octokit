package octokat

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestPullRequest(t *testing.T) {
	c := NewClient().WithToken(os.Getenv("GITHUB_TOKEN"))
	repo := Repo{"octokat", "jingweno"}

	pr, err := c.PullRequest(repo, "1")

	assert.Equal(t, nil, err)
	assert.Equal(t, "ezbercih", pr.User.Login)
	assert.Equal(t, "ezbercih:patch-1", pr.Head.Label)
	assert.Equal(t, "jingweno:master", pr.Base.Label)
}
