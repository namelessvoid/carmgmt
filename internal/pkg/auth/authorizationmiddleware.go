package auth

import (
	"net/http"

	"github.com/gorilla/mux"
)

// AuthorizationMiddleware performs authorization check based on the provided
// authorize function.
func AuthorizationMiddleware(authorize func(User) (bool, error)) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			u := GetUserFromContext(req.Context())

			isAuthorized, err := authorize(u)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}

			if !isAuthorized {
				http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				return
			}

			next.ServeHTTP(w, req)
		})
	}
}
