package r

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Get(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	res, err := Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "GET")
}

func Test_Options(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	res, err := Options(server.URL)

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "OPTIONS")
}

func Test_Head(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	res, err := Head(server.URL)

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Request.Method, "HEAD")
}

func Test_Post(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprint(w, r.Form.Get("test"))
	}))

	defer ts.Close()

	res, err := Post(ts.URL, map[string]string{"test": "testValue"})

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Request.Method, "POST")
	expect(t, res.Text(), "testValue")
}

func Test_Put(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprint(w, string(body))
	}))

	defer ts.Close()

	res, err := Put(ts.URL, "test")

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Request.Method, "PUT")
	expect(t, res.Text(), "test")
}

func Test_Delete(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprint(w, string(body))
	}))

	defer ts.Close()

	res, err := Delete(ts.URL, "test")

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Request.Method, "DELETE")
	expect(t, res.Text(), "test")
}
