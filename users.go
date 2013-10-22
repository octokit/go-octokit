package octokit

import (
	"github.com/lostisland/go-sawyer"
	"net/url"
	"time"
)

func (c *Client) Users(url *url.URL) *UsersService {
	return &UsersService{client: c, URL: url}
}

type UsersService struct {
	client *Client
	URL    *url.URL
}

func (u *UsersService) Get() (user *User, result *Result) {
	return
}

type User struct {
	*sawyer.HALResource

	Login             string           `json:"login,omitempty"`
	ID                int              `json:"id,omitempty"`
	AvatarURL         string           `json:"avatar_url,omitempty"`
	GravatarID        string           `json:"gravatar_id,omitempty"`
	URL               string           `json:"url,omitempty"`
	Name              string           `json:"name,omitempty"`
	Company           string           `json:"company,omitempty"`
	Blog              string           `json:"blog,omitempty"`
	Location          string           `json:"location,omitempty"`
	Email             string           `json:"email,omitempty"`
	Hireable          bool             `json:"hireable,omitempty"`
	Bio               string           `json:"bio,omitempty"`
	PublicRepos       int              `json:"public_repos,omitempty"`
	PublicGists       int              `json:"public_gists,omitempty"`
	Followers         int              `json:"followers,omitempty"`
	Following         int              `json:"following,omitempty"`
	HTMLURL           string           `json:"html_url,omitempty"`
	CreatedAt         *time.Time       `json:"created_at,omitempty"`
	UpdatedAt         *time.Time       `json:"updated_at,omitempty"`
	Type              string           `json:"type,omitempty"`
	FollowingURL      sawyer.Hyperlink `json:"following_url,omitempty"`
	FollowersURL      sawyer.Hyperlink `json:"followers_url,omitempty"`
	GistsURL          sawyer.Hyperlink `json:"gists_url,omitempty"`
	StarredURL        sawyer.Hyperlink `json:"starred_url,omitempty"`
	SubscriptionsURL  sawyer.Hyperlink `json:"subscriptions_url,omitempty"`
	OrganizationsURL  sawyer.Hyperlink `json:"organizations_url,omitempty"`
	ReposURL          sawyer.Hyperlink `json:"repos_url,omitempty"`
	EventsURL         sawyer.Hyperlink `json:"events_url,omitempty"`
	ReceivedEventsURL sawyer.Hyperlink `json:"received_events_url,omitempty"`
}
