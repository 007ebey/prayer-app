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

// ---------------------------------------------------------
// PrayerSession repository stub
// ---------------------------------------------------------

type updatePrayerSessionRepositoryStub struct {
	session    *domain.PrayerSession
	findErr    error
	updateErr  error
	saved      *domain.PrayerSession
	findCalled bool
	updateCalled bool
}

func (r *updatePrayerSessionRepositoryStub) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	r.findCalled = true

	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.session, nil
}

func (r *updatePrayerSessionRepositoryStub) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {
	return nil, nil
}

func (r *updatePrayerSessionRepositoryStub) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (r *updatePrayerSessionRepositoryStub) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	r.updateCalled = true
	r.saved = session

	return r.updateErr
}

func (r *updatePrayerSessionRepositoryStub) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	return nil
}

// ---------------------------------------------------------
// User repository stub
// ---------------------------------------------------------

type updateUserRepositoryStub struct {
	user       *domainuser.User
	findErr    error
	findCalled bool
	externalID string
}

func (r *updateUserRepositoryStub) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {
	r.findCalled = true
	r.externalID = externalID

	if r.findErr != nil {
		return nil, r.findErr
	}

	return r.user, nil
}

func (r *updateUserRepositoryStub) Save(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

func (r *updateUserRepositoryStub) FindByEmail(
	ctx context.Context,
	email string,
) (*domainuser.User, error) {
	return nil, nil
}

func (r *updateUserRepositoryStub) Update(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

// ---------------------------------------------------------
// Success
// ---------------------------------------------------------

func TestService_Update_Success(t *testing.T) {
	t.Parallel()

	actorID := identity.UserID("user-1")
	groupID := identity.PrayerGroupID("group-1")
	sessionID := identity.PrayerSessionID("session-1")

	oldDate := time.Date(
		2026, 9, 20,
		0, 0, 0, 0,
		time.UTC,
	)

	newDate := time.Date(
		2026, 9, 21,
		0, 0, 0, 0,
		time.UTC,
	)

	existing := &domain.PrayerSession{
		ID:            sessionID,
		PrayerGroupID: groupID,
		Title:         "Morning Prayer",
		Date:          oldDate,
		Time:          "06:00",
		Duration:      30,
		PrayerPointIDs: []identity.PrayerPointID{
			identity.PrayerPointID("point-1"),
		},
	}

	user := &domainuser.User{
		ID: actorID,
		PrayerGroupIDs: []identity.PrayerGroupID{
			groupID,
		},
	}

	sessions := &updatePrayerSessionRepositoryStub{
		session: existing,
	}

	users := &updateUserRepositoryStub{
		user: user,
	}

	service := NewService(sessions, users)

	result, err := service.Update(
		context.Background(),
		UpdateCommand{
			ActorID:   actorID,
			SessionID: sessionID,
			Title:     "Evening Prayer",
			Date:      newDate,
			Time:      "19:30",
			Duration:  45,
			PrayerPointIDs: []identity.PrayerPointID{
				identity.PrayerPointID("point-2"),
				identity.PrayerPointID("point-3"),
			},
		},
	)

	require.NoError(t, err)
	require.NotNil(t, result)
	require.NotNil(t, sessions.saved)

	assert.Equal(t, sessionID, result.ID)
	assert.Equal(t, groupID, result.PrayerGroupID)
	assert.Equal(t, "Evening Prayer", result.Title)
	assert.Equal(t, newDate, result.Date)
	assert.Equal(t, "19:30", result.Time)
	assert.Equal(t, 45, result.Duration)

	assert.Equal(
		t,
		[]identity.PrayerPointID{
			identity.PrayerPointID("point-2"),
			identity.PrayerPointID("point-3"),
		},
		result.PrayerPointIDs,
	)

	assert.True(t, sessions.findCalled)
	assert.True(t, sessions.updateCalled)
	assert.True(t, users.findCalled)
	assert.Equal(t, "user-1", users.externalID)
}

// ---------------------------------------------------------
// Session not found
// ---------------------------------------------------------

func TestService_Update_SessionNotFound(t *testing.T) {
	t.Parallel()

	sessions := &updatePrayerSessionRepositoryStub{
		session: nil,
	}

	users := &updateUserRepositoryStub{}

	service := NewService(sessions, users)

	result, err := service.Update(
		context.Background(),
		UpdateCommand{
			ActorID:   identity.UserID("user-1"),
			SessionID: identity.PrayerSessionID("session-1"),
			Title:     "Evening Prayer",
			Date:      time.Now(),
			Time:      "19:30",
			Duration:  45,
		},
	)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrNotFound)

	assert.True(t, sessions.findCalled)

	// User lookup should not happen.
	assert.False(t, users.findCalled)

	// Update should not happen.
	assert.False(t, sessions.updateCalled)
}

// ---------------------------------------------------------
// Session repository error
// ---------------------------------------------------------

func TestService_Update_FindSessionError(t *testing.T) {
	t.Parallel()

	repoErr := errors.New("session repository failure")

	sessions := &updatePrayerSessionRepositoryStub{
		findErr: repoErr,
	}

	users := &updateUserRepositoryStub{}

	service := NewService(sessions, users)

	result, err := service.Update(
		context.Background(),
		UpdateCommand{
			ActorID:   identity.UserID("user-1"),
			SessionID: identity.PrayerSessionID("session-1"),
			Title:     "Evening Prayer",
			Date:      time.Now(),
			Time:      "19:30",
			Duration:  45,
		},
	)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, repoErr)

	assert.False(t, users.findCalled)
	assert.False(t, sessions.updateCalled)
}

