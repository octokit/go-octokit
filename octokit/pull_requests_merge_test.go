package octokit

import (
	"testing"

	"net/url"

	"bytes"
	"io"
	"net/http"

	"github.com/stretchr/testify/assert"
)

func TestPullRequestsMergeURL(t *testing.T) {
	url, err := PullRequestsMergeURL.Expand(M{"owner": "jingweno", "repo": "gh", "number": 11})
	assert.NoError(t, err)
	assert.Equal(t, "repos/jingweno/gh/pulls/11/merge", url.String())
}

func TestPullRequestsMerge(t *testing.T) {
	url := &url.URL{}
	c := &Client{}
	service := c.PullRequestsMerge(url)
	assert.NotNil(t, service)
	assert.Equal(t, service.client, c)
	assert.Equal(t, service.URL, url)
}

func TestFoo(t *testing.T) {
	commitMessage := "requestCommitMessage"
	sha := "requestSha"
	url, err := PullRequestsMergeURL.Expand(M{"owner": "jingweno", "repo": "gh", "number": 11})
	assert.NoError(t, err)

	transport := new(mockTransport)
	transport.err = nil
	transport.resp = &http.Response{
		Header: http.Header{},
		Body: nopCloser{bytes.NewBufferString(`{
  "sha": "6dcb09b5b57875f334f61aebed695e2e4193db5e",
  "merged": true,
  "message": "Pull Request successfully merged"
}`)},
	}
	transport.resp.Header.Add("Content-Type", "application/json")

	c := NewClientWith(
		server.URL,
		userAgent,
		BasicAuth{
			Login:           "jingweno",
			Password:        "password",
			OneTimePassword: "OTP",
		},
		&http.Client{
			Transport: transport,
		},
	)

	service := c.PullRequestsMerge(url)
	response, result := service.Merge(&PullRequestsMergeRequest{
		CommitMessage: commitMessage,
		Sha:           sha,
	})
	assert.NotNil(t, result)
	assert.Nil(t, result.Err)
	assert.NotNil(t, response)
	assert.Equal(t, "6dcb09b5b57875f334f61aebed695e2e4193db5e", response.Sha)
	assert.Equal(t, true, response.Merged)
	assert.Equal(t, "Pull Request successfully merged", response.Message)
}

type mockTransport struct {
	req  *http.Request
	resp *http.Response
	err  error
}

func (m *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	m.req = req
	return m.resp, m.err
}

type nopCloser struct {
	io.Reader
}

func (nopCloser) Close() error {
	return nil
}
