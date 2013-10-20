package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

type ResponseErrorType int

const (
	ErrorClientError             ResponseErrorType = iota // 400-499
	ErrorBadRequest              ResponseErrorType = iota // 400
	ErrorUnauthorized            ResponseErrorType = iota // 401
	ErrorOneTimePasswordRequired ResponseErrorType = iota // 401
	ErrorForbidden               ResponseErrorType = iota // 403
	ErrorTooManyRequests         ResponseErrorType = iota // 403
	ErrorTooManyLoginAttempts    ResponseErrorType = iota // 403
	ErrorNotFound                ResponseErrorType = iota // 404
	ErrorNotAcceptable           ResponseErrorType = iota // 406
	ErrorUnsupportedMediaType    ResponseErrorType = iota // 414
	ErrorUnprocessableEntity     ResponseErrorType = iota // 422
	ErrorServerError             ResponseErrorType = iota // 500-599
	ErrorInternalServerError     ResponseErrorType = iota // 500
	ErrorNotImplemented          ResponseErrorType = iota // 501
	ErrorBadGateway              ResponseErrorType = iota // 502
	ErrorServiceUnavailable      ResponseErrorType = iota // 503
	ErrorMissingContentType      ResponseErrorType = iota
	ErrorUnknownError            ResponseErrorType = iota
)

type ErrorObject struct {
	Resource string `json:"resource,omitempty"`
	Code     string `json:"code,omitempty"`
	Field    string `json:"field,omitempty"`
}

func (e *ErrorObject) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

type ResponseError struct {
	Response         *Response
	Type             ResponseErrorType
	Message          string        `json:"message,omitempty"`
	Err              string        `json:"error,omitempty"`
	Errors           []ErrorObject `json:"errors,omitempty"`
	DocumentationURL string        `json:"documentation_url,omitempty"`
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("%v %v: %d - %s",
		e.Response.Request.Method, e.Response.Request.URL,
		e.Response.StatusCode, e.errorMessage())
}

func (e *ResponseError) errorMessage() string {
	messages := []string{}

	if e.Message != "" {
		messages = append(messages, e.Message)
	}

	if e.Err != "" {
		m := fmt.Sprintf("Error: %s", e.Err)
		messages = append(messages, m)
	}

	if len(e.Errors) > 0 {
		m := []string{}
		m = append(m, "\nError summary:")
		for _, e := range e.Errors {
			m = append(m, fmt.Sprintf("\t%s", e.Error()))
		}
		messages = append(messages, strings.Join(m, "\n"))
	}

	if e.DocumentationURL != "" {
		messages = append(messages, fmt.Sprintf("// See: %s", e.DocumentationURL))
	}

	return strings.Join(messages, "\n")
}

func getResponseErrorType(resp *http.Response) ResponseErrorType {
	code := resp.StatusCode
	switch {
	case code == http.StatusBadRequest:
		return ErrorBadRequest

	case code == http.StatusUnauthorized:
		header := resp.Header.Get("X-GitHub-OTP")
		r := regexp.MustCompile(`(?i)required; (\\w+)`)
		if r.MatchString(header) {
			return ErrorOneTimePasswordRequired
		}

		return ErrorUnauthorized

	case code == http.StatusForbidden:
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return ErrorForbidden
		}
		rr := regexp.MustCompile("(?i)rate limit exceeded")
		if rr.MatchString(string(body)) {
			return ErrorTooManyRequests
		}
		lr := regexp.MustCompile("(?i)login attempts exceeded")
		if lr.MatchString(string(body)) {
			return ErrorTooManyLoginAttempts
		}

		return ErrorForbidden

	case code == http.StatusNotFound:
		return ErrorNotFound

	case code == http.StatusNotAcceptable:
		return ErrorNotAcceptable

	case code == http.StatusUnsupportedMediaType:
		return ErrorUnsupportedMediaType

	case code == 422:
		return ErrorUnprocessableEntity

	case code >= 400 && code <= 499:
		return ErrorClientError

	case code == http.StatusInternalServerError:
		return ErrorInternalServerError

	case code == http.StatusNotImplemented:
		return ErrorNotImplemented

	case code == http.StatusBadGateway:
		return ErrorBadGateway

	case code == http.StatusServiceUnavailable:
		return ErrorServiceUnavailable

	case code >= 500 && code <= 599:
		return ErrorServerError
	}

	return ErrorUnknownError
}
