package octokat

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestRepositories(t *testing.T) {
	c := NewClient().WithToken(os.Getenv("GITHUB_TOKEN"))
	repo := Repo{"octokat", "jingweno"}

	repository, err := c.Repository(repo)
	assert.Equal(t, nil, err)
	assert.Equal(t, "octokat", repository.Name)
	assert.Equal(t, "jingweno", repository.Owner.Login)
}
