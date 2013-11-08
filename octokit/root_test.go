package octokit

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestRootService_Get(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("root.json"))
	})

	rootService, err := client.Root(nil)
	assert.Equal(t, nil, err)

	root, result := rootService.Get()
	assert.T(t, !result.HasError())
	assert.Equal(t, "https://api.github.com/users/{user}", string(root.UserURL))
}
