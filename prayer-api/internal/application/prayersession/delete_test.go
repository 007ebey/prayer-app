package prayersession

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	domainuser "prayer-api/internal/domain/user"
)

type deletePrayerSessionRepositoryStub struct {
	session    *domain.PrayerSession
	findErr    error
	deleteErr  error
	deletedID  identity.PrayerSessionID
	deleteCall bool
}

func (r *deletePrayerSessionRepositoryStub) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.session, nil
}

func (r *deletePrayerSessionRepositoryStub) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {
	return nil, nil
}

func (r *deletePrayerSessionRepositoryStub) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (r *deletePrayerSessionRepositoryStub) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (r *deletePrayerSessionRepositoryStub) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	r.deleteCall = true
	r.deletedID = id

	return r.deleteErr
}

type deleteUserRepositoryStub struct {
	user       *domainuser.User
	findErr    error
	externalID string
}

func (r *deleteUserRepositoryStub) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {
	r.externalID = externalID

	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.user, nil
}

func (r *deleteUserRepositoryStub) Save(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

func (r *deleteUserRepositoryStub) FindByEmail(
	ctx context.Context,
	email string,
) (*domainuser.User, error) {
	return nil, nil
}

func (r *deleteUserRepositoryStub) Update(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

func TestService_Delete(t *testing.T) {
	t.Parallel()

	sessionID := identity.PrayerSessionID("session-1")
	groupID := identity.PrayerGroupID("group-1")
	userID := identity.UserID("user-1")

	session := &domain.PrayerSession{
		ID:            sessionID,
		PrayerGroupID: groupID,
		Title:         "Morning Prayer",
	}

	t.Run("deletes prayer session when user has access", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID:             userID,
			PrayerGroupIDs: []identity.PrayerGroupID{groupID},
		}

		sessionRepo := &deletePrayerSessionRepositoryStub{
			session: session,
		}

		userRepo := &deleteUserRepositoryStub{
			user: user,
		}

		service := &Service{
			sessions: sessionRepo,
			users:    userRepo,
		}

		err := service.Delete(
			context.Background(),
			userID,
			sessionID,
		)

		require.NoError(t, err)

		assert.True(t, sessionRepo.deleteCall)
		assert.Equal(t, sessionID, sessionRepo.deletedID)
		assert.Equal(t, "user-1", userRepo.externalID)
	})

	t.Run("returns repository error when finding session fails", func(t *testing.T) {
		t.Parallel()

		findErr := errors.New("find session failed")

		sessionRepo := &deletePrayerSessionRepositoryStub{
			findErr: findErr,
		}

		service := &Service{
			sessions: sessionRepo,
			users:    &deleteUserRepositoryStub{},
		}

		err := service.Delete(
			context.Background(),
			userID,
			sessionID,
		)

		assert.ErrorIs(t, err, findErr)
		assert.False(t, sessionRepo.deleteCall)
	})

	t.Run("returns not found when session does not exist", func(t *testing.T) {
		t.Parallel()

		sessionRepo := &deletePrayerSessionRepositoryStub{
			session: nil,
		}

		service := &Service{
			sessions: sessionRepo,
			users:    &deleteUserRepositoryStub{},
		}

		err := service.Delete(
			context.Background(),
			userID,
			sessionID,
		)

		assert.ErrorIs(t, err, ErrNotFound)
		assert.False(t, sessionRepo.deleteCall)
	})

	t.Run("returns error when finding user fails", func(t *testing.T) {
		t.Parallel()

		findErr := errors.New("find user failed")

		sessionRepo := &deletePrayerSessionRepositoryStub{
			session: session,
		}

		userRepo := &deleteUserRepositoryStub{
			findErr: findErr,
		}

		service := &Service{
			sessions: sessionRepo,
			users:    userRepo,
		}

		err := service.Delete(
			context.Background(),
			userID,
			sessionID,
		)

		assert.ErrorIs(t, err, findErr)
		assert.False(t, sessionRepo.deleteCall)
	})

	t.Run("returns user not found when user does not exist", func(t *testing.T) {
		t.Parallel()

		sessionRepo := &deletePrayerSessionRepositoryStub{
			session: session,
		}

		userRepo := &deleteUserRepositoryStub{
			user: nil,
		}

		service := &Service{
			sessions: sessionRepo,
			users:    userRepo,
		}

		err := service.Delete(
			context.Background(),
			userID,
			sessionID,
		)

		assert.ErrorIs(t, err, ErrUserNotFound)
		assert.False(t, sessionRepo.deleteCall)
	})

	t.Run("returns forbidden when user has no group access", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID: userID,
			PrayerGroupIDs: []identity.PrayerGroupID{
				identity.PrayerGroupID("different-group"),
			},
		}

		sessionRepo := &deletePrayerSessionRepositoryStub{
			session: session,
		}

		userRepo := &deleteUserRepositoryStub{
			user: user,
		}

		service := &Service{
			sessions: sessionRepo,
			users:    userRepo,
		}

		err := service.Delete(
			context.Background(),
			userID,
			sessionID,
		)

		assert.ErrorIs(t, err, ErrForbidden)
		assert.False(t, sessionRepo.deleteCall)
	})

	t.Run("returns error when deleting session fails", func(t *testing.T) {
		t.Parallel()

		deleteErr := errors.New("delete session failed")

		user := &domainuser.User{
			ID:             userID,
			PrayerGroupIDs: []identity.PrayerGroupID{groupID},
		}

		sessionRepo := &deletePrayerSessionRepositoryStub{
			session:   session,
			deleteErr: deleteErr,
		}

		userRepo := &deleteUserRepositoryStub{
			user: user,
		}

		service := &Service{
			sessions: sessionRepo,
			users:    userRepo,
		}

		err := service.Delete(
			context.Background(),
			userID,
			sessionID,
		)

		assert.ErrorIs(t, err, deleteErr)
		assert.True(t, sessionRepo.deleteCall)
		assert.Equal(t, sessionID, sessionRepo.deletedID)
	})

	t.Run("trims external user ID before lookup", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID:             userID,
			PrayerGroupIDs: []identity.PrayerGroupID{groupID},
		}

		sessionRepo := &deletePrayerSessionRepositoryStub{
			session: session,
		}

		userRepo := &deleteUserRepositoryStub{
			user: user,
		}

		service := &Service{
			sessions: sessionRepo,
			users:    userRepo,
		}

		err := service.Delete(
			context.Background(),
			identity.UserID("  user-1  "),
			sessionID,
		)

		require.NoError(t, err)
		assert.Equal(t, "user-1", userRepo.externalID)
	})
}