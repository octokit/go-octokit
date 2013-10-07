package octokat

type Result struct {
	Response *Response
	Err      error
}

func (r *Result) HasError() bool {
	return r.Err != nil || (r.Response != nil && r.Response.HasError())
}

func (r *Result) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}

	if r.Response != nil && r.Response.HasError() {
		return r.Response.Error.Error()
	}

	return ""
}

func newResult(resp *Response, err error) *Result {
	return &Result{Response: resp, Err: err}
}
