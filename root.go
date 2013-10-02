package octokat

import (
	"encoding/json"
	"regexp"
)

type Root struct {
	client *Client
	links  map[string]Hyperlink
}

func (r *Root) Rel(rel string) *Hyperlink {
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

	r.links = make(map[string]Hyperlink, len(out))

	for rel, link := range out {
		rel = parseRelNameFromURL(rel)
		r.links[rel] = Hyperlink{client: r.client, Rel: rel, Href: link}
	}

	return nil
}

func (c *Client) Root(options *Options) (root *Root, err error) {
	root = &Root{}
	root.client = c
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
