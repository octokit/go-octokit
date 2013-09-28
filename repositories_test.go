package octokat

import (
	"encoding/json"
	"github.com/bmizerany/assert"
	"net/http"
	"reflect"
	"testing"
)

func TestCreateRepository(t *testing.T) {
	setup()
	defer tearDown()

	params := RepositoryParams{}
	params.Name = "Hello-World"
	params.Description = "This is your first repo"
	params.Homepage = "https://github.com"
	params.Private = false
	params.HasIssues = true
	params.HasWiki = true
	params.HasDownloads = true

	mux.HandleFunc("/user/repos", func(w http.ResponseWriter, r *http.Request) {
		var repoParams RepositoryParams
		json.NewDecoder(r.Body).Decode(&repoParams)
		assert.T(t, reflect.DeepEqual(repoParams, params))

		testMethod(t, r, "POST")
		respondWith(w, loadFixture("create_repository.json"))
	})

	options := Options{Params: params}
	repo, _ := client.CreateRepository("", &options)

	assert.Equal(t, 1296269, repo.ID)
	assert.Equal(t, "Hello-World", repo.Name)
	assert.Equal(t, "octocat/Hello-World", repo.FullName)
	assert.Equal(t, "This is your first repo", repo.Description)
	assert.T(t, !repo.Private)
	assert.T(t, repo.Fork)
	assert.Equal(t, "https://api.github.com/repos/octocat/Hello-World", repo.URL)
	assert.Equal(t, "https://github.com/octocat/Hello-World", repo.HTMLURL)
	assert.Equal(t, "https://github.com/octocat/Hello-World.git", repo.CloneURL)
	assert.Equal(t, "git://github.com/octocat/Hello-World.git", repo.GitURL)
	assert.Equal(t, "git@github.com:octocat/Hello-World.git", repo.SSHURL)
	assert.Equal(t, "master", repo.MasterBranch)
}

func TestFork(t *testing.T) {
	//c := NewClient().WithToken(os.Getenv("GITHUB_TOKEN"))
	//repo := Repo{"octokat", "jingweno"}

	//repository, err := c.Fork(repo, nil)
	//assert.Equal(t, nil, err)
	//assert.Equal(t, "octokat", repository.Name)
}
