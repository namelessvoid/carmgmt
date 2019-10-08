package api

import (
	"io"
	"net/http"
)

func httpError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, "[\"error.unknown\"]")
}
