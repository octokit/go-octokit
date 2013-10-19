package http

type Headers map[string]string

type Request struct {
	client  *Client
	URL     string
	Headers Headers
}

func (r *Request) Head(output interface{}) (resp *Response, err error) {
	var respErr *ResponseError
	sawyerReq, err := r.client.sawyerClient.NewRequest(r.URL, respErr)
	if err != nil {
		return
	}

	sawyerResp := sawyerReq.Head(output)
	resp = &Response{Response: sawyerResp.Response, Error: respErr}

	return
}
