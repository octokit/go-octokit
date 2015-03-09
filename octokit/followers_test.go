package octokit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFollowersService_AllFollowers(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users/obsc/followers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/users/obsc/followers?page=2"), testURLOf("/users/obsc/followers?page=3"))
		header.Set("Link", link)

		respondWithJSON(w, loadFixture("followers.json"))
	})

	url, _ := FollowerUrl.Expand(M{"user": "obsc"})
	followers, result := client.Followers(url).All()

	assert.False(t, result.HasError())

	validateUser(t, followers)

	assert.Equal(t, testURLStringOf("/users/obsc/followers?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/users/obsc/followers?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestFollowersService_AllFollowersCurrent(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/followers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/user/followers?page=2"), testURLOf("/user/followers?page=3"))
		header.Set("Link", link)

		respondWithJSON(w, loadFixture("followers.json"))
	})

	url, _ := CurrentFollowerUrl.Expand(nil)
	followers, result := client.Followers(url).All()

	assert.False(t, result.HasError())

	validateUser(t, followers)

	assert.Equal(t, testURLStringOf("/user/followers?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/user/followers?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestFollowersService_AllFollowing(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users/obsc/following", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/users/obsc/following?page=2"), testURLOf("/users/obsc/following?page=3"))
		header.Set("Link", link)

		respondWithJSON(w, loadFixture("followers.json"))
	})

	url, _ := FollowingUrl.Expand(M{"user": "obsc"})
	allFollowing, result := client.Followers(url).All()

	assert.False(t, result.HasError())

	validateUser(t, allFollowing)

	assert.Equal(t, testURLStringOf("/users/obsc/following?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/users/obsc/following?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestFollowersService_AllFollowingCurrent(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/following", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/user/following?page=2"), testURLOf("/user/following?page=3"))
		header.Set("Link", link)

		respondWithJSON(w, loadFixture("followers.json"))
	})

	url, _ := CurrentFollowingUrl.Expand(nil)
	allFollowing, result := client.Followers(url).All()

	assert.False(t, result.HasError())

	validateUser(t, allFollowing)

	assert.Equal(t, testURLStringOf("/user/following?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/user/following?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestFollowersService_CheckFollowing(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/users/harrisonzhao/following/obsc", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	url, _ := CheckFollowingUrl.Expand(M{"user": "harrisonzhao", "target": "obsc"})
	result := client.Followers(url).Check()

	assert.False(t, result.HasError())
	assert.Equal(t, 204, result.Response.StatusCode)
}

func TestFollowersService_CheckCurrentFollowing(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/following/obsc", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	url, _ := CheckCurrentFollowingUrl.Expand(M{"target": "obsc"})
	result := client.Followers(url).Check()

	assert.False(t, result.HasError())
	assert.Equal(t, 204, result.Response.StatusCode)
}

func validateUser(t *testing.T, followers []User) {
	assert.Len(t, followers, 1)
	first := followers[0]

	assert.Equal(t, "harrisonzhao", first.Login)
	assert.Equal(t, 5186533, first.ID)
	assert.Equal(t, "https://avatars.githubusercontent.com/u/5186533?v=3", first.AvatarURL)
	assert.Equal(t, "", first.GravatarID)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao", first.URL)
	assert.Equal(t, "https://github.com/harrisonzhao", first.HTMLURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/followers", first.FollowersURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/following{/other_user}", first.FollowingURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/gists{/gist_id}", first.GistsURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/starred{/owner}{/repo}", first.StarredURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/subscriptions", first.SubscriptionsURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/orgs", first.OrganizationsURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/repos", first.ReposURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/events{/privacy}", first.EventsURL)
	assert.Equal(t, "https://api.github.com/users/harrisonzhao/received_events", first.ReceivedEventsURL)
	assert.Equal(t, "User", first.Type)
	assert.Equal(t, false, first.SiteAdmin)
}

func validateNextPage(t *testing.T, result *Result) {
	nextPageURL, err := result.NextPage.Expand(nil)
	assert.NoError(t, err)

	followers, result := client.Followers(nextPageURL).All()
	assert.False(t, result.HasError())
	assert.Len(t, followers, 1)
}
