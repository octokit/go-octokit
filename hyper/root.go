package hyper

import (
	"encoding/json"
	"regexp"
)

type Root struct {
	links map[string]Link
}

func (r *Root) Rel(rel string) *Link {
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

	r.links = make(map[string]Link, len(out))

	for rel, link := range out {
		rel = parseRelNameFromURL(rel)
		r.links[rel] = Link(link)
	}

	return nil
}

func parseRelNameFromURL(url string) string {
	re := regexp.MustCompile("^(.+)_url")
	if re.MatchString(url) {
		return re.FindStringSubmatch(url)[1]
	}

	return url
}
