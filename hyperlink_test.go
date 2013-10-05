package octokat

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestHyperlink_Expand(t *testing.T) {
	link := Hyperlink("https://api.github.com/users/{user}")
	url, _ := link.Expand(M{"user": "jingweno"})
	assert.Equal(t, "https://api.github.com/users/jingweno", url)

	link = Hyperlink("https://api.github.com/user")
	url, _ = link.Expand(nil)
	assert.Equal(t, "https://api.github.com/user", url)

	url, _ = link.Expand(M{})
	assert.Equal(t, "https://api.github.com/user", url)
}
