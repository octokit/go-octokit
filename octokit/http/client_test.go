package http

import (
	"github.com/bmizerany/assert"
	"net/http"
	"testing"
)

func TestSuccessfulGet(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		head := w.Header()
		head.Set("Content-Type", "application/json")
		respondWith(w, `{"login": "octokit"}`)
	})

	req, err := client.NewRequest("foo")
	assert.Equal(t, nil, err)

	var output map[string]interface{}
	resp, err := req.Get(&output)
	assert.Equal(t, nil, err)
	assert.T(t, !resp.IsError())
	assert.Equal(t, "octokit", output["login"])
}

func TestGetWithoutDecoder(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		head := w.Header()
		head.Set("Content-Type", "application/booya+booya")
		respondWith(w, `{"login": "octokit"}`)
	})

	req, err := client.NewRequest("foo")
	assert.Equal(t, nil, err)

	var output map[string]interface{}
	_, err = req.Get(output)
	assert.NotEqual(t, nil, err)
}

func TestGetResponseError(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		head := w.Header()
		head.Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		respondWith(w, `{"message": "not found"}`)
	})

	req, err := client.NewRequest("foo")
	assert.Equal(t, nil, err)

	var output map[string]interface{}
	resp, err := req.Get(output)
	assert.Equal(t, nil, err)
	assert.T(t, resp.IsError())
	respErr := resp.Error
	assert.Equal(t, "not found", respErr.Message)
}

func TestSuccessfulPost(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testHeader(t, r, "Content-Type", "application/json")
		testBody(t, r, "{\"input\":\"bar\"}\n")
		head := w.Header()
		head.Set("Content-Type", "application/json")
		respondWith(w, `{"login": "octokit"}`)
	})

	req, err := client.NewRequest("foo")
	assert.Equal(t, nil, err)

	input := map[string]interface{}{"input": "bar"}
	var output map[string]interface{}
	resp, err := req.Post(input, &output)
	assert.Equal(t, nil, err)
	assert.T(t, !resp.IsError())
	assert.Equal(t, "octokit", output["login"])
}