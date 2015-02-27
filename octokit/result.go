package octokit

import (
	"github.com/jingweno/go-sawyer/mediaheader"
)

type pageable struct {
	NextPage  *Hyperlink
	LastPage  *Hyperlink
	FirstPage *Hyperlink
	PrevPage  *Hyperlink
}

// Result is a pageable set of data, with hyperlinks to the first, last,
// previous, and next pages, containing a response to some request and
// associated error, if any
type Result struct {
	Response *Response
	Err      error
	pageable
}

// HasError returns true if the error field of the Result is not nil; false
// otherwise
func (r *Result) HasError() bool {
	return r.Err != nil
}

// Error returns the string representation of the error if it exists; the
// empty string is returned otherwise
func (r *Result) Error() string {
	if r.Err != nil {
		return r.Err.Error()
	}

	return ""
}

func newResult(resp *Response, err error) *Result {
	pageable := pageable{}
	if resp != nil {
		fillPageable(&pageable, resp.MediaHeader)
	}

	return &Result{Response: resp, pageable: pageable, Err: err}
}

func fillPageable(pageable *pageable, header *mediaheader.MediaHeader) {
	if link, ok := header.Relations["next"]; ok {
		l := Hyperlink(link)
		pageable.NextPage = &l
	}

	if link, ok := header.Relations["prev"]; ok {
		l := Hyperlink(link)
		pageable.PrevPage = &l
	}

	if link, ok := header.Relations["first"]; ok {
		l := Hyperlink(link)
		pageable.FirstPage = &l
	}

	if link, ok := header.Relations["last"]; ok {
		l := Hyperlink(link)
		pageable.LastPage = &l
	}
}
