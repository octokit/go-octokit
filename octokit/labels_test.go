package octokit

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestLabelsService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octokit/go-octokit/labels", "labels", nil)

	labels, result := client.Labels().All(nil, M{"owner": "octokit", "repo": "go-octokit"})

	assert.False(t, result.HasError())

	assert.Equal(t, 2, len(labels))

  assert.Equal(t, "https://api.github.com/repos/octokit/go-octokit/labels/bug", labels[0].URL)
  assert.Equal(t, "bug", labels[0].Name)
	assert.Equal(t, "fc2929", labels[0].Color)

  assert.Equal(t, "https://api.github.com/repos/octokit/go-octokit/labels/duplicate", labels[1].URL)
  assert.Equal(t, "duplicate", labels[1].Name)
	assert.Equal(t, "cccccc", labels[1].Color)
}
