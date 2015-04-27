package octokit

import (
	"fmt"
	"net/http"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUsersService_GetCurrentUser(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/user", "user", nil)

	url, _ := CurrentUserURL.Expand(nil)
	user, result := client.Users(url).One()

	assert.False(t, result.HasError())
	assert.EqualValues(t, 169064, user.ID)
	assert.EqualValues(t, "jingweno", user.Login)
	assert.EqualValues(t, "jingweno@gmail.com", user.Email)
	assert.EqualValues(t, "User", user.Type)
	assert.EqualValues(t, 17, user.Following)
	assert.EqualValues(t, 28, user.Followers)
	assert.EqualValues(t, 90, user.PublicRepos)
	assert.EqualValues(t, false, user.SiteAdmin)
	assert.EqualValues(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}

func TestUsersService_UpdateCurrentUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, "{\"email\":\"jingweno@gmail.com\"}\n")
		respondWithJSON(w, loadFixture("user.json"))
	})

	url, _ := CurrentUserURL.Expand(nil)
	userToUpdate := User{Email: "jingweno@gmail.com"}
	user, result := client.Users(url).Update(userToUpdate)

	assert.False(t, result.HasError())
	assert.EqualValues(t, 169064, user.ID)
	assert.EqualValues(t, "jingweno", user.Login)
	assert.EqualValues(t, "jingweno@gmail.com", user.Email)
	assert.EqualValues(t, "User", user.Type)
	assert.EqualValues(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}

func TestUsersService_GetUser(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/users/jingweno", "user", nil)

	url, err := UserURL.Expand(M{"user": "jingweno"})
	assert.NoError(t, err)
	user, result := client.Users(url).One()

	assert.False(t, result.HasError())
	assert.EqualValues(t, 169064, user.ID)
	assert.EqualValues(t, "jingweno", user.Login)
	assert.EqualValues(t, "jingweno@gmail.com", user.Email)
	assert.EqualValues(t, "User", user.Type)
	assert.EqualValues(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}

func TestUsersService_All(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		rr := regexp.MustCompile(`users\?since=\d+`)
		assert.True(t, rr.MatchString(r.URL.String()), "Regexp should match users?since=\\d+")

		header := w.Header()
		link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="first"`, testURLOf("users?since=135"), testURLOf("users{?since}"))
		header.Set("Link", link)
		respondWithJSON(w, loadFixture("users.json"))
	})

	url, err := UserURL.Expand(M{"since": 1})
	assert.NoError(t, err)

	q := url.Query()
	q.Set("since", "1")
	url.RawQuery = q.Encode()
	allUsers, result := client.Users(url).All()

	assert.False(t, result.HasError())
	assert.Len(t, allUsers, 1)
	assert.EqualValues(t, testURLStringOf("users?since=135"), string(*result.NextPage))

	nextPageURL, err := result.NextPage.Expand(nil)
	assert.NoError(t, err)

	allUsers, result = client.Users(nextPageURL).All()
	assert.False(t, result.HasError())
	assert.Len(t, allUsers, 1)
}
