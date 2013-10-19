package octokit

import (
	"net/url"
)

type pageable struct {
	NextPage  *url.URL
	LastPage  *url.URL
	FirstPage *url.URL
	PrevPage  *url.URL
}

type Result struct {
	Response *Response
	Err      error
	pageable
}

func (r *Result) HasError() bool {
	return r.Err != nil
}

func (r *Result) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}

	return ""
}

func newResult(resp *Response, err error) *Result {
	var pageable pageable
	if resp != nil {
		parser := paginationParser{header: resp.Header}
		pageable = parser.Parse()
	}

	return &Result{Response: resp, pageable: pageable, Err: err}
}
