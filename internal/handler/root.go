package handler

import "net/http"

type RootHandler struct{}

func NewRootHandler() *RootHandler {
	return &RootHandler{}
}

func (h *RootHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	w.Header().Set("Content-Type", "text/plain")

	_, _ = w.Write([]byte("Welcome to EdgeGo"))
}
