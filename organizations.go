package octokat

import (
	"fmt"
)

type Organization User

func (c *Client) Organizations(user string, params *Params) ([]Organization, error) {
	var path string
	if user == "" {
		path = "user/orgs"
	} else {
		path = fmt.Sprintf("users/%s/orgs", user)
	}

	var orgs []Organization
	err := c.jsonGet(path, nil, &orgs)

	if err != nil {
		return nil, err
	}

	return orgs, err
}

func (c *Client) OrganizationRepositories(org string, params *Params) ([]Repository, error) {
	path := fmt.Sprintf("orgs/%s/repos", org)

	var repos []Repository
	err := c.jsonGet(path, nil, &repos)

	if err != nil {
		return nil, err
	}

	return repos, err
}
