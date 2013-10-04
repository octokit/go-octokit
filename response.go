package octokat

type Response struct {
	RawBody []byte
	Error   error
}

func (resp *Response) HasError() bool {
	return resp.Error != nil
}

func (resp *Response) Data(v interface{}) error {
	return jsonUnmarshal(resp.RawBody, v)
}
