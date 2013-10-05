package octokat

type Root struct {
	UserSearchURL               Hyperlink `json:"user_search_url,omitempty"`
	UserRepositoriesURL         Hyperlink `json:"user_repositories_url,omitempty"`
	UserOrganizationsURL        Hyperlink `json:"user_organizations_url,omitempty"`
	UserURL                     Hyperlink `json:"user_url,omitempty"`
	TeamURL                     Hyperlink `json:"team_url,omitempty"`
	StarredGistsURL             Hyperlink `json:"starred_gists_url,omitempty"`
	StarredURL                  Hyperlink `json:"starred_url,omitempty"`
	CurrentUserRepositoriesURL  Hyperlink `json:"current_user_repositories_url,omitempty"`
	RepositorySearchURL         Hyperlink `json:"repository_search_url,omitempty"`
	RepositoryURL               Hyperlink `json:"repository_url,omitempty"`
	RateLimitURL                Hyperlink `json:"rate_limit_url,omitempty"`
	GistsURL                    Hyperlink `json:"gists_url,omitempty"`
	FollowingURL                Hyperlink `json:"following_url,omitempty"`
	FeedsURL                    Hyperlink `json:"feeds_url,omitempty"`
	EventsURL                   Hyperlink `json:"events_url,omitempty"`
	EmojisURL                   Hyperlink `json:"emojis_url,omitempty"`
	EmailsURL                   Hyperlink `json:"emails_url,omitempty"`
	AuthorizationsURL           Hyperlink `json:"authorizations_url,omitempty"`
	CurrentUserURL              Hyperlink `json:"current_user_url,omitempty"`
	HubURL                      Hyperlink `json:"hub_url,omitempty"`
	IssueSearchURL              Hyperlink `json:"issue_search_url,omitempty"`
	IssuesURL                   Hyperlink `json:"issues_url,omitempty"`
	KeysURL                     Hyperlink `json:"keys_url,omitempty"`
	NotificationsURL            Hyperlink `json:"notifications_url,omitempty"`
	OrganizationRepositoriesURL Hyperlink `json:"organization_repositories_url,omitempty"`
	OrganizationsURL            Hyperlink `json:"organization_url,omitempty"`
	PublicGistsURL              Hyperlink `json:"public_gists_url,omitempty"`
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
