package octokit

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCommentsService_AllRepoComments(t *testing.T) {
	setup()
	defer tearDown()

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/repos/octokit/go-octokit/comments?page=2"), testURLOf("/repos/octokit/go-octokit/comments?page=3"))
	stubGet(t, "/repos/octokit/go-octokit/comments", "comments", map[string]string{"Link": link})

	comments, result := client.Comments().All(nil, M{"owner": "octokit", "repo": "go-octokit"})
	assert.False(t, result.HasError())
	assert.Len(t, comments, 1)

	comment := comments[0]
	validateComment(t, comment)

	assert.Equal(t, testURLStringOf("/repos/octokit/go-octokit/comments?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/repos/octokit/go-octokit/comments?page=3"), string(*result.LastPage))

	validateNextPage_Comments(t, result)
}

func TestCommentsService_AllCommitComments(t *testing.T) {
	setup()
	defer tearDown()

	link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/repos/octokit/go-octokit/commits/8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09/comments?page=2"), testURLOf("/repos/octokit/go-octokit/commits/8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09/comments?page=3"))
	stubGet(t, "/repos/octokit/go-octokit/commits/8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09/comments", "comments", map[string]string{"Link": link})

	comments, result := client.Comments().All(&CommitCommentsURL, M{"owner": "octokit", "repo": "go-octokit", "sha": "8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09"})
	assert.False(t, result.HasError())
	assert.Len(t, comments, 1)

	comment := comments[0]
	validateComment(t, comment)

	assert.Equal(t, testURLStringOf("/repos/octokit/go-octokit/commits/8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09/comments?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/repos/octokit/go-octokit/commits/8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09/comments?page=3"), string(*result.LastPage))

	validateNextPage_Comments(t, result)
}

func TestCommentsService_OneRepoComment(t *testing.T) {
	setup()
	defer tearDown()

	stubGet(t, "/repos/octokit/go-octokit/comments/4236029", "comment", nil)

	comment, result := client.Comments().One(nil, M{"owner": "octokit", "repo": "go-octokit", "id": 4236029})
	assert.False(t, result.HasError())

	validateComment(t, *comment)
}

func TestCommentsService_CreateCommitComment(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/go-octokit/commits/8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09/comments", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, "{\"body\":\"I am a comment\",\"path\":\"root.go\",\"position\":46}\n")

		respondWithJSON(w, loadFixture("comment.json"))
	})

	input := M{
		"body":     "I am a comment",
		"path":     "root.go",
		"position": 46,
	}

	comment, result := client.Comments().Create(nil, M{"owner": "octokit", "repo": "go-octokit", "sha": "8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09"}, input)
	assert.False(t, result.HasError())

	validateComment(t, *comment)
}

func TestCommentsService_UpdateCommitComment(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/go-octokit/comments/4236029", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "PATCH")
		testBody(t, r, "{\"body\":\"I am a comment\"}\n")

		respondWithJSON(w, loadFixture("comment.json"))
	})

	input := M{"body": "I am a comment"}

	comment, result := client.Comments().Update(nil, M{"owner": "octokit", "repo": "go-octokit", "id": 4236029}, input)
	assert.False(t, result.HasError())

	validateComment(t, *comment)
}

func TestCommentsService_DeleteCommitComment(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/repos/octokit/go-octokit/comments/4236029", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	success, result := client.Comments().Delete(nil, M{"owner": "octokit", "repo": "go-octokit", "id": 4236029})
	assert.False(t, result.HasError())

	assert.True(t, success)
}

func TestPublicKeysService_Failure(t *testing.T) {
	setup()
	defer tearDown()

	url := Hyperlink("}")
	comments, result := client.Comments().All(&url, nil)
	assert.True(t, result.HasError())
	assert.Len(t, comments, 0)

	comment, result := client.Comments().One(&url, nil)
	assert.True(t, result.HasError())
	assert.Nil(t, comment)

	comment, result = client.Comments().Create(&url, nil, nil)
	assert.True(t, result.HasError())
	assert.Nil(t, comment)

	comment, result = client.Comments().Update(&url, nil, nil)
	assert.True(t, result.HasError())
	assert.Nil(t, comment)

	success, result := client.Comments().Delete(&url, nil)
	assert.True(t, result.HasError())
	assert.False(t, success)
}

func validateComment(t *testing.T, comment Comment) {
	testTime, _ := time.Parse("2006-01-02T15:04:05Z", "2013-10-02T19:32:40Z")

	assert.Equal(t, "https://api.github.com/repos/octokit/go-octokit/comments/4236029", comment.URL)
	assert.Equal(t, "https://github.com/octokit/go-octokit/commit/8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09#commitcomment-4236029", comment.HTMLURL)
	assert.Equal(t, 4236029, comment.ID)
	assert.Equal(t, 46, comment.Position)
	assert.Equal(t, 46, comment.Line)
	assert.Equal(t, "root.go", comment.Path)
	assert.Equal(t, "8b8347dc11c81b64fdd9938d34dc4ef6a07dbf09", comment.CommitID)
	assert.Equal(t, &testTime, comment.CreatedAt)
	assert.Equal(t, &testTime, comment.UpdatedAt)
	assert.Equal(t, ":heart:\r\n\r\nAre you handling plain `url`, too? In Octokit.rb, we parse those as a `self` relation.", comment.Body)

	user := comment.User

	assert.Equal(t, "pengwynn", user.Login)
	assert.Equal(t, 865, user.ID)
}

func validateNextPage_Comments(t *testing.T, result *Result) {
	comments, result := client.Comments().All(result.NextPage, nil)
	assert.False(t, result.HasError())
	assert.Len(t, comments, 1)
}
