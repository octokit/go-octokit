package octokat

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestCreateAccessToken(t *testing.T) {
	_, err := CreateAccessToken(&Params{})

	assert.NotEqual(t, nil, err)
	assert.Equal(t, "Missing fields: client_id, client_secret, code", err.Error())
}
