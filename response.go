package requests

import (
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int
	Status     string
	Body       io.ReadCloser
	Header	   http.Header 
}

func (r *Response) Text() string {

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		//	return err.Error()
	}

	return string(body)
}

func (r *Response) Headers(header string) {
	
}
