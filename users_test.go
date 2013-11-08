package octokit

import (
	"fmt"
	"github.com/bmizerany/assert"
	"net/http"
	"regexp"
	"testing"
)

func TestUsersService_GetCurrentUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("user.json"))
	})

	users, err := client.Users(nil, nil)
	assert.Equal(t, nil, err)

	user, result := users.Get()

	assert.T(t, !result.HasError())
	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}

func TestUsersService_UpdateCurrentUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, "{\"email\":\"jingweno@gmail.com\"}\n")
		respondWithJSON(w, loadFixture("user.json"))
	})

	users, err := client.Users(nil, nil)
	assert.Equal(t, nil, err)

	userToUpdate := User{Email: "jingweno@gmail.com"}
	user, result := users.Update(userToUpdate)

	assert.T(t, !result.HasError())
	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}

func TestUsersService_GetUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users/jingweno", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("user.json"))
	})

	users, err := client.Users(&UserURL, M{"user": "jingweno"})
	assert.Equal(t, nil, err)

	user, result := users.Get()

	assert.T(t, !result.HasError())
	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}

func TestUsersService_GetAll(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		rr := regexp.MustCompile(`users\?since=\d+`)
		assert.Tf(t, rr.MatchString(r.URL.String()), "Regexp should match users?since=\\d+")

		header := w.Header()
		link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="first"`, testURLOf("users?since=135"), testURLOf("users{?since}"))
		header.Set("Link", link)
		respondWithJSON(w, loadFixture("users.json"))
	})

	users, err := client.Users(&UserURL, M{"since": 1})
	assert.Equal(t, nil, err)

	allUsers, result := users.GetAll()

	assert.T(t, !result.HasError())
	assert.Equal(t, 1, len(allUsers))
	assert.Equal(t, testURLStringOf("users?since=135"), string(*result.NextPage))

	users, err = client.Users(result.NextPage, nil)
	assert.Equal(t, nil, err)

	allUsers, result = users.GetAll()
	assert.T(t, !result.HasError())
	assert.Equal(t, 1, len(allUsers))
}
