package octokit

import (
	"net/url"
)

type GenericService struct {
	client      *Client
	uriTemplate string
}

func mergeMaps(first M, second M) M {
	combinedMap := make(M)
	for k, v := range second {
		combinedMap[k] = v
	}
	for k, v := range first {
		combinedMap[k] = v
	}
	return combinedMap
}

func (g *GenericService) getURL(params M, staticParams M) (*url.URL, error) {
	return Hyperlink(g.uriTemplate).Expand(mergeMaps(params, staticParams))
}
