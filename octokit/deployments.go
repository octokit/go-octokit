package octokit

import (
	"time"
	"github.com/jingweno/go-sawyer/hypermedia"
)

// https://developer.github.com/v3/repos/
var DeploymentsURL = Hyperlink("/repos/{owner}/{repo}/deployments")

// Deployments creates a DeploymentsService with a base url
func (c *Client) Deployments() (deployments *DeploymentsService) {
	deployments = &DeploymentsService{client: c}
	return
}

// DeploymentsService is a service providing access to deployments from a particular url
type DeploymentsService struct {
	client *Client
}

// All gets a list of all deployments associated with the URL of the service
func (c *DeploymentsService) All(uri *Hyperlink, params M) (deployments []Deployment, result *Result) {
	if uri == nil {
		uri = &DeploymentsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return make([]Deployment, 0), &Result{Err: err}
	}
	result = c.client.get(url, &deployments)

	return
}

// Create posts a new deployment based on parameters in a Deployment struct to
// the deployment service url
//
func (r *DeploymentsService) Create(uri *Hyperlink, uriParams M,
	params interface{}) (repo *Deployment, result *Result) {
	if uri == nil {
		uri = &DeploymentsURL
	}

	url, err := uri.Expand(uriParams)
	if err != nil {
		return nil, &Result{Err: err}
	}
	result = r.client.post(url, params, &repo)
	return
}

// Payload represents the payload on the changes made in a deployment
type DeploymentPayload struct {
	Task string `json:"task,omitempty"`
}

// Deployment is a representation of a full deployment in git
type Deployment struct {
	*hypermedia.HALResource

	URL						string							`json:"url,omitempty"`
	ID						int									`json:"id,omitempty"`
	NodeID				string							`json:"node_id,omitempty"`
	Sha						string							`json:"sha,omitempty"`
	Ref						string							`json:"ref,omitempty"`
	Task					string							`json:"task,omitempty"`
	Payload       DeploymentPayload		`json:"payload,omitempty"`
	Environment		string							`json:"environment,omitempty"`
	Description		string							`json:"description,omitempty"`
	Creator       *User								`json:"creator,omitempty"`
	CreatedAt     *time.Time					`json:"created_at,omitempty"`
	UpdatedAt     *time.Time					`json:"updated_at,omitempty"`
	StatusesURL   string							`json:"statuses_url,omitempty"`
	RepositoryURL string							`json:"repository_url,omitempty"`
}

// DeploymentParams represents the struture used to create a Deployment
type DeploymentParams struct {
	Ref							 string   `json:"ref,omitempty"`
	Task						 string   `json:"body,omitempty"`
	AutoMerge				 bool     `json:"auto_merge,omitempty"`
	RequiredContexts []string `json:"required_contexts,omitempty"`
	Payload					 string   `json:"payload,omitempty"`
	Environment      string		`json:"environment,omitempty"`
	Description      string   `json:"description,omitempty"`
}
