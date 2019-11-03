package auth_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/namelessvoid/carmgmt/internal/pkg/auth"
)

type stubHandler struct {
}

func Test_AuthenticationMiddleware(t *testing.T) {
	var actualContext context.Context
	next := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		actualContext = req.Context()
	})

	const (
		correctUsername = "correctUsername"
		correctPassword = "correctPassword"
	)

	middleware := auth.AuthenticationMiddleware(correctUsername, correctPassword)(next)

	t.Run("adds unauthenticated user when no authentication is provided", func(t *testing.T) {
		actualContext = nil
		res := httptest.NewRecorder()
		req, err := http.NewRequest("", "", nil)
		if err != nil {
			t.Fatal(err)
		}

		middleware.ServeHTTP(res, req)

		actualUser := auth.GetUserFromContext(actualContext)
		if actualUser.IsAuthenticated() != false {
			t.Error("actualUser.IsAuthenticated() has unexpected value: got true want false")
		}
	})

	t.Run("adds authenticated user when valid authentication is provided", func(t *testing.T) {
		actualContext = nil
		res := httptest.NewRecorder()
		req, err := http.NewRequest("", "", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.SetBasicAuth(correctUsername, correctPassword)

		middleware.ServeHTTP(res, req)
		actualUser := auth.GetUserFromContext(actualContext)

		if actualUser.IsAuthenticated() != true {
			t.Error("actualUser.IsAuthenticated() has unexpected value: got false want true")
		}
	})

	invalidAuthenticationTests := []struct {
		reason   string
		username string
		password string
	}{{
		reason:   "username is invalid",
		username: "invalidUsername",
		password: correctPassword,
	}, {
		reason:   "password is invalid",
		username: correctUsername,
		password: "invalidPassword",
	}, {
		reason:   "username and password are invalid",
		username: "invalidUsername",
		password: "invalidPassword",
	}}
	for _, run := range invalidAuthenticationTests {
		t.Run("returns with 401 if invalid"+run.reason, func(t *testing.T) {
			actualContext = nil
			res := httptest.NewRecorder()
			req, err := http.NewRequest("", "", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetBasicAuth(run.username, run.password)

			middleware.ServeHTTP(res, req)

			if actualContext != nil {
				t.Error("expected request to be aborted but AuthenticationMiddleware forwarded request to next handler")
			}

			if res.Code != http.StatusUnauthorized {
				t.Errorf("AuthenticationMiddleware returned unexpected status code: got %v want %v", res.Code, http.StatusUnauthorized)
			}
		})
	}
}
