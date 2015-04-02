package octokit

// Organization is a representation of an organization on GitHub, containing
// all identifying information related to the specific organization.

import (
	"time"
)

var (
	OrganizationReposURL = Hyperlink("/orgs/{org}/repos{?type,page,per_page,sort}")
	OrganizationURL      = Hyperlink("/orgs/{org}")
	YourOrganizationsURL = Hyperlink("/user/orgs")
	UserOrganizationsURL = Hyperlink("/users/{username}/orgs")
)

func (c *Client) Organization() (organization *OrganizationService) {
	organization = &OrganizationService{client: c}
	return
}

// A service to return organization information
type OrganizationService struct {
	client *Client
}

// Get the user search results based on OrganizationService#URL
func (g *OrganizationService) OrganizationRepos(uri *Hyperlink, params M) (
	repos []Repository, result *Result) {
	if uri == nil {
		uri = &OrganizationReposURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return make([]Repository, 0), &Result{Err: err}
	}
	result = g.client.get(url, &repos)
	return
}

// Get the user search results based on OrganizationService#URL
func (g *OrganizationService) OrganizationInfo(uri *Hyperlink, params M) (
	organization Organization, result *Result) {
	if uri == nil {
		uri = &OrganizationURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return Organization{}, &Result{Err: err}
	}
	result = g.client.get(url, &organization)
	return
}

// Get the issue search results based on OrganizationService#URL
func (g *OrganizationService) YourOrganizations(uri *Hyperlink, params M) (
	organizations []Organization, result *Result) {
	if uri == nil {
		uri = &YourOrganizationsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return make([]Organization, 0), &Result{Err: err}
	}
	result = g.client.get(url, &organizations)
	return
}

// Get the issue search results based on OrganizationService#URL
func (g *OrganizationService) UserOrganizations(uri *Hyperlink, params M) (
	organizations []Organization, result *Result) {
	if uri == nil {
		uri = &UserOrganizationsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return make([]Organization, 0), &Result{Err: err}
	}
	result = g.client.get(url, &organizations)
	return
}

type Organization struct {
	Description      string    `json:"description, omitempty"`
	AvatarURL        string    `json:"avatar_url,omitempty"`
	PublicMembersURL Hyperlink `json:"public_member_url,omitempty"`
	MembersURL       Hyperlink `json:"members_url,omitempty"`
	EventsURL        Hyperlink `json:"events_url,omitempty"`
	ReposURL         Hyperlink `json:"repos_url,omitempty"`
	URL              string    `json:"url,omitempty"`
	ID               int       `json:"id,omitempty"`
	Login            string    `json:"login,omitempty"`

	Name        string     `json:"name, omitempty"`
	Company     string     `json:"company, omitempty"`
	Blog        string     `json:"blog, omitempty"`
	Location    string     `json:"location, omitempty"`
	Email       string     `json:"email, omitempty"`
	PublicRepos int        `json:"public_repos,omitempty"`
	PublicGists int        `json:"public_gists,omitempty"`
	Followers   int        `json:"followers,omitempty"`
	Followering int        `json:"following,omitempty"`
	HTMLURL     string     `json:"html_url,omitempty"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	Type        string     `json:"type,omitempty"`
}
