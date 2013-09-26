package octokat

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestReleases(t *testing.T) {
	c := NewClient().WithToken(os.Getenv("GITHUB_TOKEN"))
	repo := Repo{"gh", "jingweno"}

	releases, err := c.Releases(repo)
	assert.Equal(t, nil, err)
	assert.T(t, len(releases) != 0)
}
