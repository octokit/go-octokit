package octokat

import (
	"net/url"
)

type Requester struct {
	client  *Client
	URL     *url.URL
	Headers Headers
}

func (r *Requester) Get(v interface{}) (resp *Response, err error) {
	resp, err = r.client.Get(r.URL, r.Headers)
	if !canUnmarshal(resp, err) {
		err = resp.Data(v)
	}

	return
}

func (r *Requester) Patch(params interface{}, v interface{}) (resp *Response, err error) {
	resp, err = r.client.Patch(r.URL, r.Headers, params)
	if !canUnmarshal(resp, err) {
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
