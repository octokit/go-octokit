package octokit

import (
	"net/url"
	"time"
)

const (
	RepositoryHooksURL   = Hyperlink("repos/{owner}/{repo}/hooks")
	RepositoryHookURL    = Hyperlink("repos/{owner}/{repo}/hooks/{id}")
	OrganizationHooksURL = Hyperlink("orgs/{org}/hooks")
	OrganizationHookURL  = Hyperlink("orgs/{org}/hooks/{id}")
)

// Create an HooksService with the base url.URL
func (c *Client) Hooks(url *url.URL) *HooksService {
	return &HooksService{client: c, URL: url}
}

// A service to return hook records
type HooksService struct {
	client *Client
	URL    *url.URL
}

func (u *HooksService) One() (hook *Hook, result *Result) {
	result = u.client.get(u.URL, &hook)
	return
}

func (u *HooksService) OnePreview() (hook *Hook, result *Result) {
	result = u.client.getPreview(u.URL, previewOrgHooksType, &hook)
	return
}

// Update a hook based on HookService#URL
func (u *HooksService) Create(params interface{}) (hook *Hook, result *Result) {
	result = u.client.post(u.URL, params, &hook)
	return
}

// Update a hook based on HookService#URL
func (u *HooksService) CreatePreview(params interface{}) (hook *Hook, result *Result) {
	result = u.client.postPreview(u.URL, previewOrgHooksType, params, &hook)
	return
}

// Update a hook based on HookService#URL
func (u *HooksService) Update(params interface{}) (hook *Hook, result *Result) {
	result = u.client.put(u.URL, params, &hook)
	return
}

// Get a list of hooks based on HookService#URL
func (u *HooksService) All() (hooks []Hook, result *Result) {
	result = u.client.get(u.URL, &hooks)
	return
}

func (u *HooksService) AllPreview() (hooks []Hook, result *Result) {
	result = u.client.getPreview(u.URL, previewOrgHooksType, &hooks)
	return
}

type Hook struct {
	ID        uint        `json:"id,omitempty"`
	URL       string      `json:"url,omitempty"`
	Name      string      `json:"name,omitempty"`
	Events    []string    `json:"events,omitempty"`
	Active    bool        `json:"active,omitempty"`
	Config    *HookConfig `json:"config,omitempty"`
	UpdatedAt time.Time   `json:"updated_at,omitempty"`
	CreatedAt time.Time   `json:"created_at,omitempty"`
}

type HookConfig struct {
	URL         string `json:"url,omitempty"`
	ContentType string `json:"content_type,omitempty"`
	Secret      string `json:"secret,omitempty"`
	InsecureSSL string `json:"insecure_ssl,omitempty"`
}
