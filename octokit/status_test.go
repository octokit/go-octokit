package octokit

import (
	"testing"

	"net/url"

	"github.com/stretchr/testify/assert"
)

func TestStatusUrl(t *testing.T) {
	sha := "740211b9c6cd8e526a7124fe2b33115602fbc637"
	url, err := StatusURL.Expand(M{"owner": "jingweno", "repo": "gh", "ref": sha})
	assert.NoError(t, err)
	assert.Equal(t, "repos/jingweno/gh/status/740211b9c6cd8e526a7124fe2b33115602fbc637", url.String())
}

func TestStatus(t *testing.T) {
	url := &url.URL{}
	c := &Client{}
	service := c.Status(url)
	assert.NotNil(t, service)
	assert.Equal(t, service.client, c)
	assert.Equal(t, service.URL, url)
}
