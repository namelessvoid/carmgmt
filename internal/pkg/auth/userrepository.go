package auth

import (
	"context"
	"errors"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

// A UserRepository is used to persist and query User objects.
type UserRepository interface {
	CreateUser(ctx context.Context, user User) (User, error)
	FindUserByUsername(ctx context.Context, username string) (User, error)
}

// ErrUserNotFound indicates that a find query on the UserRepository did not find
// any matching user.
var ErrUserNotFound = errors.New("User not found")

// ErrUserAlradyExists is returned by UserRepository.CreateUser() if the provided
// User.Username is not unique
var ErrUserAlreadyExists = errors.New("User already exists")

type appengineUserRepository struct {
	client *datastore.Client
}

// NewAppengineUserRepository creates a new UserRepository which stores its
// data in a google Datastore.
func NewAppengineUserRepository(client *datastore.Client) *appengineUserRepository {
	return &appengineUserRepository{client: client}
}

// TODO: Add transaction to avoid query races
func (repo *appengineUserRepository) CreateUser(ctx context.Context, user User) (User, error) {
	key := datastore.IncompleteKey("User", nil)

	_, err := repo.FindUserByUsername(ctx, user.Username)
	if err == nil {
		return User{}, ErrUserAlreadyExists
	} else if err != ErrUserNotFound {
		return User{}, err
	}

	key, err = repo.client.Put(ctx, key, &user)
	if err != nil {
		return User{}, err
	}

	user.ID = key.ID
	return user, nil
}

func (repo *appengineUserRepository) FindUserByUsername(ctx context.Context, username string) (User, error) {
	user := User{}

	query := datastore.NewQuery("User").Filter("Username =", username)
	it := repo.client.Run(ctx, query)
	key, err := it.Next(&user)

	if err == iterator.Done {
		return User{}, ErrUserNotFound
	}

	if err != nil {
		return User{}, err
	}

	user.ID = key.ID
	return user, nil
}
