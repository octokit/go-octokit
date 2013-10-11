package octokit

import (
	"bytes"
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

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

	resp, _ := client.Get(testURLOf("foo"), nil)
	assert.Equal(t, 200, resp.StatusCode)

	// path doesn't exist
	_, err := client.Get(testURLOf("bar"), nil)
	assert.T(t, err != nil)
}

func TestPatch(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testHeader(t, r, "Accept", MediaType)
		testHeader(t, r, "User-Agent", UserAgent)
		testHeader(t, r, "Content-Type", DefaultContentType)
		testBody(t, r, `{"foo":"bar"}`)
		respondWith(w, "ok")
	})

	m := make(map[string]string)
	m["foo"] = "bar"
	resp, _ := client.Patch(testURLOf("foo"), nil, m)
	assert.Equal(t, 200, resp.StatusCode)

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testHeader(t, r, "Accept", MediaType)
		testHeader(t, r, "User-Agent", UserAgent)
		testHeader(t, r, "Content-Type", DefaultContentType)
		respondWith(w, "ok")
	})

	resp, _ = client.Patch(testURLOf("bar"), nil, nil)
	assert.Equal(t, 200, resp.StatusCode)

	// path doesn't exist
	_, err := client.Patch(testURLOf("baz"), nil, m)
	assert.T(t, err != nil)
}

func TestDeprecatedGet(t *testing.T) {
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

func TestDeprecatedPost(t *testing.T) {
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

func TestJSONPost(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", MediaType)
		testHeader(t, r, "User-Agent", UserAgent)
		testHeader(t, r, "Content-Type", DefaultContentType)
		testBody(t, r, "")
		respondWith(w, `{"ok": "foo"}`)
	})

	m := make(map[string]interface{})
	client.jsonPost("foo", nil, &m)
}

func TestBuildURL(t *testing.T) {
	url, _ := client.buildURL("https://api.github.com")
	assert.Equal(t, "https://api.github.com", url.String())

	url, _ = client.buildURL("repos")
	assert.Equal(t, testURLOf("repos"), url)
}
