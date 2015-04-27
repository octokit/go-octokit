package octokit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicAuth(t *testing.T) {
	basicAuth := BasicAuth{Login: "jingweno", Password: "password"}
	assert.EqualValues(t, "Basic amluZ3dlbm86cGFzc3dvcmQ=", basicAuth.String())
}

func TestNetrcAuth(t *testing.T) {
	netrcAuth := NetrcAuth{NetrcPath: "../fixtures/example.netrc"}
	assert.EqualValues(t, "Basic Y2F0c2J5OnY1UDZmZ2huN19hX2Zha2VfY29kZV9QR3VlbHZiRmF4QlBrVWcxaWI=", netrcAuth.String())
}

func TestTokenAuth(t *testing.T) {
	tokenAuth := TokenAuth{AccessToken: "token"}
	assert.EqualValues(t, "token token", tokenAuth.String())
}
