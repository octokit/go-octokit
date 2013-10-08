package octokit

type Result struct {
	Response *Response
	Err      error
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
	return &Result{Response: resp, Err: err}
}
