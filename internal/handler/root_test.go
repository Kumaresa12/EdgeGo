package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRootHandler(t *testing.T) {
	handler := NewRootHandler()

	request := httptest.NewRequest(
		http.MethodGet,
		"/",
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

	expected := "Welcome to EdgeGo"

	if recorder.Body.String() != expected {
		t.Fatalf(
			"expected body %q, got %q",
			expected,
			recorder.Body.String(),
		)
	}
}
