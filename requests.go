// Go lang http for humans
// simple request powerfull response
package requests

import (
	"net/http"
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

func Post(url string) {

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
