package octokit

import (
	"net/http"
	"testing"

	"github.com/bmizerany/assert"
)

func TestCommitsService_One(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/go-octokit/commits/4351fb69b8d5ed075e9cd844e67ad2114b335c82", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("commit.json"))
	})

	url, err := CommitsURL.Expand(M{
		"owner": "octokit",
		"repo":  "go-octokit",
		"sha":   "4351fb69b8d5ed075e9cd844e67ad2114b335c82",
	})
	assert.Equal(t, nil, err)
	commit, result := client.Commits(url).One()

	assert.T(t, !result.HasError())
	assert.Equal(t, "4351fb69b8d5ed075e9cd844e67ad2114b335c82", commit.Sha)
	assert.Equal(t, "https://api.github.com/repos/octokit/go-octokit/commits/4351fb69b8d5ed075e9cd844e67ad2114b335c82", commit.URL)

	files := commit.Files
	assert.Equal(t, 35, len(files))
}
