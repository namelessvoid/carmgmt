package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AuthorizationMiddleware performs authorization check based on the provided
// authorize function.
func AuthorizationMiddleware(authorize func(UserInfo) bool) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			info := GetUserInfoFromContext(req.Context())

			if !authorize(info) {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
