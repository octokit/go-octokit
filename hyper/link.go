package hyper

import (
	"github.com/jtacoma/uritemplates"
)

type M map[string]interface{}

type Link string

func (l *Link) Expand(m M) (string, error) {
	template, err := uritemplates.Parse(string(*l))
	if err != nil {
		return "", err
	}

	expanded, err := template.Expand(m)
	if err != nil {
		return "", err
	}

	return expanded, nil
}
