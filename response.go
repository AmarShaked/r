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
	Bytes []byte

	// The http.Response type, This type is the original type
	// That return from the server.
	BaseResponse *http.Response

	Close bool

	// The Final URL location.
	Url string

	// The response status as boolean type
	Ok bool
}

// NewResponse get a http.Response and return a Response type.
func NewResponse(resp *http.Response) (*Response, error) {

	r := &Response{
		BaseResponse: resp,
		StatusCode:   resp.StatusCode,
		Status:       resp.Status,
		Proto:        resp.Proto,
		Request:      resp.Request,
		Close:        resp.Close,

		// Return the url as string type.
		Url: resp.Request.URL.String(),

		// Return the status code as boolean value.
		Ok: resp.StatusCode < 400,
	}

	// Read the body and save it as bytes
	err := r.responseBodyToBytes()

	if err != nil {
		return r, err
	}

	return r, nil
}

// Convert the request body(io.Reader) to bytes
// Becouse the io.Reader.Close function
func (r *Response) responseBodyToBytes() error {

	defer r.BaseResponse.Body.Close()
	body, err := ioutil.ReadAll(r.BaseResponse.Body)

	if err != nil {
		return err
	}

	// Save the body for future usage
	r.Bytes = body

	return nil
}

// Get the response body as a string
func (r *Response) Text() string {
	return string(r.Bytes)
}

// Get response header value by name
func (r *Response) Headers(key string) string {

	return r.BaseResponse.Header.Get(key)
}

// Parses the JSON-encoded data and stores the result in the value pointed to by v.
// Used the json.Unmarshal function.
func (r *Response) Json(v interface{}) error {

	err := json.Unmarshal(r.Bytes, v)

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
