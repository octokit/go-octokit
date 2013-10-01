package octokat

import (
	"fmt"
	"time"
)

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

type ReleaseParams struct {
	TagName         string `json:"tag_name,omitempty"`
	TargetCommitish string `json:"target_commitish,omitempty"`
	Name            string `json:"name,omitempty"`
	Body            string `json:"body,omitempty"`
	Draft           bool   `json:"draft,omitempty"`
	Prerelease      bool   `json:"prerelease,omitempty"`
}

// Create a release
//
// See http://developer.github.com/v3/repos/releases/#create-a-release
func (c *Client) CreateRelease(repo Repo, options *Options) (release *Release, err error) {
	path := fmt.Sprintf("repos/%s/releases", repo)
	options = addPreviewMediaType(options)
	err = c.jsonPost(path, options, &release)
	return
}

func addPreviewMediaType(options *Options) *Options {
	if options == nil {
		options = &Options{}
	}

	if options.Headers == nil {
		options.Headers = Headers{}
	}

	if options.Headers["Accept"] == "" {
		options.Headers["Accept"] = PreviewMediaType
	}

	return options
}
