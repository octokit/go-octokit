package octokit

import (
	"fmt"
	"time"
)

type Status struct {
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	State       string    `json:"state,omitempty"`
	TargetURL   string    `json:"target_url,omitempty"`
	Description string    `json:"description,omitempty"`
	ID          int       `json:"id,omitempty"`
	URL         string    `json:"url,omitempty"`
	Creator     User      `json:"creator,omitempty"`
}

// List all statuses for a given commit
//
// See http://developer.github.com/v3/repos/statuses
func (c *Client) Statuses(repo Repo, sha string, options *Options) (statuses []Status, err error) {
	path := fmt.Sprintf("repos/%s/statuses/%s", repo, sha)
	err = c.jsonGet(path, options, &statuses)
	return
}
