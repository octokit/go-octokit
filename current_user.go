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

func (u *CurrentUserService) Get() (user *User, result *Result) {
	result = u.client.Get(u.URL, &user)
	return
}

func (u *CurrentUserService) Update(params interface{}) (user *User, result *Result) {
	result = u.client.Put(u.URL, params, &user)
	return
}
