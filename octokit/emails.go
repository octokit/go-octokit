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

// Get a list of emails for the current user
func (e *EmailsService) All() (emails []Email, result *Result) {
	result = e.client.get(e.URL, &emails)
	return
}

// Adds a list of emails for the current user
func (e *EmailsService) Create(params interface{}) (emails []Email, result *Result) {
	result = e.client.post(e.URL, params, &emails)
	return
}

// Deletes a list of emails for the current user
func (e *EmailsService) Delete(params interface{}) (result *Result) {
	result = e.client.delete(e.URL, params, nil)
	return
}

type Email struct {
	*hypermedia.HALResource

	Email    string `json:"email,omitempty"`
	Verified bool   `json:"verified,omitempty"`
	Primary  bool   `json:"primary,omitempty"`
}
