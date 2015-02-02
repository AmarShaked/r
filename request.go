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
	Url         string
	Method      string
	Body        interface{}
	Headers     map[string]string
	QueryString map[string]string
	Auth        [2]string
}

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

	req.Header = createHeaderObject(r.Headers)

	if len(r.Auth) == 2 {
		req.SetBasicAuth(r.Auth[0], r.Auth[1])
	}

	return responseHandler(DefaultClient.Do(req))
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

func createHeaderObject(uh map[string]string) http.Header {
	var h http.Header

	for key, value := range uh {
		h.Add(key, value)
	}

	return h
}

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
