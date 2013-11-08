package octokit

import (
	"github.com/lostisland/go-sawyer/hypermedia"
	"net/url"
	"time"
)

var (
	StatusesURL = Hyperlink("/repos/{owner}/{repo}/statuses/{ref}")
)

// Create a StatusesService with the base Hyperlink and the params M to expand the Hyperlink
// If no Hyperlink is passed in, it will use StatusesURL.
func (c *Client) Statuses(link *Hyperlink, m M) (statuses *StatusesService, err error) {
	if link == nil {
		link = &StatusesURL
	}

	url, err := link.Expand(m)
	if err != nil {
		return
	}

	statuses = &StatusesService{client: c, URL: url}
	return
}

type StatusesService struct {
	client *Client
	URL    *url.URL
}

func (s *StatusesService) GetAll() (statuses []Status, result *Result) {
	result = s.client.Get(s.URL, &statuses)
	return
}

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
