package auth

// The User represents authentication information about a user.
type User struct {
	isAuthenticated bool
}

// IsAuthenticated returns if the user has been authenticated.
func (u User) IsAuthenticated() bool {
	return u.isAuthenticated
}
