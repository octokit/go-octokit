package octokit

import (
	"github.com/lostisland/go-sawyer"
	"net/url"
)

type M map[string]interface{}

type Hyperlink string

// TODO: find out a way to not wrapping sawyer.Hyperlink like this
func (l *Hyperlink) Expand(m M) (u *url.URL, err error) {
	link := sawyer.Hyperlink(string(*l))
	sawyerM := sawyer.M{}
	for k, v := range m {
		sawyerM[k] = v
	}

	u, err = link.Expand(sawyerM)
	return
}
