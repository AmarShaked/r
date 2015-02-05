package r

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func Test_Method_Get(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	r := &Request{Method: "GET", Url: server.URL}

	res, err := r.Do()

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "GET")
}

func Test_Method_Get_Direct(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	r := &Request{Url: server.URL}

	res, err := r.Get()

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "GET")
}

func Test_Method_Post(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	r := &Request{Method: "POST", Url: server.URL}

	res, err := r.Do()

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "POST")
}

func Test_Body_String(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprint(w, string(body))
	}))

	defer ts.Close()

	r := &Request{Method: "POST", Url: ts.URL, Body: "Test"}

	res, err := r.Do()

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "Test")
}

func Test_Body_IOReader(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprint(w, string(body))
	}))

	defer ts.Close()

	r := &Request{Method: "POST", Url: ts.URL, Body: strings.NewReader("Test")}

	res, err := r.Do()

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "Test")
}

func Test_Body_Bytes(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprint(w, string(body))
	}))

	defer ts.Close()

	r := &Request{Method: "POST", Url: ts.URL, Body: []byte("Test")}

	res, err := r.Do()

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "Test")
}

func Test_Body_Json(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()
		body, _ := ioutil.ReadAll(r.Body)
		fmt.Fprint(w, string(body))
	}))

	defer ts.Close()

	type Test struct {
		Name string `json: "name"`
	}

	j := &Test{Name: "Test"}
	r := &Request{Method: "POST", Url: ts.URL, Body: j}

	res, err := r.Do()

	if err != nil {
		t.Error(err)
	}

	b, _ := json.Marshal(j)

	expect(t, res.Text(), string(b))
}

func Test_Request_Headers(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	r := &Request{Method: "POST", Url: server.URL, Headers: map[string]string{"test": "Test"}}

	res, err := r.Do()

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Request.Header.Get("test"), "Test")
}

func Test_Request_Auth(t *testing.T) {

	server := startTestServer()
	defer server.Close()

	r := &Request{Method: "POST", Url: server.URL, Auth: []string{"test", "test"}}

	res, err := r.Do()

	if err != nil {
		t.Error(err)
	}

	user, pass, ok := res.Request.BasicAuth()

	if !ok {
		t.Error("Error in the basic auth..")
	}

	expect(t, user, "test")
	expect(t, pass, "test")
}
