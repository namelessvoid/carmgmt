package auth_test

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/namelessvoid/carmgmt/internal/pkg/auth"
	auth_mock "github.com/namelessvoid/carmgmt/internal/pkg/auth/mock"
	"golang.org/x/crypto/bcrypt"
)

func Test_Authenticator_LoginViaCredentials(t *testing.T) {
	const (
		correctUsername = "correctUsername"
		correctPassword = "correctPassword"
	)
	correctPasswordHash, _ := bcrypt.GenerateFromPassword([]byte(correctPassword), bcrypt.DefaultCost)

	t.Run("returns valid token if credentials are correct", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		req, err := http.NewRequest("", "", nil)
		if err != nil {
			t.Fatal(err)
		}
		req.SetBasicAuth(correctUsername, correctPassword)

		userRepository := auth_mock.NewMockUserRepository(mockCtrl)
		userRepository.
			EXPECT().
			FindUserByUsername(gomock.Any(), correctUsername).
			Return(auth.User{ID: 2010, Username: correctUsername, Password: string(correctPasswordHash)}, nil)
		authenticator := auth.NewAuthenticator(userRepository, nil)

		token, err := authenticator.LoginViaCredentials(req)

		if err != nil {
			t.Errorf("LoginViaCredentials() returned unexpected error: got '%v' want no error", err)
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
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			req, err := http.NewRequest("", "", nil)
			if err != nil {
				t.Fatal(err)
			}
			req.SetBasicAuth(run.username, run.password)

			userRepository := auth_mock.NewMockUserRepository(mockCtrl)
			userRepository.
				EXPECT().
				FindUserByUsername(gomock.Any(), correctUsername).
				Return(auth.User{ID: 2010, Username: correctUsername, Password: string(correctPasswordHash)}, nil).
				AnyTimes()
			userRepository.
				EXPECT().
				FindUserByUsername(gomock.Any(), gomock.Any()).
				Return(auth.User{}, auth.ErrUserNotFound).
				AnyTimes()
			authenticator := auth.NewAuthenticator(userRepository, nil)

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
	t.Run("saves new token and adds it to the response", func(t *testing.T) {
		mockCtrl := gomock.NewController(t)
		defer mockCtrl.Finish()

		expectedCookieRegex := "FLEETMGMT_SESSION=\\d+; Path=/; Domain=localhost; HttpOnly"

		res := httptest.NewRecorder()

		sessionRepository := auth_mock.NewMockSessionRepository(mockCtrl)
		sessionRepository.EXPECT().CreateSession(gomock.Any())

		authenticator := auth.NewAuthenticator(nil, sessionRepository)

		err := authenticator.CreateSession(res)

		if err != nil {
			t.Errorf("CreateSession() returned unexpected error: %v", err)
		}

		actualCookie := res.Header().Get("Set-Cookie")
		if matched, _ := regexp.MatchString(expectedCookieRegex, actualCookie); !matched {
			t.Errorf("CreateSession() added unexpected Set-Cookie header: got '%v' did not match regexp '%v'", actualCookie, expectedCookieRegex)
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
