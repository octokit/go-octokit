package octokat

import (
	"fmt"
	"time"
)

type Release struct {
	ID             int       `json:"id"`
	HTMLURL        string    `json:"html_url"`
	AssetsURL      string    `json:"assets_url"`
	UploadURL      string    `json:"upload_url"`
	TagName        string    `json:"tag_name"`
	TargetCommitsh string    `json:"target_commitish"`
	Name           string    `json:"name"`
	Body           string    `json:"body"`
	Draft          bool      `json:"draft"`
	Prerelease     bool      `json:"prerelease"`
	CreatedAt      time.Time `json:"created_at"`
	PublishedAt    time.Time `json:"published_at"`
}

func (c *Client) Releases(repo Repo) ([]Release, error) {
	path := fmt.Sprintf("repos/%s/releases", repo)
	var releases []Release

	headers := make(map[string]string)
	headers["Accept"] = "application/vnd.github.manifold-preview"
	err := c.jsonGet(path, headers, &releases)
	if err != nil {
		return nil, err
	}

	return releases, nil
}
