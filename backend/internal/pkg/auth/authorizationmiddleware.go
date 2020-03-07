package auth

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
)

// AuthorizationMiddleware performs authorization check based on the provided
// authorize function.
func AuthorizationMiddleware(authorize func(UserInfo) bool) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			user := req.Context().Value("user")
			sub := user.(*jwt.Token).Claims.(jwt.MapClaims)["sub"]

			fmt.Println(sub)

			if !authorize(UserInfo{}) {
				println("Authorization failed. URL:", req.URL.String())
				// http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
				// return
			}

			next.ServeHTTP(w, req)
		})
	}
}
