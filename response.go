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
	Text             string
}

func bodyToText(r *Response) error {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		return err
	}

	r.Bytes = body
	r.Text = string(body)

	return nil
}

func (r *Response) Headers(key string) string {
	return r.Header.Get(key)
}

func (r *Response) Json(v interface{}) error {

	err := json.Unmarshal(r.Bytes, v)

	if err != nil {
		return err
	}

	return nil
}
