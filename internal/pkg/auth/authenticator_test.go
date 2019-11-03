package auth_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/namelessvoid/carmgmt/internal/pkg/auth"
)

func Test_Authenticator_LoginViaCredentials(t *testing.T) {
	const (
		correctUsername = "correctUsername"
		correctPassword = "correctPassword"
	)

	t.Run("returns valid token if credentials are correct", func(t *testing.T) {
		req, err := http.NewRequest("", "", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.SetBasicAuth(correctUsername, correctPassword)

		authenticator := auth.NewAuthenticator(correctUsername, correctPassword)

		token, err := authenticator.LoginViaCredentials(req)

		if err != nil {
			t.Errorf("LoginViaCredentials() returned unexpected error: got %v want nil", err)
		}

		if token.IsEmpty() {
			t.Errorf("LoginViaCredentials() returned empty token but non-empty token was expected")
		}

		// Verify that returned token is valid
		req, err = http.NewRequest("", "", nil)
		if err != nil {
			t.Fatal(err)
		}

		req.Header.Set("Authorization", "Bearer "+string(token))
		if !authenticator.VerifyToken(req) {
			t.Errorf("LoginViaCredentials() returned token token which cannot be verified by VerifyToken(): %v", token)
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
		t.Run("returns empty token if "+run.reason, func(t *testing.T) {
			req, err := http.NewRequest("", "", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetBasicAuth(run.username, run.password)

			authenticator := auth.NewAuthenticator(correctUsername, correctPassword)

			token, err := authenticator.LoginViaCredentials(req)

			if err != nil {
				t.Errorf("LoginViaCredentials() returned unexpected error: got %v want nil", err)
			}

			if token.IsEmpty() != true {
				t.Errorf("LoginViaCredentials() returned unexpected token: got %v want \"\"", token)
			}
		})
	}
}

func Test_Authenticator_CreateSession(t *testing.T) {
	t.Run("adds token to the response", func(t *testing.T) {
		expectedCookie := "FLEETMGMT_SESSION=somesession; Path=/; Domain=localhost; HttpOnly"

		res := httptest.NewRecorder()

		authenticator := auth.NewAuthenticator("foo", "bar")

		authenticator.CreateSession(res)

		actualCookie := res.Header().Get("Set-Cookie")
		if actualCookie != expectedCookie {
			t.Errorf("CreateSession() added unexpected Set-Cookie header: got '%v' want '%v'", actualCookie, expectedCookie)
		}
	})
}

// t.Run("adds unauthenticated user when no authentication is provided", func(t *testing.T) {
// 	actualContext = nil
// 	res := httptest.NewRecorder()
// 	req, err := http.NewRequest("", "", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	middleware.ServeHTTP(res, req)

// 	actualUser := auth.GetUserFromContext(actualContext)
// 	if auth.IsAuthenticated(actualUser) != false {
// 		t.Error("actualUser.IsAuthenticated() has unexpected value: got true want false")
// 	}
// })

// t.Run("adds authenticated user when valid authentication is provided", func(t *testing.T) {
// 	actualContext = nil
// 	res := httptest.NewRecorder()
// 	req, err := http.NewRequest("", "", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.SetBasicAuth(correctUsername, correctPassword)

// 	middleware.ServeHTTP(res, req)
// 	actualUser := auth.GetUserFromContext(actualContext)

// 	if auth.IsAuthenticated(actualUser) != true {
// 		t.Error("actualUser.IsAuthenticated() has unexpected value: got false want true")
// 	}
// })
// }
