package octokit

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestStatuses(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/jingweno/gh/statuses/740211b9c6cd8e526a7124fe2b33115602fbc637", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("statuses.json"))
	})

	repo := Repo{UserName: "jingweno", Name: "gh"}
	sha := "740211b9c6cd8e526a7124fe2b33115602fbc637"
	statuses, _ := client.Statuses(repo, sha, nil)
	assert.Equal(t, 2, len(statuses))

	firstStatus := statuses[0]
	assert.Equal(t, "pending", firstStatus.State)
	assert.Equal(t, "The Travis CI build is in progress", firstStatus.Description)
	assert.Equal(t, "https://travis-ci.org/jingweno/gh/builds/11911500", firstStatus.TargetURL)
}
