package octokat

import (
	"github.com/bmizerany/assert"
	"testing"
)

func TestParamsPut(t *testing.T) {
	p := Params{"FOO": "BAR"}
	v := p.Put("BAZ", "BAR")

	assert.Equal(t, 2, p.Size())
	assert.Equal(t, nil, v)

	v = p.Put("FOO", "FOO")
	assert.Equal(t, 2, p.Size())
	assert.Equal(t, "BAR", v)
}

func TestParamsDelete(t *testing.T) {
	p := Params{"FOO": "BAR"}
	v := p.Delete("FOO")

	assert.Equal(t, 0, p.Size())
	assert.Equal(t, "BAR", v)

	v = p.Delete("BAR")
	assert.Equal(t, 0, p.Size())
	assert.Equal(t, nil, v)
}
