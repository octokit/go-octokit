package octokit

import (
	"net/url"
)

var DefaultCurrentUserHyperlink = Hyperlink("user")

type CurrentUserService struct {
	client *Client
	URL    *url.URL
}

func (c *Client) CurrentUser(link *Hyperlink, m M) (users *CurrentUserService, err error) {
	if link == nil {
		link = &DefaultCurrentUserHyperlink
	}

	url, err := link.Expand(m)
	if err != nil {
		return
	}

	users = &CurrentUserService{client: c, URL: url}
	return
}

func (u *CurrentUserService) Update(params interface{}) (user *User, result *Result) {
	req, err := u.client.NewRequest(u.URL.String())
	if err != nil {
		result = newResult(nil, err)
		return
	}

	resp, err := req.Put(params, &user)
	result = newResult(resp, err)

	return
}
