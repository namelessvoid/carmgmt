package auth

import (
	"context"
	"net/http"
	"strings"

	oidc "github.com/coreos/go-oidc"
)

//go:generate mockgen -source authenticator.go -destination=./mock/authenticatormock.go -package=auth_mock

// The Authenticator can be used to login via cookie or user credentials,
// to create a new cookie based session and to verify tokens created by
// the login methods.
type Authenticator interface {
	VerifyToken(r *http.Request) bool
}

type authenticator struct{}

func NewAuthenticator() *authenticator {
	return &authenticator{}
}

func (a *authenticator) VerifyToken(r *http.Request) bool {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://dev-fleetmgmt.eu.auth0.com/")
	if err != nil {
		println(err)
		return false
	}

	verifier := provider.Verifier(&oidc.Config{ClientID: "Fleet Management - Local"})

	authHeader := r.Header.Get("Authorization")
	authInfo := strings.Split(authHeader, " ")
	if len(authInfo) != 2 {
		return false
	}

	if authInfo[0] != "Bearer" {
		return false
	}

	rawIDToken := authInfo[1]
	_, err = verifier.Verify(ctx, rawIDToken)

	if err != nil {
		println(err.Error())
		return false
	}

	return true
}
