package octokat

import (
	"bytes"
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestGetResponse(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", MediaType)
		testHeader(t, r, "User-Agent", UserAgent)
		testHeader(t, r, "Content-Type", DefaultContentType)
		respondWith(w, "ok")
	})

	resp := client.Get(testURLOf("foo"), nil)
	assert.Equal(t, "ok", string(resp.RawBody))

	// path doesn't exist
	resp = client.Get(testURLOf("bar"), nil)
	assert.T(t, resp.Error != nil)
}

func TestGet(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", MediaType)
		testHeader(t, r, "User-Agent", UserAgent)
		testHeader(t, r, "Content-Type", DefaultContentType)
		respondWith(w, "ok")
	})

	client.get("foo", nil)
}

func TestPost(t *testing.T) {
	setup()
	defer tearDown()

	content := "content"
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", MediaType)
		testHeader(t, r, "User-Agent", UserAgent)
		testHeader(t, r, "Content-Type", DefaultContentType)
		testBody(t, r, content)
		respondWith(w, "ok")
	})

	client.post("foo", nil, bytes.NewBufferString(content))
}

func TestBuildURL(t *testing.T) {
	url, _ := client.buildURL("https://api.github.com")
	assert.Equal(t, "https://api.github.com", url.String())

	url, _ = client.buildURL("repos")
	assert.Equal(t, testURLOf("repos"), url.String())
}
