package octokit

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGistsService_One(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/gists/a6bea192debdbec0d4ab", "gist", nil)

	url, _ := GistsURL.Expand(M{"gist_id": "a6bea192debdbec0d4ab"})
	gist, result := client.Gists(url).One()

	assert.False(t, result.HasError())
	assert.EqualValues(t, "a6bea192debdbec0d4ab", gist.ID)
	assert.Len(t, gist.Files, 1)

	file := gist.Files["grep_cellar"]
	assert.EqualValues(t, "grep_cellar", file.FileName)
	assert.EqualValues(t, "text/plain", file.Type)
	assert.EqualValues(t, "", file.Language)
	assert.EqualValues(t, "https://gist.githubusercontent.com/jingweno/a6bea192debdbec0d4ab/raw/80757419d2bd4cfddf7c6be24308eca11b3c330e/grep_cellar", file.RawURL)
	assert.EqualValues(t, 8107, file.Size)
	assert.EqualValues(t, false, file.Truncated)
}

func TestGistsService_Raw(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/gists/a6bea192debdbec0d4ab", "gist", nil)

	mux.HandleFunc("/jingweno/a6bea192debdbec0d4ab/raw/80757419d2bd4cfddf7c6be24308eca11b3c330e/grep_cellar", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		assert.EqualValues(t, "gist.githubusercontent.com", r.Host)
		testHeader(t, r, "Accept", textMediaType)
		respondWith(w, "hello")
	})

	url, _ := GistsURL.Expand(M{"gist_id": "a6bea192debdbec0d4ab"})
	body, result := client.Gists(url).Raw()

	assert.False(t, result.HasError())
	content, err := ioutil.ReadAll(body)
	assert.NoError(t, err)
	assert.EqualValues(t, "hello", string(content))
}
