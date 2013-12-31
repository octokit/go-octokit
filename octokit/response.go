package octokit

import (
	"github.com/lostisland/go-sawyer"
	"github.com/lostisland/go-sawyer/hypermedia"
	"github.com/lostisland/go-sawyer/mediatype"
	"net/http"
)

type Response struct {
	MediaType *mediatype.MediaType
	Relations hypermedia.Relations
	*http.Response
}

func NewResponse(sawyerResp *sawyer.Response) (resp *Response, err error) {
	if sawyerResp.IsError() {
		err = sawyerResp.ResponseError
		return
	}

	if sawyerResp.IsApiError() {
		err = NewResponseError(sawyerResp)
		return
	}

	relations := hypermedia.HyperHeaderRelations(sawyerResp.Header, nil)
	resp = &Response{Response: sawyerResp.Response, MediaType: sawyerResp.MediaType, Relations: relations}

	return
}
