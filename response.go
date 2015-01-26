package requests

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode       int
	Status           string
	Body             io.ReadCloser
	Header           http.Header
	Proto            string
	ProtoMajor       int
	ProtoMinor       int
	ContentLength    int64
	TransferEncoding []string
	Request          *http.Request
	Close            bool
	Trailer          http.Header
	TLS              *tls.ConnectionState
	Bytes            []byte
	BaseResponse     *http.Response
}

func bodyToBytes(r *Response) error {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	r.Bytes = body

	return nil
}

func (r *Response) Text() string {
	return string(r.Bytes)
}

func (r *Response) Headers(key string) string {
	return r.Header.Get(key)
}

// Parses the JSON-encoded data and stores the result in the value pointed to by v.
func (r *Response) Json(v interface{}) error {

	err := json.Unmarshal(r.Bytes, v)

	if err != nil {
		return err
	}

	return nil
}

// Get the cookie value by cookie name.
func (r *Response) Cookies(name string) string {
	cookies := r.BaseResponse.Cookies()

	for _, c := range cookies {
		if c.Name == name {
			return c.Value
		}
	}

	return ""
}
