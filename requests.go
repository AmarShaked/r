// Go lang http for humans
// simple request powerfull response
package requests

import (
	"net/http"
)

func Get(url string) (*Response, error) {

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	r := &Response{
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

	err = bodyToText(r)

	if err != nil {
		return nil, err
	}

	return r, nil
}

func Post(url string) {

}
