package octokat

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestRoot(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("root.json"))
	})

	root, _ := client.Root(nil)
	assert.T(t, root.client != nil)
	assert.Equal(t, 27, len(root.links))

	repoLink := root.Rel("repository")
	assert.T(t, repoLink.client != nil)
	assert.Equal(t, "repository", repoLink.Rel)
	assert.Equal(t, "https://api.github.com/repos/{owner}/{repo}", repoLink.Href)
}

func TestParseRelNameFromURL(t *testing.T) {
	assert.Equal(t, "repository", parseRelNameFromURL("repository_url"))
	assert.Equal(t, "public_gists", parseRelNameFromURL("public_gists_url"))
}
