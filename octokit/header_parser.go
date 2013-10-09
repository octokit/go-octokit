package octokit

import (
	"net/http"
	"net/url"
	"strings"
)

// TODO: need a full link header parser for http://tools.ietf.org/html/rfc5988
type paginationPraser struct {
	header http.Header
}

func (pp paginationPraser) Parse() pageable {
	link := pp.header.Get("Link")
	p := pageable{}
	if len(link) == 0 {
		return p
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

	return p
}
