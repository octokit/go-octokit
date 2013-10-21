package http

import (
	"github.com/lostisland/go-sawyer"
	"github.com/lostisland/go-sawyer/mediatype"
)

type Headers map[string]string

type Request struct {
	Headers   Headers
	sawyerReq *sawyer.Request
}

func (r *Request) Head(output interface{}) (resp *Response, err error) {
	resp, err = r.do("HEAD", nil, output)
	return
}

func (r *Request) Get(output interface{}) (resp *Response, err error) {
	resp, err = r.do("GET", nil, output)
	return
}

func (r *Request) Post(input interface{}, output interface{}) (resp *Response, err error) {
	resp, err = r.do("POST", input, output)
	return
}

func (r *Request) Put(input interface{}, output interface{}) (resp *Response, err error) {
	resp, err = r.do("PUT", input, output)
	return
}

func (r *Request) Delete(output interface{}) (resp *Response, err error) {
	resp, err = r.do("DELETE", nil, output)
	return
}

func (r *Request) do(method string, input interface{}, output interface{}) (resp *Response, err error) {
	var sawyerResp *sawyer.Response
	switch method {
	case "HEAD":
		sawyerResp = r.sawyerReq.Head(output)
	case "GET":
		sawyerResp = r.sawyerReq.Get(output)
	case "POST":
		mtype, _ := mediatype.Parse("application/json")
		r.sawyerReq.SetBody(mtype, input)
		sawyerResp = r.sawyerReq.Post(output)
	case "PUT":
		mtype, _ := mediatype.Parse("application/json")
		r.sawyerReq.SetBody(mtype, input)
		sawyerResp = r.sawyerReq.Put(output)
	case "PATCH":
		sawyerResp = r.sawyerReq.Patch(output)
	case "DELETE":
		sawyerResp = r.sawyerReq.Delete(output)
	case "OPTIONS":
		sawyerResp = r.sawyerReq.Options(output)
	}

	if sawyerResp.IsError() {
		err = sawyerResp.ResponseError
		return
	}

	if sawyerResp.IsApiError() {
		err = NewResponseError(sawyerResp)
		return
	}

	resp = &Response{Response: sawyerResp.Response}
	return
}
