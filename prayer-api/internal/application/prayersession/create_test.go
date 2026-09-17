package prayersession

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	domainuser "prayer-api/internal/domain/user"
)

type createPrayerSessionRepositoryStub struct {
	savedSession *domain.PrayerSession
	saveErr      error
}

func (r *createPrayerSessionRepositoryStub) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	return nil, nil
}

func (r *createPrayerSessionRepositoryStub) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {
	return nil, nil
}

func (r *createPrayerSessionRepositoryStub) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	r.savedSession = session
	return r.saveErr
}

func (r *createPrayerSessionRepositoryStub) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (r *createPrayerSessionRepositoryStub) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	return nil
}

type createUserRepositoryStub struct {
	user      *domainuser.User
	findErr   error
	externalID string
}

func (r *createUserRepositoryStub) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {
	r.externalID = externalID

	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.user, nil
}

func (r *createUserRepositoryStub) Save(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

func (r *createUserRepositoryStub) FindByEmail(
	ctx context.Context,
	email string,
) (*domainuser.User, error) {
	return nil, nil
}

func (r *createUserRepositoryStub) Update(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

func TestService_Create(t *testing.T) {
	t.Parallel()

	validActorID := identity.UserID("user-1")
	validGroupID := identity.PrayerGroupID("group-1")
	validDate := time.Date(
		2026,
		9,
		18,
		0,
		0,
		0,
		0,
		time.UTC,
	)

	validCommand := CreateCommand{
		ActorID:       validActorID,
		PrayerGroupID: validGroupID,
		Title:         "Morning Prayer",
		Date:          validDate,
		Time:          "07:00",
		Duration:      60,
		PrayerPointIDs: []identity.PrayerPointID{
			"point-1",
			"point-2",
		},
	}

	t.Run("creates a prayer session successfully", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID:             validActorID,
			PrayerGroupIDs: []identity.PrayerGroupID{validGroupID},
		}

		userRepo := &createUserRepositoryStub{
			user: user,
		}

		sessionRepo := &createPrayerSessionRepositoryStub{}

		service := &Service{
			users:    userRepo,
			sessions: sessionRepo,
		}

		session, err := service.Create(
			context.Background(),
			validCommand,
		)

		require.NoError(t, err)
		require.NotNil(t, session)

		assert.Equal(t, validGroupID, session.PrayerGroupID)
		assert.Equal(t, "Morning Prayer", session.Title)
		assert.Equal(t, validDate, session.Date)
		assert.Equal(t, "07:00", session.Time)
		assert.Equal(t, 60, session.Duration)
		assert.Equal(
			t,
			validCommand.PrayerPointIDs,
			session.PrayerPointIDs,
		)

		require.NotNil(t, sessionRepo.savedSession)
		assert.Equal(
			t,
			session,
			sessionRepo.savedSession,
		)
	})

	t.Run("requires prayer group ID", func(t *testing.T) {
		t.Parallel()

		cmd := validCommand
		cmd.PrayerGroupID = ""

		service := &Service{}

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidGroupID)
	})

	t.Run("requires title", func(t *testing.T) {
		t.Parallel()

		cmd := validCommand
		cmd.Title = ""

		service := &Service{}

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrTitleRequired)
	})

	t.Run("rejects whitespace-only title", func(t *testing.T) {
		t.Parallel()

		cmd := validCommand
		cmd.Title = "   "

		service := &Service{}

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrTitleRequired)
	})

	t.Run("requires date", func(t *testing.T) {
		t.Parallel()

		cmd := validCommand
		cmd.Date = time.Time{}

		service := &Service{}

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidDate)
	})

	t.Run("requires time", func(t *testing.T) {
		t.Parallel()

		cmd := validCommand
		cmd.Time = ""

		service := &Service{}

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrTimeRequired)
	})

	t.Run("rejects whitespace-only time", func(t *testing.T) {
		t.Parallel()

		cmd := validCommand
		cmd.Time = "   "

		service := &Service{}

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrTimeRequired)
	})

	t.Run("requires positive duration", func(t *testing.T) {
		t.Parallel()

		cmd := validCommand
		cmd.Duration = 0

		service := &Service{}

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidDuration)
	})

	t.Run("rejects negative duration", func(t *testing.T) {
		t.Parallel()

		cmd := validCommand
		cmd.Duration = -10

		service := &Service{}

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidDuration)
	})

	t.Run("returns unknown when user lookup fails", func(t *testing.T) {
		t.Parallel()

		userRepo := &createUserRepositoryStub{
			findErr: errors.New("database failure"),
		}

		sessionRepo := &createPrayerSessionRepositoryStub{}

		service := &Service{
			users:    userRepo,
			sessions: sessionRepo,
		}

		session, err := service.Create(
			context.Background(),
			validCommand,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrUnknown)
	})

	t.Run("returns user not found when user does not exist", func(t *testing.T) {
		t.Parallel()

		userRepo := &createUserRepositoryStub{
			user: nil,
		}

		sessionRepo := &createPrayerSessionRepositoryStub{}

		service := &Service{
			users:    userRepo,
			sessions: sessionRepo,
		}

		session, err := service.Create(
			context.Background(),
			validCommand,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrUserNotFound)
	})

	t.Run("returns forbidden when user has no group access", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID:             validActorID,
			PrayerGroupIDs: []identity.PrayerGroupID{
				identity.PrayerGroupID("different-group"),
			},
		}

		userRepo := &createUserRepositoryStub{
			user: user,
		}

		sessionRepo := &createPrayerSessionRepositoryStub{}

		service := &Service{
			users:    userRepo,
			sessions: sessionRepo,
		}

		session, err := service.Create(
			context.Background(),
			validCommand,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrForbidden)
	})

	t.Run("trims title before creating session", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID:             validActorID,
			PrayerGroupIDs: []identity.PrayerGroupID{validGroupID},
		}

		userRepo := &createUserRepositoryStub{
			user: user,
		}

		sessionRepo := &createPrayerSessionRepositoryStub{}

		service := &Service{
			users:    userRepo,
			sessions: sessionRepo,
		}

		cmd := validCommand
		cmd.Title = "  Morning Prayer  "

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		require.NoError(t, err)
		require.NotNil(t, session)

		assert.Equal(t, "Morning Prayer", session.Title)
	})

	t.Run("trims time before creating session", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID:             validActorID,
			PrayerGroupIDs: []identity.PrayerGroupID{validGroupID},
		}

		userRepo := &createUserRepositoryStub{
			user: user,
		}

		sessionRepo := &createPrayerSessionRepositoryStub{}

		service := &Service{
			users:    userRepo,
			sessions: sessionRepo,
		}

		cmd := validCommand
		cmd.Time = "  07:00  "

		session, err := service.Create(
			context.Background(),
			cmd,
		)

		require.NoError(t, err)
		require.NotNil(t, session)

		assert.Equal(t, "07:00", session.Time)
	})

	t.Run("passes actor ID as external ID to user repository", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID:             validActorID,
			PrayerGroupIDs: []identity.PrayerGroupID{validGroupID},
		}

		userRepo := &createUserRepositoryStub{
			user: user,
		}

		service := &Service{
			users:    userRepo,
			sessions: &createPrayerSessionRepositoryStub{},
		}

		cmd := validCommand
		cmd.ActorID = identity.UserID("  external-user-1  ")

		_, err := service.Create(
			context.Background(),
			cmd,
		)

		require.NoError(t, err)
		assert.Equal(t, "external-user-1", userRepo.externalID)
	})

	t.Run("returns repository error when save fails", func(t *testing.T) {
		t.Parallel()

		user := &domainuser.User{
			ID:             validActorID,
			PrayerGroupIDs: []identity.PrayerGroupID{validGroupID},
		}

		userRepo := &createUserRepositoryStub{
			user: user,
		}

		saveErr := errors.New("save failed")

		sessionRepo := &createPrayerSessionRepositoryStub{
			saveErr: saveErr,
		}

		service := &Service{
			users:    userRepo,
			sessions: sessionRepo,
		}

		session, err := service.Create(
			context.Background(),
			validCommand,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, saveErr)
	})
}