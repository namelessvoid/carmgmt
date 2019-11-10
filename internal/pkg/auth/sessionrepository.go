package auth

import (
	"context"
	"errors"

	"cloud.google.com/go/datastore"
	"google.golang.org/api/iterator"
)

//go:generate mockgen -source sessionrepository.go -destination=./mock/sessionrepositorymock.go -package=auth_mock

// The SessionRepository stores sessions e.g. created by an Authenticator.
type SessionRepository interface {
	CreateSession(session Session) error
	FindSession(token string) (Session, error)
}

type appengineSessionRepository struct {
	ctx    context.Context
	client *datastore.Client
}

// ErrSessionNotFound indicates that SessionRepository.FindSession() did not
// find a valid session.
var ErrSessionNotFound = errors.New("Session not found")

// NewAppengineSessionRepository creates a new SessionRepository which stores its
// data in a google Datastore.
func NewAppengineSessionRepository(ctx context.Context, client *datastore.Client) *appengineSessionRepository {
	return &appengineSessionRepository{ctx: ctx, client: client}
}

func (repo *appengineSessionRepository) CreateSession(session Session) error {
	key := datastore.IncompleteKey("Session", nil)
	_, err := repo.client.Put(repo.ctx, key, &session)
	return err
}

func (repo *appengineSessionRepository) FindSession(token string) (Session, error) {
	var session = Session{}

	query := datastore.NewQuery("Session").Filter("Token = ", token)
	it := repo.client.Run(repo.ctx, query)
	_, err := it.Next(&session)

	if err == iterator.Done {
		return session, ErrSessionNotFound
	}

	if err != nil {
		return Session{}, err
	}

	return session, nil
}
