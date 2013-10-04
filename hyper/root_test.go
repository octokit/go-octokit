package hyper

import (
	"encoding/json"
	"github.com/bmizerany/assert"
	"testing"
)

func TestRoot_parseRelNameFromURL(t *testing.T) {
	assert.Equal(t, "repository", parseRelNameFromURL("repository_url"))
	assert.Equal(t, "public_gists", parseRelNameFromURL("public_gists_url"))
}

func TestRoot_Rel(t *testing.T) {
	links := make(map[string]Link)
	links["user"] = Link("https://api.github.com/users/{user}")
	root := Root{links: links}

	assert.Equal(t, links["user"], *root.Rel("user"))
	assert.T(t, root.Rel("not_exist") == nil)
}

func TestRoot_AddRel(t *testing.T) {
	root := Root{}
	root.AddRel("user", Link("https://api.github.com/users/{user}"))
	assert.Equal(t, 1, len(root.links))
}

func TestRoot_Marshal(t *testing.T) {
	links := make(map[string]Link)
	links["user"] = Link("https://api.github.com/users/{user}")
	root := Root{links: links}
	b, _ := json.Marshal(root)
	assert.Equal(t, `{"user_url":"https://api.github.com/users/{user}"}`, string(b))

	root = Root{}
	b, _ = json.Marshal(root)
	assert.Equal(t, `{}`, string(b))
}

func TestRoot_Unmarshal(t *testing.T) {
	data := `{"user_url": "https://api.github.com/users/{user}"}`
	var root Root
	json.Unmarshal([]byte(data), &root)

	assert.Equal(t, 1, len(root.links))
}
