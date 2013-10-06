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

func (c *Client) User(login string) (user *User, result *Result) {
	root, r := c.Root()
	if r.HasError() {
		result = r
		return
	}

	var link Hyperlink
	if login == "" {
		link = root.CurrentUserURL
	} else {
		link = root.UserURL
	}

	userURL, e := link.Expand(M{"user": login})
	if e != nil {
		result = newResult(nil, e)
		return
	}

	requester := c.Requester(userURL)
	result = requester.Get(&user)

	return
}

func (c *Client) UpdateUser(params interface{}) (user *User, result *Result) {
	root, r := c.Root()
	if r.HasError() {
		result = r
		return
	}

	url, _ := root.CurrentUserURL.Expand(nil)
	requester := c.Requester(url)
	result = requester.Patch(params, &user)

	return
}

type allUsersRequester struct {
	client *Client
	since  int
}

func (c *Client) AllUsers(since int) (users []User, result *Result) {
	root, r := c.Root()
	if r.HasError() {
		result = r
		return
	}

	url, e := root.UserURL.Expand(M{"user": ""})
	if e != nil {
		result = newResult(nil, e)
		return
	}

	if since > 0 {
		q := url.Query()
		q.Set("since", strconv.Itoa(since))
		url.RawQuery = q.Encode()
	}

	requester := c.Requester(url)
	result = requester.Get(&users)

	return
}
