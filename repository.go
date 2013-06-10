package octokat

type Repository struct {
	Name     string
	UserName string
}

func (r Repository) String() string {
	return concatPath(r.UserName, r.Name)
}
