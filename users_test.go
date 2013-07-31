package octokat

import (
	"github.com/bmizerany/assert"
	"os"
	"testing"
)

func TestUser(t *testing.T) {
	c := NewClient()
	user, err := c.User("jingweno")

	assert.Equal(t, nil, err)
	assert.Equal(t, "jingweno", user.Login)
}

func TestAuthenticatedUser(t *testing.T) {
	c := NewClient().WithToken(os.Getenv("GITHUB_TOKEN"))
	user, err := c.User("")

	assert.Equal(t, nil, err)
	assert.NotEqual(t, "", user.Login)
}
