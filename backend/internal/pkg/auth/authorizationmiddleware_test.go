package auth_test

import (
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
	}}

	for _, run := range tests {
		t.Run(run.name, func(t *testing.T) {
			authorizationFunction := func(u auth.UserInfo) bool {
				return run.authorized
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
			req = req.WithContext(auth.AddUserToTestContext(req.Context(), true))

			middleware := auth.AuthorizationMiddleware(authorizationFunction)(next)

			middleware.ServeHTTP(res, req)

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
