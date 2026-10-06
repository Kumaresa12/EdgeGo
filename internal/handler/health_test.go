package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	handler := NewHealthHandler()

	request := httptest.NewRequest(
		http.MethodGet,
		"/health",
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

func TestHealthHandlerResponse(t *testing.T) {
	handler := NewHealthHandler()

	request := httptest.NewRequest(
		http.MethodGet,
		"/health",
		nil,
	)

	recorder := httptest.NewRecorder()

	handler.ServeHTTP(
		recorder,
		request,
	)

	expected := `{"status":"UP"}`

	if strings.TrimSpace(recorder.Body.String()) != expected {
		t.Fatalf(
			"expected body %q, got %q",
			expected,
			recorder.Body.String(),
		)
	}

}
