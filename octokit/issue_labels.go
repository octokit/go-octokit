package octokit

// IssueLabelsURL is a URL template for accessing issue labels
//
// https://developer.github.com/v3/issues/labels/
var IssueLabelsURL = Hyperlink("repos/{owner}/{repo}/issues/{number}/labels")

// IssueLabels creates an IssueLabelsService with a base url
func (c *Client) IssueLabels() (issueLabels *IssueLabelsService) {
	issueLabels = &IssueLabelsService{client: c}
	return
}

// IssueLabelsService is a service providing access to labels for an issue
type IssueLabelsService struct {
	client *Client
}

// Adds labels to an issue
//
// https://developer.github.com/v3/issues/labels/#add-labels-to-an-issue
func (l *IssueLabelsService) Add(uri *Hyperlink, uriParams M, labelsToAdd []string) (labels []Label, result *Result) {
	url, err := ExpandWithDefault(uri, &IssueLabelsURL, uriParams)
	if err != nil {
		return nil, &Result{Err: err}
	}

  result = l.client.post(url, labelsToAdd, &labels)
	return
}
