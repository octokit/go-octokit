package octokit

// MardownURL is for rendering an arbitrary markdown document
// MarkdownRawURL is for rendering raw markdown
var (
	MarkdownURL    = Hyperlink("/markdown")
	MarkdownRawURL = Hyperlink("/markdown/raw")
)

// Create a MarkdownService
func (c *Client) Markdown() (m *MarkdownService) {
	m = &MarkdownService{client: c}
	return
}

// A service to return HTML rendered markdown document
type MarkdownService struct {
	client *Client
}

// Renders a markdown document with json input
func (m *MarkdownService) Render(uri *Hyperlink, requestParams interface{}) (renderedHTML string, result *Result) {
	url, err := ExpandWithDefault(uri, &MarkdownURL, nil)
	if err != nil {
		return "", &Result{Err: err}
	}
	result = sendRequest(m.client, url, func(req *Request) (*Response, error) {
		req.setBody(requestParams)
		return req.createResponseRaw(req.Request.Post(), &renderedHTML)
	})
	return
}

// Renders a markdown document with string input
func (m *MarkdownService) RenderRaw(uri *Hyperlink, markdownText *string) (renderedHTML string, result *Result) {
	url, err := ExpandWithDefault(uri, &MarkdownURL, nil)
	if err != nil {
		return "", &Result{Err: err}
	}
	result = sendRequest(m.client, url, func(req *Request) (*Response, error) {
		req.setBodyText(markdownText)
		return req.createResponseRaw(req.Request.Post(), &renderedHTML)
	})
	return
}
