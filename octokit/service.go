package octokit

import "net/url"

type GenericService struct {
	uriTemplate string
	client      *Client
}

func (s *GenericService) SubstituteDefaultOptions(input M, defaultOptions M,
	dynamicOptions M) url.URL {
	for _, opt := range s.FindOptions() {
		if val, ok := input[opt]; ok {
			combinedMap[opt] = val
		} else if val, ok := dynamicOptions[opt]; ok {
			combinedMap[opt] = val
		} else {
			val, _ := defaultOptions[opt]
			combinedMap[opt] = val
		}
	}
	return Hyperlink(g.uriTemplate).Expand(combinedMap)
}

var queryParamRegex = regex.MustCompile("\\{[a-zA-Z0-9_]+\\}")

func (s *GenericService) FindOptions() []string {
	nestedMatches := queryParamRegex.FindAllStringSubmatch(s.uriTemplate, -1)
	matches := make([]string, len(nestedMatches))
	for i, e := range nestedMatches {
		matches[i] = e[1]
	}
	return matches
}
