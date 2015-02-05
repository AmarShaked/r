package r

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

/* Test helpers */
func expect(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

func refute(t *testing.T, a, b interface{}) {
	if a == b {
		t.Errorf("Did not expect %v (type %v) - Got %v (type %v)", b, reflect.TypeOf(b), a, reflect.TypeOf(a))
	}
}

// Return the url of the test server.
func startTestServer() *httptest.Server {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
	}))

	return ts
}

func Test_Response(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	res, err := Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Ok, true)
	expect(t, res.StatusCode, 200)
	expect(t, res.Url, server.URL)
}

func Test_Headers(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	res, err := Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Headers("content-length"), "3")
}

func Test_Text(t *testing.T) {
	server := startTestServer()
	defer server.Close()

	res, err := Get(server.URL)

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Text(), "GET")
}

func Test_Json(t *testing.T) {

	type TestJson struct {
		Env string `json:"env"`
	}

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		testJson := TestJson{Env: "test"}

		b, err := json.Marshal(testJson)

		if err != nil {
			t.Error(err)
		}

		fmt.Fprint(w, string(b))
	}))

	defer ts.Close()

	res, err := Get(ts.URL)

	if err != nil {
		t.Error(err)
	}

	var json TestJson

	err = res.Json(&json)
	if err != nil {
		t.Error(err)
	}

	expect(t, json.Env, "test")
}

func Test_Cookies(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		c := http.Cookie{Name: "cookieForTest", Value: "test"}

		http.SetCookie(w, &c)
		fmt.Fprint(w, "test")
	}))

	res, err := Get(ts.URL)

	if err != nil {
		t.Error(err)
	}

	expect(t, res.Cookies("cookieForTest"), "test")
}
