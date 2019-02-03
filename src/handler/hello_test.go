package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHello(t *testing.T) {
	s := httptest.NewServer(http.HandlerFunc(Hello))
	defer s.Close()

	res, err := http.Get(s.URL)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("status code invalid: %d", res.StatusCode)
		return
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Errorf("error: %v", err)
		return
	}

	if string(b) != "Hello, world!\n" {
		t.Errorf("bad response body: %v", string(b))
	}
}
