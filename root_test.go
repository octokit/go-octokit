package octokat

import (
	"github.com/bmizerany/assert"
	"github.com/octokit/octokat/hyper"
	"net/http"
	"testing"
)

func TestClient_Root(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("root.json"))
	})

	root, _ := client.Root(nil)
	repoLink := root.Rel("repository")
	assert.Equal(t, hyper.Link("https://api.github.com/repos/{owner}/{repo}"), *repoLink)
}
