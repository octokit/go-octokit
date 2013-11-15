package octokit

import (
	"github.com/lostisland/go-sawyer/hypermedia"
)

// Resource
type Resource struct {
	*hypermedia.HALResource
	rels hypermedia.Relations `json:"-"`
}

func (r *Resource) RelsOf(resource interface{}) hypermedia.Relations {
	if r.rels == nil || len(r.rels) == 0 {
		r.rels = hypermedia.HyperFieldDecoder(resource)
		for key, hyperlink := range r.HALResource.Rels() {
			r.rels[key] = hyperlink
		}
	}
	return r.rels
}

func NewResource() *Resource {
	return &Resource{&hypermedia.HALResource{}, nil}
}
