package octokit

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestUsersService_Get_FallbackURL(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users/jingweno", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		head := w.Header()
		head.Set("Content-Type", "application/json")
		respondWith(w, loadFixture("user.json"))
	})

	users, err := client.Users(nil, M{"user": "jingweno"})
	assert.Equal(t, nil, err)

	user, result := users.Get()

	assert.T(t, !result.HasError())
	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}

func TestUsersService_Get_PassInURL(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/my-users/jingweno", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		head := w.Header()
		head.Set("Content-Type", "application/json")
		respondWith(w, loadFixture("user.json"))
	})

	userLink := Hyperlink("my-users/{user}")
	users, err := client.Users(&userLink, M{"user": "jingweno"})
	assert.Equal(t, nil, err)

	user, result := users.Get()

	assert.T(t, !result.HasError())
	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, "https://api.github.com/users/jingweno/repos", string(user.ReposURL))
}
