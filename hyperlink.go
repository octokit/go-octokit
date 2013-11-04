package octokit

import (
	"github.com/lostisland/go-sawyer"
	"net/url"
)

type M map[string]interface{}

type Hyperlink string

func (l *Hyperlink) Expand(m M) (u *url.URL, err error) {
	sawyerHyperlink := sawyer.Hyperlink(string(*l))
	u, err = sawyerHyperlink.Expand(m)
	return
}
