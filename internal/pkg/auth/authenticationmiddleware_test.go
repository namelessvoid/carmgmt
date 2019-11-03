package auth_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/namelessvoid/carmgmt/internal/pkg/auth"
	auth_mock "github.com/namelessvoid/carmgmt/internal/pkg/auth/mock"
)

func Test_AuthenticationMiddleware(t *testing.T) {
	var actualContext context.Context
	next := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		actualContext = req.Context()
	})

	res := httptest.NewRecorder()
	req, err := http.NewRequest("", "", nil)
	if err != nil {
		t.Fatal(err)
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	authenticator := auth_mock.NewMockAuthenticator(mockCtrl)
	authenticator.EXPECT().VerifyToken(req).Return(true)

	middleware := auth.AuthenticationMiddleware(authenticator)(next)

	middleware.ServeHTTP(res, req)

	actualUser := auth.GetUserFromContext(actualContext)
	if auth.IsAuthenticated(actualUser) != true {
		t.Error("actualUser.IsAuthenticated() has unexpected value: got false want true")
	}
}
