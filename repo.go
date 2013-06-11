package octokat

type Repo struct {
	Name     string
	UserName string
}

func (r Repo) String() string {
	return concatPath(r.UserName, r.Name)
}
