package octokit

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestParsePage(t *testing.T) {
	link := `<https://api.github.com/user/repos?page=3&per_page=100>; rel="next", <https://api.github.com/user/repos?page=50&per_page=100>; rel="last"`
	page := parsePage(link)

	assert.Equal(t, "https://api.github.com/user/repos?page=3&per_page=100", page.NextPage.String())
	assert.Equal(t, "https://api.github.com/user/repos?page=50&per_page=100", page.LastPage.String())
}
