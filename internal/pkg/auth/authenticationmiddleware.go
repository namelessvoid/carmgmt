package auth

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type contextKey int

const userKey contextKey = 0

// AuthenticationMiddleware extracts the user authentication from the request.
func AuthenticationMiddleware(correctUsername, correctPassword string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			u := User{isAuthenticated: false}

			username, password, ok := req.BasicAuth()
			if ok {
				isAuthenticated := username == correctUsername && password == correctPassword
				if !isAuthenticated {
					http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
					return
				}
				u.isAuthenticated = isAuthenticated
			}

			ctx := context.WithValue(req.Context(), userKey, u)

			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}

// GetUserFromContext returns the User added to the context
// by the AuthenticationMiddleware
func GetUserFromContext(ctx context.Context) User {
	return ctx.Value(userKey).(User)
}
