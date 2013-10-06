package octokat

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestClient_User(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("user.json"))
	})

	var user User
	requester := client.User("")
	requester.Request(&user)

	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, 25, user.PublicGists)

	mux.HandleFunc("/users/jingweno", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("user.json"))
	})

	requester = client.User("jingweno")
	requester.Request(&user)

	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, Hyperlink("https://api.github.com/users/jingweno/repos"), user.ReposURL)
}

func TestUser_UpdateUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testBody(t, r, `{"name":"name","email":"email"}`)
		respondWith(w, loadFixture("user.json"))
	})

	var userToUpdate = User{
		Name:  "name",
		Email: "email",
	}

	var user *User
	requester := client.UpdateUser(userToUpdate)
	requester.Request(&user)
	assert.Equal(t, 169064, user.ID)
}

func TestUser_AllUsers(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		assert.Equal(t, "/users/?since=1", r.URL.String())
		respondWith(w, loadFixture("users.json"))
	})

	var users []User
	requester := client.AllUsers(1)
	requester.Request(&users)
	assert.Equal(t, 1, len(users))
}
