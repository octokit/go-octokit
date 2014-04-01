package octokit

type Middleware interface {
	PrepareRequest(req *Request) error
	PrepareResponse(resp *Response) error
}
