// Package r is powerful package for quick and simple HTTP requests in Go language.
//
// For a full guide visit https://github.com/AmarShaked/r
// 	package main
//
//	import (
//		"github.com/AmarShaked/r"
//		"fmt"
//	)
//
//	func main() {
//		res, _ := r.Get('https://api.github.com/events')
//		fmt.Println(res.Text())
//	}
package r

import (
	"bytes"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// The DefaultClient is the client that send all the requests
var DefaultClient = &http.Client{}

// Get send a simple GET http request.
func Get(url string) (*Response, error) {

	// Used the default Get function of the DefaultClient
	return httpResponseHandler(DefaultClient.Get(url))
}

// Post send a simple POST http request.
// Post get the map of strings and used it like PostForm.
func Post(url string, data map[string]string) (*Response, error) {

	return httpResponseHandler(DefaultClient.Post(url, "application/x-www-form-urlencoded", strings.NewReader(encode(data))))
}

// Head send a simple HEAD http request.
func Head(url string) (*Response, error) {

	return httpResponseHandler(DefaultClient.Head(url))
}

// Put send a simple PUT http request.
// Put send request with body
func Put(url string, body string) (*Response, error) {

	req, err := http.NewRequest("PUT", url, strings.NewReader(body))

	if err != nil {
		return nil, err
	}

	return httpResponseHandler(DefaultClient.Do(req))
}

// Delete send a quick DELETE http request.
// Delete send a body with the request.
func Delete(url string, body string) (*Response, error) {

	req, err := http.NewRequest("DELETE", url, strings.NewReader(body))

	if err != nil {
		return nil, err
	}

	return httpResponseHandler(DefaultClient.Do(req))
}

// Options send a quick OPTIONS http request.
func Options(url string) (*Response, error) {

	req, err := http.NewRequest("OPTIONS", url, nil)

	if err != nil {
		return nil, err
	}

	return httpResponseHandler(DefaultClient.Do(req))
}

// httpResponseHandler handle all the logic after we get a http.Request.
// Handle all the errors and set the new Response object
func httpResponseHandler(resp *http.Response, err error) (*Response, error) {

	if err != nil {
		return nil, err
	}

	return NewResponse(resp)
}

// This function encode the map to formData or queryString
// {"exs": "shaked", "exs2": "shaked2"} == "exs=shaked&exs2=shaked2"
func encode(v map[string]string) string {

	if v == nil {
		return ""
	}

	var buf bytes.Buffer

	keys := make([]string, 0, len(v))
	for k := range v {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		prefix := url.QueryEscape(k) + "="

		if buf.Len() > 0 {
			buf.WriteByte('&')
		}
		buf.WriteString(prefix)
		buf.WriteString(url.QueryEscape(string(v[k])))

	}
	return buf.String()
}
