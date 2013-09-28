package octokat

import (
	"fmt"
)

type Repo struct {
	Name     string
	UserName string
}

func (r Repo) String() string {
	return fmt.Sprintf("%s/%s", r.UserName, r.Name)
}
