package octokit

import (
	"fmt"
	"github.com/bmizerany/assert"
	"testing"
)

func TestResponse_HasError(t *testing.T) {
	resp := Response{}
	assert.T(t, !resp.HasError())

	resp = Response{Error: fmt.Errorf("an error")}
	assert.T(t, resp.HasError())
}

func TestResponse_Data(t *testing.T) {
	resp := Response{RawBody: []byte(loadFixture("user.json"))}
	var user User
	err := resp.Data(&user)

	assert.T(t, err == nil)
	assert.Equal(t, 169064, user.ID)
}
