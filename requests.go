// Go lang http for humans
// simple request powerfull response
package requests

import (
	"net/http"
)

// A Request represents an HTTP request received by a server.
type Request struct {
	Headers map[string]string
}

func Get(url string) (Response, error) {

	resp, err := http.Get(url)

	if err != nil {
		return Response{}, err
	}

	return Response{
		StatusCode: resp.StatusCode,
		Status:     resp.Status,
		Body:       resp.Body,
		Header:     resp.Header}, nil
}
