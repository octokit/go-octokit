package octokit

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGistsService_One(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/gists/a6bea192debdbec0d4ab", "gist", nil)

	gist, result := client.Gists().One(&GistsURL, M{"gist_id": "a6bea192debdbec0d4ab"})

	assert.False(t, result.HasError())
	assert.Equal(t, "a6bea192debdbec0d4ab", gist.ID)
	assert.Len(t, gist.Files, 1)

	file := gist.Files["grep_cellar"]
	assert.Equal(t, "grep_cellar", file.FileName)
	assert.Equal(t, "text/plain", file.Type)
	assert.Equal(t, "", file.Language)
	assert.Equal(t, "https://gist.githubusercontent.com/jingweno/a6bea192debdbec0d4ab/raw/80757419d2bd4cfddf7c6be24308eca11b3c330e/grep_cellar", file.RawURL)
	assert.Equal(t, 8107, file.Size)
	assert.Equal(t, false, file.Truncated)

	gistNil, resultNil := client.Gists().One(nil, M{"gist_id": "a6bea192debdbec0d4ab"})
	assert.False(t, resultNil.HasError())
	assert.Equal(t, gist, gistNil)

	//Error case
	var invalid = Hyperlink("{")
	gistErr, resultErr := client.Gists().One(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.Nil(t, gistErr)
}

func TestGistsService_Raw(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/gists/a6bea192debdbec0d4ab", "gist", nil)

	mux.HandleFunc("/jingweno/a6bea192debdbec0d4ab/raw/80757419d2bd4cfddf7c6be24308eca11b3c330e/grep_cellar", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		assert.Equal(t, "gist.githubusercontent.com", r.Host)
		testHeader(t, r, "Accept", textMediaType)
		respondWith(w, "hello")
	})

	body, result := client.Gists().Raw(&GistsURL, M{"gist_id": "a6bea192debdbec0d4ab"})

	assert.False(t, result.HasError())
	content, err := ioutil.ReadAll(body)
	assert.NoError(t, err)
	assert.Equal(t, "hello", string(content))
}

