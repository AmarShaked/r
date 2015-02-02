// Go lang http for humans
// simple request powerfull response
package requests

import (
	"bytes"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

var DefaultClient = &http.Client{}

// Get send a quick GET http request.
func Get(url string) (*Response, error) {

	return responseHandler(DefaultClient.Get(url))
}

// Post send a quick POST http request.
func Post(url string, data map[string]string) (*Response, error) {

	return responseHandler(DefaultClient.Post(url, "application/x-www-form-urlencoded", strings.NewReader(encode(data))))
}

// Head send a quick HEAD http request.
func Head(url string) (*Response, error) {

	return responseHandler(DefaultClient.Head(url))
}

// Put send a quick PUT http request.
func Put(url string, data map[string]string) (*Response, error) {

	req, err := http.NewRequest("PUT", url, strings.NewReader(encode(data)))

	if err != nil {
		return nil, err
	}

	return responseHandler(DefaultClient.Do(req))

}

// Delete send a quick DELETE http request.
func Delete(url string, data map[string]string) (*Response, error) {

	req, err := http.NewRequest("DELETE", url, strings.NewReader(encode(data)))

	if err != nil {
		return nil, err
	}

	return responseHandler(DefaultClient.Do(req))
}

// Options send a quick OPTIONS http request.
func Options(url string) (*Response, error) {

	req, err := http.NewRequest("OPTIONS", url, nil)

	if err != nil {
		return nil, err
	}

	return responseHandler(DefaultClient.Do(req))
}

// This function handle all the logic after we get a http.Request.
// Handle all the errors and set the new Response object
func responseHandler(resp *http.Response, err error) (*Response, error) {

	if err != nil {
		return nil, err
	}

	r := setNewResponse(resp)

	// Read the body and saveit as bytes
	err = responseBodyToBytes(r)

	if err != nil {
		return nil, err
	}

	return r, nil
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

func setNewResponse(resp *http.Response) *Response {
	return &Response{
		BaseResponse: resp,
		StatusCode:   resp.StatusCode,
		Status:       resp.Status,
		Proto:        resp.Proto,
		Request:      resp.Request,
	}
}
