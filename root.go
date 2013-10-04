package octokat

import (
	"encoding/json"
	"github.com/octokit/octokat/hyper"
	"regexp"
)

type Root struct {
	links map[string]hyper.Link
}

func (r *Root) Rel(rel string) *hyper.Link {
	if link, ok := r.links[rel]; ok {
		return &link
	}

	return nil
}

func (r *Root) UnmarshalJSON(d []byte) error {
	var out map[string]string

	if err := json.Unmarshal(d, &out); err != nil {
		return err
	}

	r.links = make(map[string]hyper.Link, len(out))

	for rel, link := range out {
		rel = parseRelNameFromURL(rel)
		r.links[rel] = hyper.Link(link)
	}

	return nil
}

func (c *Client) Root(options *Options) (root *Root, err error) {
	err = c.jsonGet("", options, &root)
	return
}

func parseRelNameFromURL(url string) string {
	re := regexp.MustCompile("^(.+)_url")
	if re.MatchString(url) {
		return re.FindStringSubmatch(url)[1]
	}

	return url
}
