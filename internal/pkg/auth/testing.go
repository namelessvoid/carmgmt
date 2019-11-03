package auth

import "context"

// AddUserToTestContext adds a user with given properties to the context to
// fake the AuthenticationMiddleware in test
func AddUserToTestContext(ctx context.Context, isAuthenticated bool) context.Context {
	u := User{isAuthenticated: isAuthenticated}
	return addUserToContext(ctx, u)
}
