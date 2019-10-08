package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

// CORSMiddleware adds the Access-Control-Allow-Origin header to the response
// At the moment, all origins are allowed, i.e. the header is set to '*'.
func CORSMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, req)
		})
	}
}
