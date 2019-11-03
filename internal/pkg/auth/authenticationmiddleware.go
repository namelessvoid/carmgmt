package auth

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type contextKey int

const userKey contextKey = 0

// AuthenticationMiddleware extracts the user authentication from the request.
func AuthenticationMiddleware(auth Authenticator) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			u := User{isAuthenticated: false}
			u.isAuthenticated = auth.VerifyToken(req)

			ctx := addUserToContext(req.Context(), u)

			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}

// GetUserFromContext returns the User added to the context
// by the AuthenticationMiddleware
func GetUserFromContext(ctx context.Context) User {
	return ctx.Value(userKey).(User)
}

func addUserToContext(ctx context.Context, user User) context.Context {
	return context.WithValue(ctx, userKey, user)
}
