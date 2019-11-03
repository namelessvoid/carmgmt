package auth

// The User represents authentication information about a user.
type User struct {
	isAuthenticated bool
}

// IsAuthenticated returns if the user has been authenticated.
// This method can be used with the AuthorizationMiddleware
func IsAuthenticated(u User) bool {
	return u.isAuthenticated
}
