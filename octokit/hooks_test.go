package octokit

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHooksService_OrganizationOne(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/orgs/octokat/hooks/1", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("hook.json"))
	})

	url, err := OrganizationHookURL.Expand(M{"org": "octokat", "id": 1})
	assert.Equal(t, nil, err)

	hook, result := client.Hooks(url).One()
	fmt.Println(result)

	assert.T(t, !result.HasError())
	assert.Equal(t, uint(1), hook.ID)
	assert.Equal(t, "web", hook.Name)
	assert.Equal(t, "https://api.github.com/orgs/octocat/hooks/1", hook.URL)
	assert.T(t, hook.Active)
	assert.Equal(t, []string{"push", "pull_request"}, hook.Events)
	assert.Equal(t, "http://example.com", hook.Config.URL)
	assert.Equal(t, "json", hook.Config.ContentType)
}

func TestHooksService_OrganizationAll(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/orgs/octokat/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		respondWithJSON(w, loadFixture("hooks.json"))
	})

	url, err := OrganizationHooksURL.Expand(M{"org": "octokat"})
	assert.Equal(t, nil, err)

	hooks, result := client.Hooks(url).All()

	assert.T(t, !result.HasError())
	assert.Equal(t, 1, len(hooks))

	hook := hooks[0]

	assert.Equal(t, uint(1), hook.ID)
	assert.Equal(t, "web", hook.Name)
	assert.Equal(t, "https://api.github.com/orgs/octocat/hooks/1", hook.URL)
	assert.T(t, hook.Active)
	assert.Equal(t, []string{"push", "pull_request"}, hook.Events)
	assert.Equal(t, "http://example.com", hook.Config.URL)
	assert.Equal(t, "json", hook.Config.ContentType)
}

func TestHooksService_OrganizationCreate(t *testing.T) {
	setup()
	defer tearDown()

	params := Hook{}
	params.Name = "test-hook"
	params.Active = true
	params.Events = []string{"membership"}
	params.Config = &HookConfig{
		URL:         "https://test-example.com",
		ContentType: "json",
	}

	mux.HandleFunc("/orgs/octocat/hooks", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")

		var hookParams Hook
		json.NewDecoder(r.Body).Decode(&hookParams)
		assert.Equal(t, params.Name, hookParams.Name)
		assert.Equal(t, params.Active, hookParams.Active)
		assert.Equal(t, params.Events, hookParams.Events)
		assert.Equal(t, params.Config.URL, hookParams.Config.URL)
		assert.Equal(t, params.Config.ContentType, hookParams.Config.ContentType)

		respondWithJSON(w, loadFixture("create_organization_hook.json"))
	})

	url, err := OrganizationHooksURL.Expand(M{"org": "octocat"})
	assert.Equal(t, nil, err)

	hook, result := client.Hooks(url).Create(params)
	fmt.Println(result)

	assert.T(t, !result.HasError())
	assert.Equal(t, uint(2), hook.ID)
	assert.Equal(t, "test-hook", hook.Name)
	assert.T(t, hook.Active)
	assert.Equal(t, "https://api.github.com/orgs/octocat/hooks/2", hook.URL)
	assert.Equal(t, []string{"membership"}, hook.Events)
	assert.Equal(t, "https://test-example.com", hook.Config.URL)
	assert.Equal(t, "json", hook.Config.ContentType)
}
