package auth

import (
	"fmt"
	"net/http"
	"strings"
)

type user struct {
	username string
	password string
}

type Token string
type tokens map[Token]bool

func (t Token) IsEmpty() bool {
	return len(t) == 0
}

type session string
type sessions map[session]bool

// The Authenticator can be used to login via cookie or user credentials,
// to create a new cookie based session and to verify tokens created by
// the login methods.
type Authenticator interface {
	LoginViaSession(r *http.Request) (Token, error)
	LoginViaCredentials(r *http.Request) (Token, error)
	CreateSession(w http.ResponseWriter)
	VerifyToken(r *http.Request) bool
}

type authenticator struct {
	user     user
	sessions sessions
	tokens   tokens
}

func NewAuthenticator(username, password string) *authenticator {
	u := user{username: username, password: password}
	return &authenticator{user: u, sessions: sessions{}, tokens: tokens{}}
}

func (a *authenticator) LoginViaSession(r *http.Request) (Token, error) {
	c, err := r.Cookie("FLEETMGMT_SESSION")
	if err == http.ErrNoCookie {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	_, exists := a.sessions[session(c.Value)]
	if exists {
		return a.generateToken(), nil
	}

	return "", nil
}

func (a *authenticator) LoginViaCredentials(r *http.Request) (Token, error) {
	username, password, ok := r.BasicAuth()

	if !ok {
		return "", nil
	}

	fmt.Println("via credentials: ", a.user.username, a.user.password)
	if a.user.username == username && a.user.password == password {
		return a.generateToken(), nil
	}

	return "", nil
}

func (a *authenticator) CreateSession(w http.ResponseWriter) {
	s := session("somesession")
	a.sessions[s] = true
	c := &http.Cookie{Name: "FLEETMGMT_SESSION", Value: string(s), HttpOnly: true, Path: "/", Domain: "localhost"}
	fmt.Println(c.Secure)
	http.SetCookie(w, c)
}

func (a *authenticator) VerifyToken(r *http.Request) bool {
	authHeader := r.Header.Get("Authorization")
	authInfo := strings.Split(authHeader, " ")
	if len(authInfo) != 2 {
		return false
	}

	if authInfo[0] != "Bearer" {
		return false
	}

	t := Token(authInfo[1])
	_, exists := a.tokens[t]
	return exists
}

func (a *authenticator) generateToken() Token {
	a.tokens["foo"] = true
	return "foo"
}
