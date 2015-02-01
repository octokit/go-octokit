package octokit

import (
	"io"
	"net/url"
)

// Create an UploadsService with the base url.URL
func (c *Client) Uploads(url *url.URL) (uploads *UploadsService) {
	uploads = &UploadsService{client: c, URL: url}
	return
}

type UploadsService struct {
	client *Client
	URL    *url.URL
}

func (u *UploadsService) UploadAsset(asset io.ReadCloser, contentType string, contentLength int64) (result *Result) {
	return u.client.upload(u.URL, asset, contentType, contentLength)
}
