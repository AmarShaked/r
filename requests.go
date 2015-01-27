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

func Get(url string) (*Response, error) {

	resp, err := DefaultClient.Get(url)

	if err != nil {
		return nil, err
	}

	r := setNewResponse(resp)

	err = bodyToBytes(r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func Post(url string, data map[string]string) (*Response, error) {

	resp, err := DefaultClient.Post(url, "application/x-www-form-urlencoded", strings.NewReader(encode(data)))

	if err != nil {
		return nil, err
	}

	r := setNewResponse(resp)

	err = bodyToBytes(r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

// This function encode the map to formData
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
		BaseResponse:     resp,
		StatusCode:       resp.StatusCode,
		Status:           resp.Status,
		Body:             resp.Body,
		Header:           resp.Header,
		Proto:            resp.Proto,
		ProtoMajor:       resp.ProtoMajor,
		ProtoMinor:       resp.ProtoMinor,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		Request:          resp.Request,
		Close:            resp.Close,
		Trailer:          resp.Trailer,
		TLS:              resp.TLS}
}
