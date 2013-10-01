package octokat

import (
	"time"
)

type Asset struct {
	ID            int       `json:"id,omitempty"`
	Name          string    `json:"name,omitempty"`
	Label         string    `json:"label,omitempty"`
	ContentType   string    `json:"content_type,omitempty"`
	State         string    `json:"state,omitempty"`
	Size          int       `json:"size,omitempty"`
	DownloadCount int       `json:"download_count,omitempty"`
	URL           string    `json:"url,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}
