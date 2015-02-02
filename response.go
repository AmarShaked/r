// Copyright 2015 Shaked Amar.
package r

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Response struct {
	StatusCode int    // e.g. 200
	Status     string // e.g. "200 OK"
	Proto      string // e.g. "HTTP/1.0"

	// The Request that was sent to obtain this Response.
	// Request's Body is nil (having already been consumed).
	// This is only populated for Client requests.
	Request *http.Request

	// The value is the BaseResponse Body as bytes
	// Used by the Text function and the Json function.
	bytes []byte

	// The http.Response type, This type is the original type
	// That return from the server.
	BaseResponse *http.Response
}

// Convert the request body(io.Reader) to bytes
// Becouse the io.Reader.Close function
func responseBodyToBytes(r *Response) error {

	defer r.BaseResponse.Body.Close()
	body, err := ioutil.ReadAll(r.BaseResponse.Body)

	if err != nil {
		return err
	}

	// Save the body for future usage
	r.bytes = body

	return nil
}

// Get the response body as a string
func (r *Response) Text() string {
	return string(r.bytes)
}

// Get response header value by name
func (r *Response) Headers(key string) string {
	return r.BaseResponse.Header.Get(key)
}

// Parses the JSON-encoded data and stores the result in the value pointed to by v.
// Used the json.Unmarshal function.
func (r *Response) Json(v interface{}) error {

	err := json.Unmarshal(r.bytes, v)

	if err != nil {
		return err
	}

	return nil
}

// Get the cookie value by cookie name.
// If the cookie doesnt exists return a empty string.
func (r *Response) Cookies(name string) string {
	cookies := r.BaseResponse.Cookies()

	for _, c := range cookies {
		if c.Name == name {
			return c.Value
		}
	}

	return ""
}

// TODO add more features to the response
