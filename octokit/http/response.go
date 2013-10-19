package http

import (
	"net/http"
)

type Response struct {
	Error *ResponseError
	*http.Response
}
