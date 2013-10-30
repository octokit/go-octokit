package octokit

import (
	"github.com/jtacoma/uritemplates"
	"net/url"
)

// TODO: use sawyer.Hyperlink

type M map[string]interface{}

type Hyperlink string

func (l *Hyperlink) Expand(m M) (u *url.URL, err error) {
	template, e := uritemplates.Parse(string(*l))
	if e != nil {
		err = e
		return
	}

	expanded, e := template.Expand(m)
	if e != nil {
		err = e
		return
	}

	u, err = url.Parse(expanded)
	return
}
