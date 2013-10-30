package octokit

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestCurrentUserService_Get(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("user.json"))
	})

	currentUser, err := client.CurrentUser(nil, nil)
	assert.Equal(t, nil, err)

	user, result := currentUser.Get()

	assert.T(t, !result.HasError())
	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}

func TestCurrentUserService_Update(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")
		testBody(t, r, "{\"email\":\"jingweno@gmail.com\"}\n")
		respondWithJSON(w, loadFixture("user.json"))
	})

	currentUser, err := client.CurrentUser(nil, nil)
	assert.Equal(t, nil, err)

	userToUpdate := User{Email: "jingweno@gmail.com"}
	user, result := currentUser.Update(userToUpdate)

	assert.T(t, !result.HasError())
	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}
