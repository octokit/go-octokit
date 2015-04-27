package octokit

import (
	"time"
)

var (
	PublicEventsURL = Hyperlink("/events")
)

func (c *Client) ActivityEvents() *ActivityEventsService {
	return &ActivityEventsService{client: c}
}

// ActivityEventsService is a service providing access to event activity
type ActivityEventsService struct {
	client *Client
}

func (r *ActivityEventsService) One(uri *Hyperlink, params M) (event *Event,
	result *Result) {
	if uri == nil {
		uri = &PublicEventsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}
	result = r.client.get(url, &event)
	return
}

// All gets a list of all events associated with the url of the service
func (r *ActivityEventsService) All(uri *Hyperlink, params M) (
	events []Event, result *Result) {
	if uri == nil {
		uri = &PublicEventsURL
	}
	url, err := uri.Expand(params)
	if err != nil {
		return nil, &Result{Err: err}
	}
	result = r.client.get(url, &events)
	return
}

type Actor struct {
	ID         int    `json:"id,omitempty"`
	Login      string `json:"login,omitempty"`
	GravatarID string `json:"gravatar_id,omitempty"`
	URL        string `json:"url,omitempty"`
	AvatarURL  string `json:"avatar_url,omitempty"`
}

type Payload struct {
	PushID       int      `json:"push_id,omitempty"`
	Size         int      `json:"size,omitempty"`
	DistinctSize int      `json:"distinct_size,omitempty"`
	Ref          string   `json:"ref,omitempty"`
	Head         string   `json:"head,omitempty"`
	Before       string   `json:"before,omitempty"`
	Commits      []Commit `json:"commits,omitempty"`
}

type Event struct {
	ID        string     `json:"id,omitempty"`
	Type      string     `json:"type,omitempty"`
	Actor     Actor      `json:"actor,omitempty"`
	Repo      Repository `json:"repo,omitempty"`
	Payload   Payload    `json:"payload,omitempty"`
	Public    bool       `json:"public,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}
