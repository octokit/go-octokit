package octokat

type Hyperlink struct {
	client *Client
	Rel    string
	Href   string
}

func (l *Hyperlink) Get(v interface{}, options *Options) (err error) {
	err = l.client.jsonGet(l.Href, options, v)
	return
}
