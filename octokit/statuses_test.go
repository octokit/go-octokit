package octokit

import (
	"testing"
	"encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestStatusesService_All(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/jingweno/gh/statuses/740211b9c6cd8e526a7124fe2b33115602fbc637", "statuses", nil)

	sha := "740211b9c6cd8e526a7124fe2b33115602fbc637"
	url, err := StatusesURL.Expand(M{"owner": "jingweno", "repo": "gh", "ref": sha})
	assert.NoError(t, err)

	statuses, err := client.Statuses(url).All()

	assert.Len(t, statuses, 2)
	firstStatus := statuses[0]
	assert.Equal(t, "pending", firstStatus.State)
	assert.Equal(t, "The Travis CI build is in progress", firstStatus.Description)
	assert.Equal(t, "https://travis-ci.org/jingweno/gh/builds/11911500", firstStatus.TargetURL)
}

func TestStatusesService_Create(t *testing.T) {
	setup()
	defer tearDown()

	sha := "740211b9c6cd8e526a7124fe2b33115602fbc637"
	url, err := StatusesURL.Expand(M{"owner": "jingweno", "repo": "gh", "ref": sha})
	assert.NoError(t, err)

	params := Status{
		State:       "success",
		TargetURL:   "https://example.com/build/status",
		Description: "The build succeeded!",
		Context:     "continuous-integration/jenkins",
	}
	wantReqBody, _ := json.Marshal(params)
	stubPost(t, "/repos/jingweno/gh/statuses/740211b9c6cd8e526a7124fe2b33115602fbc637",
		"create_status", nil, string(wantReqBody)+"\n", nil)

	status, result := client.Statuses(url).Create(params)

	assert.False(t, result.HasError())
	assert.Equal(t, "success", status.State)
	assert.Equal(t, "https://example.com/build/status", status.TargetURL)
	assert.Equal(t, "The build succeeded!", status.Description)
	assert.Equal(t, "continuous-integration/jenkins", status.Context)
}
