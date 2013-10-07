package octokat

import (
	"fmt"
)

type Organization struct {
	AvatarURL        string `json:"avatar_url,omitempty"`
	PublicMembersURL string `json:"public_member_url,omitempty"`
	MembersURL       string `json:"members_url,omitempty"`
	EventsURL        string `json:"events_url,omitempty"`
	ReposURL         string `json:"repos_url,omitempty"`
	URL              string `json:"url,omitempty"`
	ID               int    `json:"id,omitempty"`
	Login            string `json:"login,omitempty"`
}

func (c *Client) Organizations(user string, options *Options) (orgs []Organization, err error) {
	var path string
	if user == "" {
		path = "user/orgs"
	} else {
		path = fmt.Sprintf("users/%s/orgs", user)
	}

	err = c.jsonGet(path, options, &orgs)
	return
}

func (c *Client) OrganizationRepositories(org string, options *Options) (repos []Repository, err error) {
	path := fmt.Sprintf("orgs/%s/repos", org)
	err = c.jsonGet(path, options, &repos)
	return
}
