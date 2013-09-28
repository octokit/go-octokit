package octokat

import (
	"fmt"
	"time"
)

type User struct {
	Login       string    `json:"login,omitempty"`
	ID          int       `json:"id,omitempty"`
	AvatarURL   string    `json:"avatar_url,omitempty"`
	GravatarID  string    `json:"gravatar_id,omitempty"`
	URL         string    `json:"url,omitempty"`
	Name        string    `json:"name,omitempty"`
	Company     string    `json:"company,omitempty"`
	Blog        string    `json:"blog,omitempty"`
	Location    string    `json:"location,omitempty"`
	Email       string    `json:"email,omitempty"`
	Hireable    bool      `json:"hireable,omitempty"`
	Bio         string    `json:"bio,omitempty"`
	PublicRepos int       `json:"public_repos,omitempty"`
	PublicGists int       `json:"jsonpublic_gists,omitempty"`
	Followers   int       `json:"followers,omitempty"`
	Following   int       `json:"following,omitempty"`
	HTMLURL     string    `json:"html_url,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	Type        string    `json:"type,omitempty"`
}

func (c *Client) User(login string, options *Options) (user *User, err error) {
	var path string
	if login == "" {
		path = "user"
	} else {
		path = fmt.Sprintf("users/%s", login)
	}

	err = c.jsonGet(path, options, &user)
	return
}
