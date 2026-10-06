package handler

import (
	"encoding/json"
	"net/http"
)

const Version = "0.1.0"

type VersionHandler struct{}

func NewVersionHandler() *VersionHandler {
	return &VersionHandler{}
}

func (h *VersionHandler) ServeHTTP(
	w http.ResponseWriter,
	r *http.Request,
) {
	response := map[string]string{
		"version": Version,
		"name":    "EdgeGo",
	}

	w.Header().Set("content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		return
	}
}
