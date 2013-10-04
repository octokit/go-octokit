package hyper

import (
	"encoding/json"
	"fmt"
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

func (r Root) MarshalJSON() ([]byte, error) {
	out := make(map[string]Link)
	for rel, link := range r.links {
		rel = fmt.Sprintf("%s_url", rel)
		out[rel] = link
	}

	return json.Marshal(out)
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

// TODO extract it into hyper.Parser
func parseRelNameFromURL(url string) string {
	re := regexp.MustCompile("^(.+)_url")
	if re.MatchString(url) {
		return re.FindStringSubmatch(url)[1]
	}

	return url
}
