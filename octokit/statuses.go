package octokit

import (
	"net/url"
	"time"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// StatusesURL is a template for accessing statuses, such as build state, with of a particular
// reference or hash in a particular repository for a particular owner that can be expanded
// to a full address.
var StatusesURL = Hyperlink("repos/{owner}/{repo}/statuses/{ref}")

// Statuses creates a StatusesService with a base url
func (c *Client) Statuses(url *url.URL) (statuses *StatusesService) {
	statuses = &StatusesService{client: c, URL: url}
	return
}

// StatusesService is a service providing access to status from a particular url
type StatusesService struct {
	client *Client
	URL    *url.URL
}

// All gets a list of all the statuses associated with the url of the service
func (s *StatusesService) All() (statuses []Status, result *Result) {
	result = s.client.get(s.URL, &statuses)
	return
}

// Status represents a state marked from an external service regarding the
// current state of a commit, including success, failure, error or pending
type Status struct {
	*hypermedia.HALResource

	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	State       string    `json:"state,omitempty"`
	TargetURL   string    `json:"target_url,omitempty"`
	Description string    `json:"description,omitempty"`
	ID          int       `json:"id,omitempty"`
	URL         string    `json:"url,omitempty"`
	Creator     User      `json:"creator,omitempty"`
}
