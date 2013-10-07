package octokat

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestOrganizations(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users/jingweno/orgs", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("organizations.json"))
	})

	orgs, _ := client.Organizations("jingweno", nil)
	assert.Equal(t, 1, len(orgs))

	firstOrg := orgs[0]
	assert.Equal(t, 1388703, firstOrg.ID)
	assert.Equal(t, "acl-services", firstOrg.Login)
	assert.Equal(t, "https://api.github.com/orgs/acl-services", firstOrg.URL)
	assert.Equal(t, "https://api.github.com/orgs/acl-services/repos", firstOrg.ReposURL)
}
