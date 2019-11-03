package auth

import (
	"fmt"
	"io"
	"net/http"
)

func LoginHandler(auth *authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token, err := auth.LoginViaSession(r)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		fmt.Println("token from cookie auth", token)
		if token.IsEmpty() {
			token, err = auth.LoginViaCredentials(r)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
			if !token.IsEmpty() {
				auth.CreateSession(w)
			}
		}
		fmt.Println("token from credentials auth", token)
		if token.IsEmpty() {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		io.WriteString(w, "{\"token\":\""+string(token)+"\"}")
	}
}
