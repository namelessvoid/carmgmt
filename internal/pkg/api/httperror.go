package api

import (
	"io"
	"net/http"
)

func internalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "[\"error.unknown\"]")
}

func invalidJSONError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "[\"error.invalidJson\"]")
}
