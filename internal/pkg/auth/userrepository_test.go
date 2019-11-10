package auth_test

import (
	"context"
	"math/rand"
	"reflect"
	"strconv"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/namelessvoid/carmgmt/internal/pkg/auth"
)

func Test_AppengineUserRepository(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip integration test in short mode")
	}

	rand.Seed(time.Now().UnixNano())

	// Ensure AppengineUserRepository implements UserRepository
	var repo auth.UserRepository

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, "integration-test")
	if err != nil {
		t.Fatal(err)
	}
	defer dsClient.Close()

	t.Run("FindUserByUsername() returns ErrUserNotFound when user does not exist", func(t *testing.T) {
		expectedUser := auth.User{}

		repo = auth.NewAppengineUserRepository(dsClient)

		actualUser, err := repo.FindUserByUsername(ctx, "invalid username")

		if err != auth.ErrUserNotFound {
			t.Errorf("FindUserByUsername() returned unexpected error: got '%v' want '%v'", err, auth.ErrUserNotFound)
		}

		if !reflect.DeepEqual(actualUser, expectedUser) {
			t.Errorf("FindUserByUsername() returned unexpected user: got '%v' want '%v'", actualUser, expectedUser)
		}
	})

	t.Run("can store user and retrieve her", func(t *testing.T) {
		repo = auth.NewAppengineUserRepository(dsClient)

		// Create
		username := strconv.FormatInt(rand.Int63(), 10)
		user := auth.User{ID: -1, Username: username, Password: "SomePasswordThing"}
		insertedUser, err := repo.CreateUser(ctx, user)

		if err != nil {
			t.Errorf("CreateUser() returned unexpected error: got '%v' want no error", err)
		}

		if insertedUser.ID == -1 {
			t.Errorf("CreateUser() did not assign User.ID correctly: got '%v' want id != -1", insertedUser.ID)
		}

		expectedUser := user
		expectedUser.ID = insertedUser.ID
		if !reflect.DeepEqual(insertedUser, expectedUser) {
			t.Errorf("CreateUser() returned unexpected user: got '%v' want '%v'", insertedUser, expectedUser)
		}

		// Query user.
		expectedUser = insertedUser
		queriedUser, err := repo.FindUserByUsername(ctx, expectedUser.Username)
		if err != nil {
			t.Errorf("FindUserByUsername() returned unexpected error: got '%v' want no error", err)
		}

		if !reflect.DeepEqual(queriedUser, expectedUser) {
			t.Errorf("FindUserByUsername() returned unexpected user: got '%v' want '%v'", queriedUser, expectedUser)
		}
	})

	t.Run("CreateUser() returns ErrUserAlreadyExists if User.Username is not unique", func(t *testing.T) {
		repo = auth.NewAppengineUserRepository(dsClient)

		user := auth.User{Username: "DuplicateUserName", Password: "DoesNotMatter"}

		repo.CreateUser(ctx, user)
		_, err := repo.CreateUser(ctx, user)

		if err != auth.ErrUserAlreadyExists {
			t.Errorf("CreateUser() returned unexpected error: got '%v' want '%v'", err, auth.ErrUserAlreadyExists)
		}
	})
}
