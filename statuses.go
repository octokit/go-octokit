package octokat

import (
	"fmt"
	"time"
)

type Status struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	State       string    `json:"state"`
	TargetURL   string    `json:"target_url"`
	Description string    `json:"description"`
	ID          int       `json:"id"`
	URL         string    `json:"url"`
	Creator     User      `json:"creator"`
}

// List all statuses for a given commit
//
// See http://developer.github.com/v3/repos/statuses
func (c *Client) Statuses(repo Repo, sha string, options *Options) (statuses []Status, err error) {
	path := fmt.Sprintf("repos/%s/statuses/%s", repo, sha)
	err = c.jsonGet(path, options, &statuses)
	return
}
