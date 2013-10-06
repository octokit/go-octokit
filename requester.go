package octokat

import (
	"net/url"
)

type Requester struct {
	client  *Client
	URL     *url.URL
	Headers Headers
}

func (r *Requester) Get(v interface{}) (result *Result) {
	resp, err := r.client.Get(r.URL, r.Headers)
	result = newResult(resp, err)
	if !result.HasError() {
		err = resp.Data(v)
	}

	return
}

func (r *Requester) Patch(params interface{}, v interface{}) (result *Result) {
	resp, err := r.client.Patch(r.URL, r.Headers, params)
	result = newResult(resp, err)
	if !result.HasError() {
		err = resp.Data(v)
	}

	return
}

func canUnmarshal(resp *Response, err error) bool {
	if err != nil && resp.HasError() {
		return true
	}

	return false
}
