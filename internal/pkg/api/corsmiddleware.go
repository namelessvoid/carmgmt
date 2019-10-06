package api

import (
	"net/http"

	"github.com/gorilla/mux"
)

func CORSMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, req)
		})
	}
}
