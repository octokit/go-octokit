package octokat

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("user.json"))
	})

	user, _ := client.User("", nil)

	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, 25, user.PublicGists)

	mux.HandleFunc("/users/jingweno", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWith(w, loadFixture("user.json"))
	})

	user, _ = client.User("jingweno", nil)

	assert.Equal(t, 169064, user.ID)
	assert.Equal(t, "jingweno", user.Login)
	assert.Equal(t, "jingweno@gmail.com", user.Email)
	assert.Equal(t, "User", user.Type)
	assert.Equal(t, "https://api.github.com/users/jingweno/repos", user.ReposURL)
}
