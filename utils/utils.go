package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/xid"
)

func ExecuteRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	http.DefaultServeMux.ServeHTTP(rr, req)
	return rr
}

func CheckResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Status code incorrect. Expected: %d, Obtained: %d", expected, actual)
	}
}

func GenerateUID() string {
	id := xid.New()
	return id.String()
}
