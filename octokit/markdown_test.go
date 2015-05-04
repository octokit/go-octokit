package octokit

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarkdownService_JSON(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/markdown", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, "{\"context\":\"github/gollum\",\"mode\":\"gfm\",\"text\":\"Hello world github/linguist#1 **cool**, and #1!\"}\n")

		respondWith(w, loadFixture("markdown.html"))
	})

	input := M{
		"text":    "Hello world github/linguist#1 **cool**, and #1!",
		"mode":    "gfm",
		"context": "github/gollum",
	}

	markdown, result := client.Markdown().Render(nil, input)
	assert.False(t, result.HasError())

	expected := "<p>Hello world \n" +
		"    <a href=\"https://github.com/github/linguist/issues/1\" class=\"issue-link\" title=\"Binary detection issues on extensionless files\">github/linguist#1</a>\n" +
		"    <strong>cool</strong>, and \n" +
		"    <a href=\"https://github.com/gollum/gollum/issues/1\" class=\"issue-link\" title=\"no method to write a file?\">#1</a>!\n" +
		"</p>"
	assert.Equal(t, expected, markdown)
}

func TestMarkdownService_RAW(t *testing.T) {
	setup()
	defer tearDown()

	mux.HandleFunc("/markdown/raw", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		testBody(t, r, "Hello world github/linguist#1 **cool**, and #1!")

		respondWith(w, loadFixture("markdown_raw.html"))
	})

	input := "Hello world github/linguist#1 **cool**, and #1!"

	markdown, result := client.Markdown().RenderRaw(&MarkdownRawURL, &input)
	assert.False(t, result.HasError())

	expected := "<p>Hello world github/linguist#1 \n" +
		"    <strong>cool</strong>, and #1!\n" +
		"</p>"
	assert.Equal(t, expected, markdown)
}

func TestMarkdownService_Failure(t *testing.T) {
	setup()
	defer tearDown()

	url := Hyperlink("}")
	markdown, result := client.Markdown().Render(&url, nil)
	assert.True(t, result.HasError())
	assert.Empty(t, markdown)
}
