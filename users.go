package octokat

import (
	"strconv"
	"time"
)

type User struct {
	Login             string     `json:"login,omitempty"`
	ID                int        `json:"id,omitempty"`
	AvatarURL         string     `json:"avatar_url,omitempty"`
	GravatarID        string     `json:"gravatar_id,omitempty"`
	URL               string     `json:"url,omitempty"`
	Name              string     `json:"name,omitempty"`
	Company           string     `json:"company,omitempty"`
	Blog              string     `json:"blog,omitempty"`
	Location          string     `json:"location,omitempty"`
	Email             string     `json:"email,omitempty"`
	Hireable          bool       `json:"hireable,omitempty"`
	Bio               string     `json:"bio,omitempty"`
	PublicRepos       int        `json:"public_repos,omitempty"`
	PublicGists       int        `json:"public_gists,omitempty"`
	Followers         int        `json:"followers,omitempty"`
	Following         int        `json:"following,omitempty"`
	HTMLURL           string     `json:"html_url,omitempty"`
	CreatedAt         *time.Time `json:"created_at,omitempty"`
	UpdatedAt         *time.Time `json:"updated_at,omitempty"`
	Type              string     `json:"type,omitempty"`
	FollowingURL      Hyperlink  `json:"following_url,omitempty"`
	FollowersURL      Hyperlink  `json:"followers_url,omitempty"`
	GistsURL          Hyperlink  `json:"gists_url,omitempty"`
	StarredURL        Hyperlink  `json:"starred_url,omitempty"`
	SubscriptionsURL  Hyperlink  `json:"subscriptions_url,omitempty"`
	OrganizationsURL  Hyperlink  `json:"organizations_url,omitempty"`
	ReposURL          Hyperlink  `json:"repos_url,omitempty"`
	EventsURL         Hyperlink  `json:"events_url,omitempty"`
	ReceivedEventsURL Hyperlink  `json:"received_events_url,omitempty"`
}

type userRequester struct {
	client *Client
	login  string
}

func (u *userRequester) Request(v interface{}) (resp *Response, err error) {
	var root Root
	resp, err = u.client.Root().Request(&root)
	if hasError(resp, err) {
		return
	}

	var link Hyperlink
	if u.login == "" {
		link = root.CurrentUserURL
	} else {
		link = root.UserURL
	}

	userURL, e := link.Expand(M{"user": u.login})
	if e != nil {
		err = e
		return
	}

	resp, err = u.client.Get(userURL, nil)
	if !hasError(resp, err) {
		err = resp.Data(v)
	}

	return
}

func (c *Client) User(login string) (req Requester) {
	return &userRequester{client: c, login: login}
}

type updateUserRequester struct {
	client *Client
	params interface{}
}

func (u *updateUserRequester) Request(v interface{}) (resp *Response, err error) {
	var root Root
	resp, err = u.client.Root().Request(&root)
	if hasError(resp, err) {
		return
	}

	url, _ := root.CurrentUserURL.Expand(nil)
	resp, err = u.client.Patch(url, nil, u.params)
	if !hasError(resp, err) {
		err = resp.Data(v)
	}

	return
}

func (c *Client) UpdateUser(params interface{}) Requester {
	return &updateUserRequester{client: c, params: params}
}

type allUsersRequester struct {
	client *Client
	since  int
}

func (a *allUsersRequester) Request(v interface{}) (resp *Response, err error) {
	var root Root
	resp, err = a.client.Root().Request(&root)
	if hasError(resp, err) {
		return
	}

	url, e := root.UserURL.Expand(M{"user": ""})
	if e != nil {
		err = e
		return
	}

	if a.since > 0 {
		q := url.Query()
		q.Set("since", strconv.Itoa(a.since))
		url.RawQuery = q.Encode()
	}

	resp, err = a.client.Get(url, nil)
	if !hasError(resp, err) {
		err = resp.Data(v)
	}

	return
}

func (c *Client) AllUsers(since int) Requester {
	return &allUsersRequester{client: c, since: since}
}
