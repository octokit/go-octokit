package octokit

import (
	"encoding/json"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizationsService_One(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/authorizations/1", "authorization", nil)

	url, err := AuthorizationsURL.Expand(M{"id": 1})
	assert.NoError(t, err)

	auth, result := client.Authorizations(url).One()

	assert.False(t, result.HasError())
	assert.EqualValues(t, 1, auth.ID)
	assert.EqualValues(t, "https://api.github.com/authorizations/1", auth.URL)
	assert.EqualValues(t, "456", auth.Token)
	assert.EqualValues(t, "", auth.Note)
	assert.EqualValues(t, "", auth.NoteURL)
	assert.EqualValues(t, "2012-11-16 01:05:51 +0000 UTC", auth.CreatedAt.String())
	assert.EqualValues(t, "2013-08-21 03:29:51 +0000 UTC", auth.UpdatedAt.String())

	app := App{ClientID: "123", URL: "http://localhost:8080", Name: "Test"}
	assert.EqualValues(t, app, auth.App)

	assert.Len(t, auth.Scopes, 2)
	scopes := []string{"repo", "user"}
	assert.True(t, reflect.DeepEqual(auth.Scopes, scopes))
}

func TestAuthorizationsService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/authorizations", "authorizations", nil)

	url, err := AuthorizationsURL.Expand(nil)
	assert.NoError(t, err)

	auths, result := client.Authorizations(url).All()
	assert.False(t, result.HasError())

	firstAuth := auths[0]
	assert.EqualValues(t, 1, firstAuth.ID)
	assert.EqualValues(t, "https://api.github.com/authorizations/1", firstAuth.URL)
	assert.EqualValues(t, "456", firstAuth.Token)
	assert.EqualValues(t, "", firstAuth.Note)
	assert.EqualValues(t, "", firstAuth.NoteURL)
	assert.EqualValues(t, "2012-11-16 01:05:51 +0000 UTC", firstAuth.CreatedAt.String())
	assert.EqualValues(t, "2013-08-21 03:29:51 +0000 UTC", firstAuth.UpdatedAt.String())

	app := App{ClientID: "123", URL: "http://localhost:8080", Name: "Test"}
	assert.EqualValues(t, app, firstAuth.App)

	assert.Len(t, firstAuth.Scopes, 2)
	scopes := []string{"repo", "user"}
	assert.True(t, reflect.DeepEqual(firstAuth.Scopes, scopes))
}

func TestAuthorizationsService_Create(t *testing.T) {
	setup()
	defer tearDown()

	params := AuthorizationParams{Scopes: []string{"public_repo"}}

	mux.HandleFunc("/authorizations", func(w http.ResponseWriter, r *http.Request) {
		var authParams AuthorizationParams
		json.NewDecoder(r.Body).Decode(&authParams)
		assert.True(t, reflect.DeepEqual(authParams, params))

		testMethod(t, r, "POST")
		respondWithJSON(w, loadFixture("create_authorization.json"))
	})

	url, err := AuthorizationsURL.Expand(nil)
	assert.NoError(t, err)

	auth, _ := client.Authorizations(url).Create(params)

	assert.EqualValues(t, 3844190, auth.ID)
	assert.EqualValues(t, "https://api.github.com/authorizations/3844190", auth.URL)
	assert.EqualValues(t, "123", auth.Token)
	assert.EqualValues(t, "", auth.Note)
	assert.EqualValues(t, "", auth.NoteURL)
	assert.EqualValues(t, "2013-09-28 18:44:39 +0000 UTC", auth.CreatedAt.String())
	assert.EqualValues(t, "2013-09-28 18:44:39 +0000 UTC", auth.UpdatedAt.String())

	app := App{ClientID: "00000000000000000000", URL: "http://developer.github.com/v3/oauth/#oauth-authorizations-api", Name: "GitHub API"}
	assert.EqualValues(t, app, auth.App)

	assert.Len(t, auth.Scopes, 1)
	scopes := []string{"public_repo"}
	assert.True(t, reflect.DeepEqual(auth.Scopes, scopes))
}
