package goget

import (
	"fmt"
	"github.com/mholt/caddy/caddyhttp/httpserver"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGoget_ServeHTTP(t *testing.T) {
	rw := Goget{
		Next: httpserver.HandlerFunc(contentHandler),
		Rule: "github.com/$1/$2",
	}
	req, err := http.NewRequest("GET", "https://tw.p/a?go-get=1", nil)
	if err != nil {
		t.Fatalf("Test : Could not create HTTP request %v", err)
	}
	rec := httptest.NewRecorder()
	result, err := rw.ServeHTTP(rec, req)

	fmt.Println(result)
	fmt.Println(rec.Body)
}

func contentHandler(w http.ResponseWriter, r *http.Request) (int, error) {
	fmt.Fprintf(w, r.URL.String())
	return http.StatusOK, nil
}