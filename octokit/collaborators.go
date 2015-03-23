package octokit

// CollaboratorsURL is the template for accessing the collaborators
// of a particular repository.
var (
	CollaboratorsURL = Hyperlink(
		"repos/{owner}/{repo}/collaborators{/username}")
)

// Collaborators creates a CollaboratorsService with a base url
func (c *Client) Collaborators() (repos *CollaboratorsService) {
	repos = &CollaboratorsService{client: c}
	return
}

// CollaboratorsService is a service providing access to a repositories'
// collaborators
type CollaboratorsService struct {
	client *Client
}

// All lists all the collaborating users on the given repository
func (r *CollaboratorsService) All(uri *Hyperlink, params M) (users []User,
	result *Result) {
	if uri == nil {
		uri = &CollaboratorsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return make([]User, 0), &Result{Err: err}
	}
	result = r.client.get(url, &users)
	return
}

// One  gets one of the collaborating user on the given repository
func (r *CollaboratorsService) One(uri *Hyperlink, params M) (user *User,
	result *Result) {
	if uri == nil {
		uri = &CollaboratorsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}
	result = r.client.get(url, &user)
	return
}
