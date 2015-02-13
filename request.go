package r

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
)

// A Request represents an HTTP request received by a server.
type Request struct {
	Url       string
	Method    string
	UserAgent string

	// Body gets string, bytes and io.Reader for other types
	// try to parse as Json.
	Body    interface{}
	Headers map[string]string

	// TODO: QueryString map[string]string

	// Auth get list of strings with the username and password
	// for basic auth header.
	Auth []string
}

// Do sends HTTP request and return a Response type.
func (r *Request) Do() (*Response, error) {
	var req *http.Request
	var err error

	b, e := prepareRequestBody(r.Body)

	if e != nil {
		return nil, e
	}

	req, err = http.NewRequest(r.Method, r.Url, b)

	if err != nil {
		return nil, err
	}

	if r.UserAgent != "" {
		req.Header.Add("User-Agent", r.UserAgent)
	}

	// Convert from map to header type.
	createBasicHeaderType(r.Headers, req)

	// Parse the auth value to basic authatication
	parseAuthValue(r.Auth, req)

	return httpResponseHandler(DefaultClient.Do(req))
}

func (r *Request) Get() (*Response, error) {
	r.Method = "GET"
	return r.Do()
}

func (r *Request) Post() (*Response, error) {
	r.Method = "POST"
	return r.Do()
}

func (r *Request) Put() (*Response, error) {
	r.Method = "PUT"
	return r.Do()
}

func (r *Request) Options() (*Response, error) {
	r.Method = "OPTIONS"
	return r.Do()
}

func (r *Request) Head() (*Response, error) {
	r.Method = "HEAD"
	return r.Do()
}

func (r *Request) Delete() (*Response, error) {
	r.Method = "DELETE"
	return r.Do()
}

func createBasicHeaderType(uh map[string]string, req *http.Request) {

	for key, value := range uh {
		req.Header.Add(key, value)
	}
}

func parseAuthValue(authSlice []string, req *http.Request) {

	if len(authSlice) == 2 {
		req.SetBasicAuth(authSlice[0], authSlice[1])
	}
}

// Get this function from the goReq package
func prepareRequestBody(b interface{}) (io.Reader, error) {
	switch b.(type) {
	case string:
		// treat is as text
		return strings.NewReader(b.(string)), nil
	case io.Reader:
		// treat is as text
		return b.(io.Reader), nil
	case []byte:
		//treat as byte array
		return bytes.NewReader(b.([]byte)), nil
	case nil:
		return nil, nil
	default:
		// try to jsonify it
		j, err := json.Marshal(b)
		if err == nil {
			return bytes.NewReader(j), nil
		}
		return nil, err
	}
}
