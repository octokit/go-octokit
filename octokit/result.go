package octokit

import (
	"net/url"
	"strings"
)

type Page struct {
	NextPage  *url.URL
	LastPage  *url.URL
	FirstPage *url.URL
	PrevPage  *url.URL
}

type Result struct {
	Response *Response
	Err      error
	Page
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
	var page Page
	if resp != nil {
		page = parsePage(resp.Header.Get("Link"))
	}

	return &Result{Response: resp, Page: page, Err: err}
}

func parsePage(link string) Page {
	page := Page{}

	if len(link) > 0 {
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
					page.NextPage = url
				case `rel="prev"`:
					page.PrevPage = url
				case `rel="first"`:
					page.FirstPage = url
				case `rel="last"`:
					page.LastPage = url
				}
			}
		}
	}

	return page
}
