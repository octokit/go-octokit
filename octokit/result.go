package octokit

import (
	"net/url"
	"strings"
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
	var p pageable
	if resp != nil {
		p = parsePageable(resp.Header.Get("Link"))
	}

	return &Result{Response: resp, pageable: p, Err: err}
}

func parsePageable(link string) (p pageable) {
	p = pageable{}
	if len(link) == 0 {
		return
	}

	for _, l := range strings.Split(link, ",") {
		l = strings.TrimSpace(l)
		segments := strings.Split(l, ";")

		if len(segments) < 2 {
			continue
		}

		if !strings.HasPrefix(segments[0], "<") || !strings.HasSuffix(segments[0], ">") {
			continue
		}

		url, err := url.Parse(segments[0][1 : len(segments[0])-1])
		if err != nil {
			continue
		}

		for _, segment := range segments[1:] {
			switch strings.TrimSpace(segment) {
			case `rel="next"`:
				p.NextPage = url
			case `rel="prev"`:
				p.PrevPage = url
			case `rel="first"`:
				p.FirstPage = url
			case `rel="last"`:
				p.LastPage = url
			}
		}
	}

	return
}
