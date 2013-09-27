package octokat

import (
	"fmt"
	"time"
)

const (
	previewMediaType string = "application/vnd.github.manifold-preview"
)

type Asset struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Label         string    `json:"label,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	State         string    `json:"state,omitempty"`
	Size          int       `json:"size,omitempty"`
	DownloadCount int       `json:"download_count,omitempty"`
	URL           string    `json:"url,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type Release struct {
	ID             int       `json:"id,omitempty"`
	URL            string    `json:"url,omitempty"`
	HTMLURL        string    `json:"html_url,omitempty"`
	AssetsURL      string    `json:"assets_url,omitempty"`
	UploadURL      string    `json:"upload_url,omitempty"`
	TagName        string    `json:"tag_name,omitempty"`
	TargetCommitsh string    `json:"target_commitish,omitempty"`
	Name           string    `json:"name,omitempty"`
	Body           string    `json:"body,omitempty"`
	Draft          bool      `json:"draft,omitempty"`
	Prerelease     bool      `json:"prerelease,omitempty"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	PublishedAt    time.Time `json:"published_at,omitempty"`
	Assets         []Asset   `json:"assets,omitempty"`
}

// List releases for a repository
//
// http://developer.github.com/v3/repos/releases/#list-releases-for-a-repository
func (c *Client) Releases(repo Repo, options *Options) (releases []Release, err error) {
	path := fmt.Sprintf("repos/%s/releases", repo)

	options = addPreviewMediaType(options)
	err = c.jsonGet(path, options, &releases)

	return
}

func addPreviewMediaType(options *Options) *Options {
	if options == nil {
		options = &Options{}
	}

	if options.Headers == nil {
		options.Headers = make(map[string]string)
	}

	options.Headers["Accept"] = previewMediaType

	return options
}
