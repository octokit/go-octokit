package octokit

import (
	"net/url"

	"github.com/jingweno/go-sawyer/hypermedia"
)

var (
	EmailUrl = Hyperlink("user/emails")
)

// Create a EmailsService with the base url.URL
func (c *Client) Emails(url *url.URL) (emails *EmailsService) {
	emails = &EmailsService{client: c, URL: url}
	return
}

// A service to return user emails
type EmailsService struct {
	client *Client
	URL    *url.URL
}

// Get a list of emails based on EmailsService#URL
func (e *EmailsService) All() (emails []Email, result *Result) {
	result = e.client.get(e.URL, &emails)
	return
}

type Email struct {
	*hypermedia.HALResource

	Email    string `json:"email,omitempty"`
	Verified bool   `json:"verified,omitempty"`
	Primary  bool   `json:"primary,omitempty"`
}
