package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionHandler(t *testing.T) {
	handler := NewVersionHandler()

	request := httptest.NewRequest(
		http.MethodGet,
		"/version",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(
		recorder,
		request,
	)

	if recorder.Code != http.StatusOK {
		t.Fatalf(
			"expected status %d, got %d",
			http.StatusOK,
			recorder.Code,
		)
	}

}
