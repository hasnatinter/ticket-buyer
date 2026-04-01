package middleware_test

import (
	"app/internal/api/health"
	"app/pkg/middleware"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testRespBody = `{"k":"v"}`

func TestAllResponsesAreJsonEncoded(t *testing.T) {
	r, _ := http.NewRequest("GET", "/healthcheck", nil)
	w := httptest.NewRecorder()

	middleware.ContentTypeJson(http.HandlerFunc(health.Read)).ServeHTTP(w, r)
	res := w.Result()

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	if contentType := res.Header.Get(middleware.HeaderKeyContentType); contentType != middleware.HeaderValueContentType {
		t.Errorf("Wrong content-type: got %v, expected %v", contentType, middleware.HeaderValueContentType)
	}
}