func TestGistsService_All(t *testing.T) {
	setup()
	defer tearDown()

	//Username URL
	stubGet(t, "/users/dannysperling/gists", "gists", nil)

	gists, result := client.Gists().All(&GistsUserURL, M{"username": "dannysperling"})

	assert.False(t, result.HasError())
	assert.Equal(t, "a6bea192debdbec0d4ab", gists[1].ID)
	assert.Len(t, gists[1].Files, 1)

	file := gists[1].Files["grep_cellar"]
	assert.Equal(t, "grep_cellar", file.FileName)
	assert.Equal(t, "text/plain", file.Type)
	assert.Equal(t, "", file.Language)
	assert.Equal(t, "https://gist.githubusercontent.com/jingweno/a6bea192debdbec0d4ab/raw/80757419d2bd4cfddf7c6be24308eca11b3c330e/grep_cellar", file.RawURL)
	assert.Equal(t, 8107, file.Size)
	assert.Equal(t, false, file.Truncated)

	//Default URL
	stubGet(t, "/gists", "gists", nil)

	gistsDef, resultDef := client.Gists().All(nil, M{})

	assert.False(t, resultDef.HasError())
	assert.Equal(t, "aa5a315d61ae9438b18d", gistsDef[0].ID)
	assert.Len(t, gistsDef, 2)
	assert.Equal(t, gistsDef, gists)

	//Error case
	var invalid = Hyperlink("{")
	gistsErr, resultErr := client.Gists().All(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.Len(t, gistsErr, 0)
}

func TestGistsService_Create(t *testing.T) {
	setup()
	defer tearDown()

	params := Gist{}
	params.Description = "the description for this gist"
	params.Public = true

	var gistFile = GistFile{}
	gistFile.Content = "String file contents"
	params.Files = map[string]*GistFile{"file1.txt": &gistFile}

	mux.HandleFunc("/gists", func(w http.ResponseWriter, r *http.Request) {
		var gistParams Gist
		json.NewDecoder(r.Body).Decode(&gistParams)
		assert.Equal(t, params.Description, gistParams.Description)
		assert.Equal(t, params.Public, gistParams.Public)
		assert.Equal(t, params.Files, gistParams.Files)

		testMethod(t, r, "POST")
		respondWithJSON(w, loadFixture("gist.json"))
	})

	gist, result := client.Gists().Create(&GistsURL, M{}, params)

	assert.False(t, result.HasError())
	assert.Equal(t, "a6bea192debdbec0d4ab", gist.ID)
	assert.Len(t, gist.Files, 1)

	file := gist.Files["grep_cellar"]
	assert.Equal(t, "grep_cellar", file.FileName)
	assert.Equal(t, "text/plain", file.Type)
	assert.Equal(t, "", file.Language)
	assert.Equal(t, "https://gist.githubusercontent.com/jingweno/a6bea192debdbec0d4ab/raw/80757419d2bd4cfddf7c6be24308eca11b3c330e/grep_cellar", file.RawURL)
	assert.Equal(t, 8107, file.Size)
	assert.Equal(t, false, file.Truncated)

	gistNil, resultNil := client.Gists().Create(nil, M{}, params)
	assert.False(t, resultNil.HasError())
	assert.Equal(t, gist, gistNil)

	//Error case
	var invalid = Hyperlink("{")
	gistErr, resultErr := client.Gists().Create(&invalid, M{}, params)
	assert.True(t, resultErr.HasError())
	assert.Nil(t, gistErr)
}

func TestGistsService_Update(t *testing.T) {
	setup()
	defer tearDown()

	params := Gist{}
	params.Description = "the description for this gist"
	params.Public = true

	var gistFile1 = GistFile{}
	gistFile1.Content = "updated file contents"
	var gistFile2 = GistFile{}
	gistFile2.FileName = "new_name.txt"
	gistFile2.Content = "modified contents"
	var gistFile3 = GistFile{}
	gistFile3.Content = "a new file"

	params.Files = map[string]*GistFile{"file1.txt": &gistFile1, "old_name.txt": &gistFile2,
		"new_file.txt": &gistFile3, "delete_this_file.txt": nil}

	mux.HandleFunc("/gists", func(w http.ResponseWriter, r *http.Request) {
		var gistParams Gist
		json.NewDecoder(r.Body).Decode(&gistParams)
		assert.Equal(t, params.Description, gistParams.Description)
		assert.Equal(t, params.Public, gistParams.Public)
		assert.Equal(t, params.Files, gistParams.Files)

		assert.Len(t, params.Files, 4)

		testMethod(t, r, "PATCH")
		respondWithJSON(w, loadFixture("gist.json"))
	})

	gist, result := client.Gists().Update(&GistsURL, M{}, params)

	assert.False(t, result.HasError())
	assert.Equal(t, "a6bea192debdbec0d4ab", gist.ID)
	assert.Len(t, gist.Files, 1)

	file := gist.Files["grep_cellar"]
	assert.Equal(t, "grep_cellar", file.FileName)
	assert.Equal(t, "text/plain", file.Type)
	assert.Equal(t, "", file.Language)
	assert.Equal(t, "https://gist.githubusercontent.com/jingweno/a6bea192debdbec0d4ab/raw/80757419d2bd4cfddf7c6be24308eca11b3c330e/grep_cellar", file.RawURL)
	assert.Equal(t, 8107, file.Size)
	assert.Equal(t, false, file.Truncated)

	gistNil, resultNil := client.Gists().Update(nil, M{}, params)
	assert.False(t, resultNil.HasError())
	assert.Equal(t, gist, gistNil)

	//Error case
	var invalid = Hyperlink("{")
	gistErr, resultErr := client.Gists().Update(&invalid, M{}, params)
	assert.True(t, resultErr.HasError())
	assert.Nil(t, gistErr)
}

func TestGistsService_Commits(t *testing.T) {
	setup()
	defer tearDown()

	//Username URL
	stubGet(t, "/gists/aa5a315d61ae9438b18d/commits", "gist_commits", nil)

	commits, result := client.Gists().Commits(&GistsCommitsURL, M{"gist_id": "aa5a315d61ae9438b18d"})

	assert.False(t, result.HasError())
	assert.Len(t, commits, 1)
	assert.Equal(t, "57a7f021a713b1c5a6a199b54cc514735d2d462f", commits[0].Version)

	assert.Equal(t, commits[0].User.AvatarURL, "https://github.com/images/error/octocat_happy.gif")
	assert.Equal(t, commits[0].ChangeStatus.Additions, 180)

	commitsNil, resultNil := client.Gists().Commits(nil, M{"gist_id": "aa5a315d61ae9438b18d"})
	assert.False(t, resultNil.HasError())
	assert.Equal(t, commitsNil, commits)

	//Error case
	var invalid = Hyperlink("{")
	commitsErr, resultErr := client.Gists().Commits(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.Len(t, commitsErr, 0)
}

func TestGistsService_Star(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/gists/aa5a315d61ae9438b18d/star", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PUT")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	success, result := client.Gists().Star(&GistsStarURL, M{"gist_id": "aa5a315d61ae9438b18d"})
	assert.False(t, result.HasError())
	assert.True(t, success)

	successNil, resultNil := client.Gists().Star(nil, M{"gist_id": "aa5a315d61ae9438b18d"})
	assert.False(t, resultNil.HasError())
	assert.Equal(t, success, successNil)

	//Error case
	var invalid = Hyperlink("{")
	starErr, resultErr := client.Gists().Star(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.False(t, starErr)
}

func TestGistsService_Unstar(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/gists/aa5a315d61ae9438b18d/star", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	success, result := client.Gists().Unstar(&GistsStarURL, M{"gist_id": "aa5a315d61ae9438b18d"})
	assert.False(t, result.HasError())
	assert.True(t, success)

	successNil, resultNil := client.Gists().Unstar(nil, M{"gist_id": "aa5a315d61ae9438b18d"})
	assert.False(t, resultNil.HasError())
	assert.Equal(t, success, successNil)

	//Error case
	var invalid = Hyperlink("{")
	starErr, resultErr := client.Gists().Unstar(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.False(t, starErr)
}

func TestGistsService_CheckStar(t *testing.T) {
	setup()
	defer tearDown()

	// Starred
	mux.HandleFunc("/gists/aa5a315d61ae9438b18d/star", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	success, result := client.Gists().CheckStar(&GistsStarURL, M{"gist_id": "aa5a315d61ae9438b18d"})
	assert.False(t, result.HasError())
	assert.True(t, success)

	// Not starred
	mux.HandleFunc("/gists/a6bea192debdbec0d4ab/star", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 404)
	})

	successNil, resultNil := client.Gists().CheckStar(nil, M{"gist_id": "a6bea192debdbec0d4ab"})
	assert.True(t, resultNil.HasError()) //404 counts as an error...
	assert.False(t, successNil)

	//Error case
	var invalid = Hyperlink("{")
	starErr, resultErr := client.Gists().CheckStar(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.False(t, starErr)
}

func TestGistsService_Fork(t *testing.T) {
	setup()
	defer tearDown()

	stubRequest(t, "POST", "/gists/a6bea192debdbec0d4ab/forks", "gist", nil)

	gist, result := client.Gists().Fork(&GistsForksURL, M{"gist_id": "a6bea192debdbec0d4ab"})

	assert.False(t, result.HasError())
	assert.Equal(t, "a6bea192debdbec0d4ab", gist.ID)
	assert.Len(t, gist.Files, 1)

	file := gist.Files["grep_cellar"]
	assert.Equal(t, "grep_cellar", file.FileName)
	assert.Equal(t, "text/plain", file.Type)
	assert.Equal(t, "", file.Language)
	assert.Equal(t, "https://gist.githubusercontent.com/jingweno/a6bea192debdbec0d4ab/raw/80757419d2bd4cfddf7c6be24308eca11b3c330e/grep_cellar", file.RawURL)
	assert.Equal(t, 8107, file.Size)
	assert.Equal(t, false, file.Truncated)

	gistNil, resultNil := client.Gists().Fork(nil, M{"gist_id": "a6bea192debdbec0d4ab"})
	assert.False(t, resultNil.HasError())
	assert.Equal(t, gist, gistNil)

	//Error case
	var invalid = Hyperlink("{")
	gistErr, resultErr := client.Gists().Fork(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.Nil(t, gistErr)
}

func TestGistsService_ListForks(t *testing.T) {
	setup()
	defer tearDown()

	//Username URL
	stubGet(t, "/gists/dee9c42e4998ce2ea439/forks", "gist_forks", nil)

	forks, result := client.Gists().ListForks(&GistsForksURL, M{"gist_id": "dee9c42e4998ce2ea439"})

	assert.False(t, result.HasError())
	assert.Len(t, forks, 1)
	assert.Equal(t, "dee9c42e4998ce2ea439", forks[0].ID)
	assert.False(t, forks[0].User.SiteAdmin)

	forksNil, resultNil := client.Gists().ListForks(nil, M{"gist_id": "dee9c42e4998ce2ea439"})

	assert.False(t, resultNil.HasError())
	assert.Equal(t, forksNil, forks)

	//Error case
	var invalid = Hyperlink("{")
	forksErr, resultErr := client.Gists().ListForks(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.Len(t, forksErr, 0)
}

func TestGistsService_Delete(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/gists/aa5a315d61ae9438b18d", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	success, result := client.Gists().Delete(&GistsURL, M{"gist_id": "aa5a315d61ae9438b18d"})
	assert.False(t, result.HasError())
	assert.True(t, success)

	successNil, resultNil := client.Gists().Delete(nil, M{"gist_id": "aa5a315d61ae9438b18d"})
	assert.False(t, resultNil.HasError())
	assert.Equal(t, success, successNil)

	//Error case
	var invalid = Hyperlink("{")
	deleteErr, resultErr := client.Gists().Delete(&invalid, M{})
	assert.True(t, resultErr.HasError())
	assert.False(t, deleteErr)
}
