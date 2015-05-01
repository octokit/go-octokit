package octokit

// GitIgnoreURL is an address for accessing various templates to apply
// to a repository upon creation.
var GitIgnoreURL = Hyperlink("/gitignore/templates{/name}")

// GitIgnore creates a GitIgnoreService to access gitignore templates
func (c *Client) GitIgnore() *GitIgnoreService {
	return &GitIgnoreService{client: c}
}

// A service to return gitignore templates
type GitIgnoreService struct {
	client *Client
}

// All gets a list all the available templates
func (s *GitIgnoreService) All(uri *Hyperlink) (templates []string, result *Result) {
	if uri == nil {
		uri = &GitIgnoreURL
	}
	url, err := uri.Expand(nil)
	if err != nil {
		return make([]string, 0), &Result{Err: err}
	}
	result = s.client.get(url, &templates)
	return
}

// One gets a specific gitignore template based on the passed url
func (s *GitIgnoreService) One(uri *Hyperlink, params M) (template GitIgnoreTemplate, result *Result) {
	if uri == nil {
		uri = &GitIgnoreURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return GitIgnoreTemplate{}, &Result{Err: err}
	}
	result = s.client.get(url, &template)
	return
}

//GitIgnoreTemplate is a representation of a given template returned by the service
type GitIgnoreTemplate struct {
	Name   string `json:"name,omitempty"`
	Source string `json:"source,omitempty"`
}
