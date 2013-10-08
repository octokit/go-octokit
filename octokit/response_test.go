package octokit

import (
	"github.com/bmizerany/assert"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestResponse_Data(t *testing.T) {
	res := &http.Response{
		Request:    &http.Request{},
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(strings.NewReader(loadFixture("user.json"))),
	}
	resp := Response{Response: res}
	user := new(User)
	err := resp.Data(user)

	assert.T(t, err == nil)
	assert.Equal(t, 169064, user.ID)
}
