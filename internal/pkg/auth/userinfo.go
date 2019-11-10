package auth

// UserInfo contains authentication information about a user.
// This must not be confused with the User object stored by UserRepository.
type UserInfo struct {
	UserID          int64
	Username        string
	IsAuthenticated bool
}

// IsAuthenticated returns if the user has been authenticated.
// This method can be used with the AuthorizationMiddleware
func IsAuthenticated(info UserInfo) bool {
	return info.IsAuthenticated
}
