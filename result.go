package octokit

type pageable struct {
	NextPage  *Hyperlink
	LastPage  *Hyperlink
	FirstPage *Hyperlink
	PrevPage  *Hyperlink
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
