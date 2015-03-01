package octokit

import (
	"net/url"
)

var (
	FollowerUrl        = Hyperlink("users/{user}/followers")
	CurrentFollowerUrl = Hyperlink("user/followers")
)

// Create a FollowersService with the base url.URL
func (c *Client) Followers(url *url.URL) (followers *FollowersService) {
	followers = &FollowersService{client: c, URL: url}
	return
}

// A service to return user followers
type FollowersService struct {
	client *Client
	URL    *url.URL
}

// Get a list of followers for the current user
func (f *FollowersService) All() (followers []User, result *Result) {
	result = f.client.get(f.URL, &followers)
	return
}
