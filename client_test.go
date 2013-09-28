package octokat

import (
	"bytes"
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