// ---------------------------------------------------------
// User not found
// ---------------------------------------------------------

func TestService_Update_UserNotFound(t *testing.T) {
	t.Parallel()

	groupID := identity.PrayerGroupID("group-1")

	existing := &domain.PrayerSession{
		ID:            identity.PrayerSessionID("session-1"),
		PrayerGroupID: groupID,
		Title:         "Morning Prayer",
		Date:          time.Now(),
		Time:          "06:00",
		Duration:      30,
	}

	sessions := &updatePrayerSessionRepositoryStub{
		session: existing,
	}

	users := &updateUserRepositoryStub{
		user: nil,
	}

	service := NewService(sessions, users)

	result, err := service.Update(
		context.Background(),
		UpdateCommand{
			ActorID:   identity.UserID("user-1"),
			SessionID: identity.PrayerSessionID("session-1"),
			Title:     "Evening Prayer",
			Date:      time.Now(),
			Time:      "19:30",
			Duration:  45,
		},
	)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, domainuser.ErrUserNotFound)

	assert.True(t, users.findCalled)
	assert.False(t, sessions.updateCalled)
}

// ---------------------------------------------------------
// User repository error
// ---------------------------------------------------------

func TestService_Update_UserRepositoryError(t *testing.T) {
	t.Parallel()

	groupID := identity.PrayerGroupID("group-1")

	existing := &domain.PrayerSession{
		ID:            identity.PrayerSessionID("session-1"),
		PrayerGroupID: groupID,
		Title:         "Morning Prayer",
		Date:          time.Now(),
		Time:          "06:00",
		Duration:      30,
	}

	repoErr := errors.New("user repository failure")

	sessions := &updatePrayerSessionRepositoryStub{
		session: existing,
	}

	users := &updateUserRepositoryStub{
		findErr: repoErr,
	}

	service := NewService(sessions, users)

	result, err := service.Update(
		context.Background(),
		UpdateCommand{
			ActorID:   identity.UserID("user-1"),
			SessionID: identity.PrayerSessionID("session-1"),
			Title:     "Evening Prayer",
			Date:      time.Now(),
			Time:      "19:30",
			Duration:  45,
		},
	)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, repoErr)

	assert.False(t, sessions.updateCalled)
}

// ---------------------------------------------------------
// Forbidden
// ---------------------------------------------------------

func TestService_Update_Forbidden(t *testing.T) {
	t.Parallel()

	groupID := identity.PrayerGroupID("group-1")

	existing := &domain.PrayerSession{
		ID:            identity.PrayerSessionID("session-1"),
		PrayerGroupID: groupID,
		Title:         "Morning Prayer",
		Date:          time.Now(),
		Time:          "06:00",
		Duration:      30,
	}

	user := &domainuser.User{
		ID: identity.UserID("user-1"),
		PrayerGroupIDs: []identity.PrayerGroupID{
			identity.PrayerGroupID("different-group"),
		},
	}

	sessions := &updatePrayerSessionRepositoryStub{
		session: existing,
	}

	users := &updateUserRepositoryStub{
		user: user,
	}

	service := NewService(sessions, users)

	result, err := service.Update(
		context.Background(),
		UpdateCommand{
			ActorID:   identity.UserID("user-1"),
			SessionID: identity.PrayerSessionID("session-1"),
			Title:     "Evening Prayer",
			Date:      time.Now(),
			Time:      "19:30",
			Duration:  45,
		},
	)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, ErrForbidden)

	assert.False(t, sessions.updateCalled)
}

// ---------------------------------------------------------
// Update repository error
// ---------------------------------------------------------

func TestService_Update_UpdateRepositoryError(t *testing.T) {
	t.Parallel()

	groupID := identity.PrayerGroupID("group-1")

	existing := &domain.PrayerSession{
		ID:            identity.PrayerSessionID("session-1"),
		PrayerGroupID: groupID,
		Title:         "Morning Prayer",
		Date:          time.Now(),
		Time:          "06:00",
		Duration:      30,
	}

	user := &domainuser.User{
		ID: identity.UserID("user-1"),
		PrayerGroupIDs: []identity.PrayerGroupID{
			groupID,
		},
	}

	repoErr := errors.New("update failed")

	sessions := &updatePrayerSessionRepositoryStub{
		session:   existing,
		updateErr: repoErr,
	}

	users := &updateUserRepositoryStub{
		user: user,
	}

	service := NewService(sessions, users)

	result, err := service.Update(
		context.Background(),
		UpdateCommand{
			ActorID:   identity.UserID("user-1"),
			SessionID: identity.PrayerSessionID("session-1"),
			Title:     "Evening Prayer",
			Date:      time.Now(),
			Time:      "19:30",
			Duration:  45,
		},
	)

	assert.Nil(t, result)
	assert.ErrorIs(t, err, repoErr)

	assert.True(t, sessions.updateCalled)
}