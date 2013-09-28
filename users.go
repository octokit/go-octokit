package octokat

import (
	"fmt"
	"time"
)

type User struct {
	Login       string    `json:"login"`
	ID          int       `json:"id"`
	AvatarURL   string    `json:"avatar_url"`
	GravatarID  string    `json:"gravatar_id"`
	URL         string    `json:"url"`
	Name        string    `json:"name"`
	Company     string    `json:"company"`
	Blog        string    `json:"blog"`
	Location    string    `json:"location"`
	Email       string    `json:"email"`
	Hireable    bool      `json:"hireable"`
	Bio         string    `json:"bio"`
	PublicRepos int       `json:"public_repos"`
	PublicGists int       `json:"jsonpublic_gists"`
	Followers   int       `json:"followers"`
	Following   int       `json:"following"`
	HTMLURL     string    `json:"html_url"`
	CreatedAt   time.Time `json:"created_at"`
	Type        string    `json:"type"`
}

func (c *Client) User(login string) (user *User, err error) {
	var path string
	if login == "" {
		path = "user"
	} else {
		path = fmt.Sprintf("users/%s", login)
	}

	err = c.jsonGet(path, nil, &user)
	return
}
