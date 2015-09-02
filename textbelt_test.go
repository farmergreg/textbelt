package textbelt

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGoodPhoneNumber(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"success\":true}")
	}))
	defer ts.Close()

	//
	c := NewClientFromURL(ts.URL)
	if err := c.Text("123-456-789", "test from go!"); err != nil {
		t.Error("Failed with: " + err.Error())
	}
}

func TestBadPhoneNumber(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "{\"success\":false, \"message\" : \"Oops, that didn't work.\"}")
	}))
	defer ts.Close()

	c := NewClientFromURL(ts.URL)
	if err := c.Text("bad", "no good"); err == nil {
		t.Error("bad is not a valid phone number. This should not have been allowed.")
	}
}
