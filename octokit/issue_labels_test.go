package octokit

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIssueLabelsService_Add(t *testing.T) {
	setup()
	defer tearDown()

	input := []string{"newLabel", "anotherNewLabel"}
	wantReqBody, _ := json.Marshal(input)
	stubPost(t, "/repos/octokit/go-octokit/issues/33/labels", "issue_labels_added", nil, string(wantReqBody)+"\n", nil)

	labels, result := client.IssueLabels().Add(nil, M{"owner": "octokit", "repo": "go-octokit", "number": 33}, input)

	assert.False(t, result.HasError())

  assert.Equal(t, 2, len(labels))

  assert.Equal(t, "https://api.github.com/repos/octokit/go-octokit/labels/newLabel", labels[0].URL)
  assert.Equal(t, "newLabel", labels[0].Name)
	assert.Equal(t, "ffffff", labels[0].Color)

  assert.Equal(t, "https://api.github.com/repos/octokit/go-octokit/labels/anotherNewLabel", labels[1].URL)
  assert.Equal(t, "anotherNewLabel", labels[1].Name)
	assert.Equal(t, "000000", labels[1].Color)
}
