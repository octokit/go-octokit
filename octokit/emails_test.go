package octokit

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmailsService_All(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/emails", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")

		header := w.Header()
		link := fmt.Sprintf(`<%s>; rel="next", <%s>; rel="last"`, testURLOf("/user/emails?page=2"), testURLOf("/user/emails?page=3"))
		header.Set("Link", link)

		respondWithJSON(w, loadFixture("emails.json"))
	})

	url, _ := EmailUrl.Expand(nil)
	allEmails, result := client.Emails(url).All()

	assert.False(t, result.HasError())
	assert.Len(t, allEmails, 1)

	email := allEmails[0]
	assert.Equal(t, "rz99@cornell.edu", email.Email)
	assert.Equal(t, true, email.Verified)
	assert.Equal(t, true, email.Primary)

	assert.Equal(t, testURLStringOf("/user/emails?page=2"), string(*result.NextPage))
	assert.Equal(t, testURLStringOf("/user/emails?page=3"), string(*result.LastPage))

	nextPageURL, err := result.NextPage.Expand(nil)
	assert.NoError(t, err)

	allEmails, result = client.Emails(nextPageURL).All()
	assert.False(t, result.HasError())
	assert.Len(t, allEmails, 1)
}

func TestEmailsService_Create(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/emails", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, "[\"test@example.com\",\"otherTest@example.com\"]\n")
		respondWithJSON(w, loadFixture("emails.json"))
	})

	url, _ := EmailUrl.Expand(nil)

	params := []string{"test@example.com", "otherTest@example.com"}
	allEmails, result := client.Emails(url).Create(params)

	assert.False(t, result.HasError())
	assert.Len(t, allEmails, 1)

	email := allEmails[0]
	assert.Equal(t, "rz99@cornell.edu", email.Email)
	assert.Equal(t, true, email.Verified)
	assert.Equal(t, true, email.Primary)
}

func TestEmailsService_Delete(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/user/emails", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		testBody(t, r, "[\"test@example.com\",\"otherTest@example.com\"]\n")

		header := w.Header()
		header.Set("Content-Type", "application/json")

		respondWithStatus(w, 204)
	})

	url, _ := EmailUrl.Expand(nil)

	params := []string{"test@example.com", "otherTest@example.com"}
	result := client.Emails(url).Delete(params)

	assert.False(t, result.HasError())
	assert.Equal(t, 204, result.Response.StatusCode)
}
