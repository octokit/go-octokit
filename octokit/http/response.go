package http

import (
	"net/http"
)

type Response struct {
	Error *ResponseError
	*http.Response
}

func (r *Response) IsError() bool {
	return r.Error != nil
}
