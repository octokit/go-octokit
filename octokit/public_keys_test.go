package octokit

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPublicKeysService_AllKeys(t *testing.T) {
	setup()
	defer tearDown()

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/users/obsc/keys?page=2"), testURLOf("/users/obsc/keys?page=3"))
	stubGet(t, "/users/obsc/keys", "public_keys", map[string]string{"Link": link})

	keys, result := client.PublicKeys().All(&PublicKeyUrl, M{"user": "obsc"})
	assert.False(t, result.HasError())
	assert.Len(t, keys, 1)

	key := keys[0]
	assert.Equal(t, 8675080, key.Id)
	assert.Equal(t, "ssh-rsa AAA...", key.Key)

	assert.Equal(t, testURLStringOf("/users/obsc/keys?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/users/obsc/keys?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestPublicKeysService_AllKeysCurrent(t *testing.T) {
	setup()
	defer tearDown()

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/user/keys?page=2"), testURLOf("/user/keys?page=3"))
	stubGet(t, "/user/keys", "keys", map[string]string{"Link": link})

	keys, result := client.PublicKeys().All(nil, nil)
	assert.False(t, result.HasError())
	assert.Len(t, keys, 1)

	validateKey(t, keys[0])

	assert.Equal(t, testURLStringOf("/user/keys?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/user/keys?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestPublicKeysService_OneKey(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/user/keys/8675080", "key", nil)

	key, result := client.PublicKeys().One(nil, M{"id": 8675080})
	assert.False(t, result.HasError())

	validateKey(t, *key)
}

func TestPublicKeysService_Create(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/keys", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, "{\"key\":\"ssh-rsa AAA...\",\"title\":\"aKey\"}\n")

		respondWithJSON(w, loadFixture("key.json"))
	})

	params := Key{Title: "aKey", Key: "ssh-rsa AAA..."}
	key, result := client.PublicKeys().Create(nil, nil, params)
	assert.False(t, result.HasError())

	validateKey(t, *key)
}

func TestPublicKeysService_Delete(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/keys/8675080", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	success, result := client.PublicKeys().Delete(nil, M{"id": 8675080})
	assert.False(t, result.HasError())

	assert.True(t, success)
}

func validateKey(t *testing.T, key Key) {
	testTime, _ := time.Parse("2006-01-02T15:04:05Z", "2014-07-23T08:42:44Z")

	assert.Equal(t, 8675080, key.Id)
	assert.Equal(t, "ssh-rsa AAA...", key.Key)
	assert.Equal(t, "https://api.github.com/user/keys/8675080", key.URL)
	assert.Equal(t, "aKey", key.Title)
	assert.Equal(t, true, key.Verified)
	assert.Equal(t, &testTime, key.CreatedAt)
}

func validateNextPage(t *testing.T, result *Result) {
	keys, result := client.PublicKeys().All(result.NextPage, nil)
	assert.False(t, result.HasError())
	assert.Len(t, keys, 1)
}
