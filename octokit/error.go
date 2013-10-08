package octokit

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
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
	Response         *http.Response
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

func checkResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 399 {
		return nil
	}

	responseError := &ResponseError{Response: resp}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if data != nil {
		err = jsonUnmarshal(data, responseError)
	}

	return responseError
}
