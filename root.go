package octokat

import (
	"github.com/octokit/octokat/hyper"
)

type Root struct {
	UserSearchURL               hyper.Link `json:"user_search_url,omitempty"`
	UserRepositoriesURL         hyper.Link `json:"user_repositories_url,omitempty"`
	UserOrganizationsURL        hyper.Link `json:"user_organizations_url,omitempty"`
	UserURL                     hyper.Link `json:"user_url,omitempty"`
	TeamURL                     hyper.Link `json:"team_url,omitempty"`
	StarredGistsURL             hyper.Link `json:"starred_gists_url,omitempty"`
	StarredURL                  hyper.Link `json:"starred_url,omitempty"`
	CurrentUserRepositoriesURL  hyper.Link `json:"current_user_repositories_url,omitempty"`
	RepositorySearchURL         hyper.Link `json:"repository_search_url,omitempty"`
	RepositoryURL               hyper.Link `json:"repository_url,omitempty"`
	RateLimitURL                hyper.Link `json:"rate_limit_url,omitempty"`
	GistsURL                    hyper.Link `json:"gists_url,omitempty"`
	FollowingURL                hyper.Link `json:"following_url,omitempty"`
	FeedsURL                    hyper.Link `json:"feeds_url,omitempty"`
	EventsURL                   hyper.Link `json:"events_url,omitempty"`
	EmojisURL                   hyper.Link `json:"emojis_url,omitempty"`
	EmailsURL                   hyper.Link `json:"emails_url,omitempty"`
	AuthorizationsURL           hyper.Link `json:"authorizations_url,omitempty"`
	CurrentUserURL              hyper.Link `json:"current_user_url,omitempty"`
	HubURL                      hyper.Link `json:"hub_url,omitempty"`
	IssueSearchURL              hyper.Link `json:"issue_search_url,omitempty"`
	IssuesURL                   hyper.Link `json:"issues_url,omitempty"`
	KeysURL                     hyper.Link `json:"keys_url,omitempty"`
	NotificationsURL            hyper.Link `json:"notifications_url,omitempty"`
	OrganizationRepositoriesURL hyper.Link `json:"organization_repositories_url,omitempty"`
	OrganizationsURL            hyper.Link `json:"organization_url,omitempty"`
	PublicGistsURL              hyper.Link `json:"public_gists_url,omitempty"`
}

func (c *Client) Root(headers Headers) (root *Root, err error) {
	resp, e := c.Get("", headers)
	if e != nil {
		err = e
		return
	}
	if resp.HasError() {
		err = resp.Error
		return
	}

	err = resp.Data(&root)
	return
}
