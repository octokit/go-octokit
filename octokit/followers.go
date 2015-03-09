package octokit

import (
	"net/url"
)

var (
	FollowerUrl              = Hyperlink("users/{user}/followers")
	CurrentFollowerUrl       = Hyperlink("user/followers")
	FollowingUrl             = Hyperlink("users/{user}/following")
	CurrentFollowingUrl      = Hyperlink("user/following")
	CheckFollowingUrl        = Hyperlink("users/{user}/following/{target}")
	CheckCurrentFollowingUrl = Hyperlink("user/following/{target}")
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

// Get a list of followers for the user
func (f *FollowersService) All() (followers []User, result *Result) {
	result = f.client.get(f.URL, &followers)
	return
}

// Checks if a user is following a target user
func (f *FollowersService) Check() (result *Result) {
	result = f.client.get(f.URL, nil)
	return
}
