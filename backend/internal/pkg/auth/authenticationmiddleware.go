package auth

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type contextKey int

const userInfoKey contextKey = 0

// AuthenticationMiddleware extracts the user authentication from the request.
func AuthenticationMiddleware(auth Authenticator) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			info := UserInfo{IsAuthenticated: false}
			info.IsAuthenticated = auth.VerifyToken(req)

			ctx := addUserInfoToContext(req.Context(), info)

			next.ServeHTTP(w, req.WithContext(ctx))
		})
	}
}

// GetUserInfoFromContext returns the User added to the context
// by the AuthenticationMiddleware
func GetUserInfoFromContext(ctx context.Context) UserInfo {
	return ctx.Value(userInfoKey).(UserInfo)
}

func addUserInfoToContext(ctx context.Context, user UserInfo) context.Context {
	return context.WithValue(ctx, userInfoKey, user)
}
