package octokit

import (
	"net"
	"strconv"

	"github.com/jingweno/go-sawyer/hypermedia"
)

// https://developer.github.com/v3/meta/
var (
	MetaURL = Hyperlink("/meta")
)

// Meta return an APIInfo with the current API meta information
func (c *Client) Meta(uri *Hyperlink) (info APIInfo, result *Result) {
	url, _ := uri.Expand(nil)
	var meta meta
	result = c.get(url, &meta)
	if !result.HasError() {
		info = meta.transform()
	}
	return
}

type ipNet struct {
	*net.IPNet
}

func (i *ipNet) UnmarshalJSON(s []byte) error {
	str, err := strconv.Unquote(string(s))
	if err != nil {
		return err
	}
	_, ipNet, err := net.ParseCIDR(str)
	i.IPNet = ipNet
	return err
}

type ip struct {
	net.IP
}

func (i *ip) UnmarshalJSON(s []byte) error {
	str, err := strconv.Unquote(string(s))
	if err != nil {
		return err
	}
	i.IP = net.ParseIP(str)
	return nil
}

type meta struct {
	*hypermedia.HALResource

	VerifiablePasswordAuthentication bool     `json:"verifiable_password_authentication,omitempty"`
	GithubServicesSha                string   `json:"github_services_sha,omitempty"`
	Hooks                            []*ipNet `json:"hooks,omitempty"`
	Git                              []*ipNet `json:"git,omitempty"`
	Pages                            []*ipNet `json:"pages,omitempty"`
	Importer                         []ip     `json:"importer,omitempty"`
}

func (m meta) transform() (info APIInfo) {
	info.VerifiablePasswordAuthentication = m.VerifiablePasswordAuthentication
	info.GithubServicesSha = m.GithubServicesSha

	hooks := make([]*net.IPNet, len(m.Hooks))
	for i, addr := range m.Hooks {
		hooks[i] = addr.IPNet
	}
	info.Hooks = hooks

	git := make([]*net.IPNet, len(m.Git))
	for i, addr := range m.Git {
		git[i] = addr.IPNet
	}
	info.Git = git

	pages := make([]*net.IPNet, len(m.Pages))
	for i, addr := range m.Pages {
		pages[i] = addr.IPNet
	}
	info.Pages = pages

	importer := make([]net.IP, len(m.Importer))
	for i, addr := range m.Importer {
		importer[i] = addr.IP
	}
	info.Importer = importer
	return
}

// APIInfo contains the information described in https://developer.github.com/v3/meta/#body
type APIInfo struct {
	*hypermedia.HALResource

	VerifiablePasswordAuthentication bool         `json:"verifiable_password_authentication,omitempty"`
	GithubServicesSha                string       `json:"github_services_sha,omitempty"`
	Hooks                            []*net.IPNet `json:"hooks,omitempty"`
	Git                              []*net.IPNet `json:"git,omitempty"`
	Pages                            []*net.IPNet `json:"pages,omitempty"`
	Importer                         []net.IP     `json:"importer,omitempty"`
}
