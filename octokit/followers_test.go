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

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/users/obsc/followers?page=2"), testURLOf("/users/obsc/followers?page=3"))
	stubGet(t, "/users/obsc/followers", "followers", map[string]string{"Link": link})

	followers, result := client.Followers().All(&FollowerUrl, M{"user": "obsc"})
	assert.False(t, result.HasError())

	validateUser(t, followers)

	assert.Equal(t, testURLStringOf("/users/obsc/followers?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/users/obsc/followers?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestFollowersService_AllFollowersCurrent(t *testing.T) {
	setup()
	defer tearDown()

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/user/followers?page=2"), testURLOf("/user/followers?page=3"))
	stubGet(t, "/user/followers", "followers", map[string]string{"Link": link})

	followers, result := client.Followers().All(nil, nil)
	assert.False(t, result.HasError())

	validateUser(t, followers)

	assert.Equal(t, testURLStringOf("/user/followers?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/user/followers?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestFollowersService_AllFollowing(t *testing.T) {
	setup()
	defer tearDown()

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/users/obsc/following?page=2"), testURLOf("/users/obsc/following?page=3"))
	stubGet(t, "/users/obsc/following", "followers", map[string]string{"Link": link})

	allFollowing, result := client.Followers().All(&FollowingUrl, M{"user": "obsc"})
	assert.False(t, result.HasError())

	validateUser(t, allFollowing)

	assert.Equal(t, testURLStringOf("/users/obsc/following?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/users/obsc/following?page=3"), string(*result.LastPage))

	validateNextPage(t, result)
}

func TestFollowersService_AllFollowingCurrent(t *testing.T) {
	setup()
	defer tearDown()

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/user/following?page=2"), testURLOf("/user/following?page=3"))
	stubGet(t, "/user/following", "followers", map[string]string{"Link": link})

	allFollowing, result := client.Followers().All(&CurrentFollowingUrl, nil)
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

	success, result := client.Followers().Check(&FollowingUrl, M{"user": "harrisonzhao", "target": "obsc"})
	assert.False(t, result.HasError())
	assert.True(t, success)
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

	success, result := client.Followers().Check(nil, M{"target": "obsc"})
	assert.False(t, result.HasError())
	assert.True(t, success)
}

func TestFollowersService_FollowUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/following/obsc", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	success, result := client.Followers().Follow(nil, M{"target": "obsc"})
	assert.False(t, result.HasError())
	assert.True(t, success)

}

func TestFollowersService_UnfollowUser(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/following/obsc", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	success, result := client.Followers().Unfollow(nil, M{"target": "obsc"})
	assert.False(t, result.HasError())
	assert.True(t, success)
}

func TestFollowersService_Failure(t *testing.T) {
	setup()
	defer tearDown()

	url := Hyperlink("}")
	followers, result := client.Followers().All(&url, nil)
	assert.True(t, result.HasError())
	assert.Len(t, followers, 0)

	success, result := client.Followers().Check(&url, nil)
	assert.True(t, result.HasError())
	assert.False(t, success)

	success, result = client.Followers().Follow(&url, nil)
	assert.True(t, result.HasError())
	assert.False(t, success)

	success, result = client.Followers().Unfollow(&url, nil)
	assert.True(t, result.HasError())
	assert.False(t, success)
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
	followers, result := client.Followers().All(result.NextPage, nil)
	assert.False(t, result.HasError())
	assert.Len(t, followers, 1)
}
