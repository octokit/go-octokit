package hyper

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestExpand(t *testing.T) {
	link := Link("https://api.github.com/users/{user}")
	url, _ := link.Expand(M{"user": "jingweno"})
	assert.Equal(t, "https://api.github.com/users/jingweno", url)
}
