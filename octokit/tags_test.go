package octokit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestTagsService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/jingweno/gh/git/refs/tags", "tags", nil)

	url, err := TagsURL.Expand(M{"owner": "jingweno", "repo": "gh"})
	assert.NoError(t, err)

	tags, result := client.Tags(url).All()
	assert.False(t, result.HasError())
	assert.Len(t, tags, 2)

	firstTag := tags[0]
	assert.Equal(t, "refs/tags/v2", firstTag.Ref)
	assert.Equal(t, "https://api.github.com/repos/jingweno/gh/git/refs/tags/v2", firstTag.URL)

	firstTagCommit := firstTag.Commit
	assert.Equal(t, "6cb80cb09fd9f624a64d85438157955751a9ac70", firstTagCommit.Sha)
	assert.Equal(t, "https://api.github.com/repos/jingweno/gh/git/commits/6cb80cb09fd9f624a64d85438157955751a9ac70", firstTagCommit.URL)

	secondTag := tags[1]
	assert.Equal(t, "refs/tags/v1", secondTag.Ref)
	assert.Equal(t, "https://api.github.com/repos/jingweno/gh/git/refs/tags/v1", secondTag.URL)

	secondTagCommit := secondTag.Commit
	assert.Equal(t, "2fffba7fe19690e038314d17a117d6b87979c89f", secondTagCommit.Sha)
	assert.Equal(t, "https://api.github.com/repos/jingweno/gh/git/commits/2fffba7fe19690e038314d17a117d6b87979c89f", secondTagCommit.URL)

}