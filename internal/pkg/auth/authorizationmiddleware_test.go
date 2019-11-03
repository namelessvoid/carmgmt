package auth_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/namelessvoid/carmgmt/internal/pkg/auth"
)

func Test_AuthorizationMiddleware(t *testing.T) {
	tests := []struct {
		name               string
		authorized         bool
		authorizationError error
		expectsHandlerCall bool
		expectedStatusCode int
	}{{
		name:               "calls next handler if authorization function returns true",
		authorized:         true,
		authorizationError: nil,
		expectsHandlerCall: true,
		expectedStatusCode: http.StatusOK,
	}, {
		name:               "returns fobidden if authorization function returns false",
		authorized:         false,
		authorizationError: nil,
		expectsHandlerCall: false,
		expectedStatusCode: http.StatusForbidden,
	}, {
		name:               "returns InternalServerError if authorization function returns error",
		authorized:         false,
		authorizationError: errors.New("Some error"),
		expectsHandlerCall: false,
		expectedStatusCode: http.StatusInternalServerError,
	}}

	for _, run := range tests {
		t.Run(run.name, func(t *testing.T) {
			authorizationFunction := func(u auth.User) (bool, error) {
				return run.authorized, run.authorizationError
			}

			handlerCalled := false
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				handlerCalled = true
			})

			res := httptest.NewRecorder()
			req, err := http.NewRequest("", "", nil)
			if err != nil {
				t.Fatal(err)
			}

			sut := auth.AuthorizationMiddleware(authorizationFunction)(next)
			// We need to wrap into AuthenticationMiddleware to fill in a User into req.Context()
			reqChain := auth.AuthenticationMiddleware("username", "password")(sut)

			reqChain.ServeHTTP(res, req)

			if res.Code != run.expectedStatusCode {
				t.Errorf("AuthorizationMiddleware returned unexpected status code: got %v want %v", res.Code, run.expectedStatusCode)
			}

			if handlerCalled != run.expectsHandlerCall {
				text := ""
				if run.expectsHandlerCall == true {
					text = "AuthorizationMiddleware did not call next handler"
				} else {
					text = "expected request to be aborted but AuthorizationMiddleware forwarded request to next handler"
				}
				t.Error(text)
			}
		})
	}
}
