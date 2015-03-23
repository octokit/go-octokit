package octokit

// Organization is a representation of an organization on GitHub, containing
// all identifying information related to the specific organization.

var (
	OrgReposURL = Hyperlink("/orgs/{org}/repos{?type,page,per_page,sort}")
	OrgURL      = Hyperlink("/orgs/{org}")
	YourOrgsURL = Hyperlink("/user/orgs")
	UserOrgsURL = Hyperlink("/users/{username}/orgs")
)

type Organization struct {
	AvatarURL        string     `json:"avatar_url,omitempty"`
	PublicMembersURL Hyperlink  `json:"public_member_url,omitempty"`
	MembersURL       Hyperlink  `json:"members_url,omitempty"`
	EventsURL        Hyperlink  `json:"events_url,omitempty"`
	ReposURL         Hyperlink  `json:"repos_url,omitempty"`
	URL              string     `json:"url,omitempty"`
	ID               int        `json:"id,omitempty"`
	Login            string     `json:"login,omitempty"`
	Description      string     `json:"description, omitempty"`
	Name             string     `json:"name, omitempty"`
	Company          string     `json:"company, omitempty"`
	Blog             string     `json:"blog, omitempty"`
	Location         string     `json:"location, omitempty"`
	Email            string     `json:"email, omitempty"`
	PublicRepos      int        `json:"public_repos,omitempty"`
	PublicGists      int        `json:"public_gists,omitempty"`
	Followers        int        `json:"followers,omitempty"`
	Followering      int        `json:"following,omitempty"`
	HTMLURL          string     `json:"html_url,omitempty"`
	CreatedAt        *time.Time `json:"created_at,omitempty"`
	UpdatedAt        *time.Time `json:"updated_at,omitempty"`
	Type             string     `json:"type,omitempty"`
}
