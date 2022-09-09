package doordash

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	c := NewClient("token")
	if c == nil {
		t.Error("Expected client to be not-nil")
	}
}

type TestStruct struct {
	Key string `json:"key"`
}

func TestNewRequest(t *testing.T) {
	c := NewClient("token")

	type foo struct {
		Key string `json:"key"`
	}

	inURL, outURL := "/foo", defaultBaseURL+"foo"
	inBody, outBody := &foo{Key: "Value"}, `{"key":"Value"}`+"\n"
	req, _ := c.NewRequest("GET", inURL, inBody)

	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := io.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%v) Body is %v, want %v", inBody, got, want)
	}
}

func TestDo(t *testing.T) {
	type foo struct {
		A string
	}

	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		reqURL := req.URL.String()
		if reqURL != "/foo" {
			t.Errorf("expected /foo for the test server url, got %s", reqURL)
		}

		rw.Write([]byte(`{"A":"a"}`))
	}))
	defer server.Close()

	url, _ := url.Parse(server.URL + "/")
	c := &Client{url, "token", server.Client()}
	req, err := c.NewRequest("GET", "/foo", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	body := new(foo)
	c.Do(req, body)

	want := &foo{"a"}
	if !reflect.DeepEqual(body, want) {
		t.Errorf("Response body = %v, want %v", body, want)
	}
}
