package middleware_test

import (
	"app/pkg/middleware"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const testRespBody = `{"k":"v"}`

func TestAllResponsesAreJsonEncoded(t *testing.T) {
	r, _ := http.NewRequest("GET", "/healthcheck", nil)
	w := httptest.NewRecorder()

	middleware.ContentTypeJson(http.HandlerFunc(testHandleFunc())).ServeHTTP(w, r)
	res := w.Result()

	if status := res.StatusCode; status != http.StatusOK {
		t.Errorf("Wrong status code: got %v, expected %v", status, http.StatusOK)
	}

	if contentType := res.Header.Get(middleware.HeaderKeyContentType); contentType != middleware.HeaderValueContentType {
		t.Errorf("Wrong content-type: got %v, expected %v", contentType, middleware.HeaderValueContentType)
	}
}

func testHandleFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, testRespBody)
	}
}
