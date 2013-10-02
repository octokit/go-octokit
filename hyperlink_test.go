package octokat

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/jingweno/octokat", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("repository.json"))
	})

	link := Hyperlink{client: client, Rel: "repository", Href: testURLOf("repos/jingweno/octokat")}
	var repo Repository
	link.Get(&repo, nil)

	assert.Equal(t, 10575811, repo.ID)
	assert.Equal(t, "jingweno/octokat", repo.FullName)
}
