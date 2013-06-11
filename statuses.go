package octokat

import (
	"fmt"
	"time"
)

type Status struct {
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	State       string    `json:"state"`
	TargetUrl   string    `json:"target_url"`
	Description string    `json:"description"`
	Id          int       `json:"id"`
	Url         string    `json:"url"`
	Creator     User      `json:"creator"`
}

func (c *Client) Statuses(repo Repo, sha string) ([]Status, error) {
	path := fmt.Sprintf("repos/%s/statuses/%s", repo, sha)
	var statuses []Status
	err := c.jsonGet(path, nil, &statuses)
	if err != nil {
		return nil, err
	}

	return statuses, nil
}
