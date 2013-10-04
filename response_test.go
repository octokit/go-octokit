package octokat

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestResponse_Data(t *testing.T) {
	resp := Response{RawBody: []byte(loadFixture("user.json"))}
	var user User
	err := resp.Data(&user)

	assert.T(t, err == nil)
	assert.Equal(t, 169064, user.ID)
}
