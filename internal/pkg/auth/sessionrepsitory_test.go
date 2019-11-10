package auth_test

import (
	"context"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/namelessvoid/carmgmt/internal/pkg/auth"
)

func Test_AppengineSessionRepsitory(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	var repo auth.SessionRepository

	if testing.Short() {
		t.Skip("Skip integration test in short mode")
	}

	ctx := context.Background()
	dsClient, err := datastore.NewClient(ctx, "integration-test")
	if err != nil {
		t.Fatal(err)
	}
	defer dsClient.Close()

	t.Run("returns ErrSessionNotFound when session does not exist", func(t *testing.T) {
		repo = auth.NewAppengineSessionRepository(ctx, dsClient)

		_, err = repo.FindSession("foo")

		if err != auth.ErrSessionNotFound {
			t.Errorf("FindSession() returned unexpected error: got %v want %v", err, auth.ErrSessionNotFound)
		}
	})

	t.Run("can store sessions and retrieve them", func(t *testing.T) {
		repo = auth.NewAppengineSessionRepository(ctx, dsClient)

		token := strconv.Itoa(rand.Int())
		userID := rand.Int()
		expectedSession := auth.Session{Token: token, UserID: userID}
		err = repo.CreateSession(expectedSession)

		if err != nil {
			t.Errorf("CreateSession() returned unexpected error: '%v'", err)
		}

		actualSession, err := repo.FindSession(expectedSession.Token)
		if err != nil {
			t.Errorf("FindSession() returned unexpected error: '%v'", err)
		}

		if actualSession.Token != expectedSession.Token {
			t.Errorf("FindSession() returned Session with unexpected Token: got '%v' want '%v'", actualSession.Token, expectedSession.Token)
		}

		if actualSession.UserID != expectedSession.UserID {
			t.Errorf("FindSession() returned Session with unexpected UserID: got '%v' want '%v'", actualSession.UserID, expectedSession.UserID)
		}
	})
}
