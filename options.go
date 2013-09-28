package octokat

type Options struct {
	Headers Headers
	Params  Params
}

type Headers map[string]string
