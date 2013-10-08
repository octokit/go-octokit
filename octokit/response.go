package octokit

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	*http.Response
}

func (resp *Response) Data(v interface{}) error {
	return json.NewDecoder(resp.Body).Decode(v)
}
