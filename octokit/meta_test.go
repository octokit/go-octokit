package octokit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeta(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/meta", "meta", nil)
	info, result := client.Meta(&MetaURL)
	fmt.Println(info, result)

	assert.False(t, result.HasError())

	assert.True(t, info.VerifiablePasswordAuthentication)
}
