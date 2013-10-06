package octokat

type Requester interface {
	Request(v interface{}) (*Response, error)
}
