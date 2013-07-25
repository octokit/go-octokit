package octokat

import (
	"fmt"
	"strings"
)

type Params map[string]interface{}

func (p Params) Put(key string, value interface{}) interface{} {
	v, ok := p[key]
	p[key] = value

	if ok {
		return v
	} else {
		return nil
	}
}

func (p Params) Delete(key string) interface{} {
	v, ok := p[key]
	if !ok {
		return nil
	}

	delete(p, key)

	return v
}

func (p Params) Size() int {
	return len(p)
}

func (p Params) Require(fields ...string) error {
	missingFields := []string{}

	for _, field := range fields {
		_, ok := p[field]
		if !ok {
			missingFields = append(missingFields, field)
		}
	}

	if len(missingFields) > 0 {
		return fmt.Errorf("Missing fields: %s", strings.Join(missingFields, ", "))
	}

	return nil
}
