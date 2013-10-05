package octokat

import (
	"github.com/octokit/octokat/hyper"
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
	FollowingURL      hyper.Link `json:"following_url,omitempty"`
	FollowersURL      hyper.Link `json:"followers_url,omitempty"`
	GistsURL          hyper.Link `json:"gists_url,omitempty"`
	StarredURL        hyper.Link `json:"starred_url,omitempty"`
	SubscriptionsURL  hyper.Link `json:"subscriptions_url,omitempty"`
	OrganizationsURL  hyper.Link `json:"organizations_url,omitempty"`
	ReposURL          hyper.Link `json:"repos_url,omitempty"`
	EventsURL         hyper.Link `json:"events_url,omitempty"`
	ReceivedEventsURL hyper.Link `json:"received_events_url,omitempty"`
}

func (c *Client) User(login string, headers Headers) (user *User, err error) {
	root, e := c.Root(headers)
	if e != nil {
		err = e
		return
	}

	var link hyper.Link
	if login == "" {
		link = root.CurrentUserURL
	} else {
		link = root.UserURL
	}

	userURL, e := link.Expand(hyper.M{"user": login})
	if e != nil {
		err = e
		return
	}

	resp, e := c.Get(userURL, headers)
	if e != nil {
		err = e
		return
	}
	if resp.HasError() {
		err = resp.Error
		return
	}

	err = resp.Data(&user)
	return
}

func (c *Client) UpdateUser(params interface{}, headers Headers) (user *User, err error) {
	root, e := c.Root(headers)
	if e != nil {
		err = e
		return
	}

	resp, e := c.Patch(string(root.CurrentUserURL), headers, params)
	if e != nil {
		err = e
		return
	}
	if resp.HasError() {
		err = resp.Error
		return
	}

	err = resp.Data(&user)
	return
}
