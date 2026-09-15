package prayersession

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	app "prayer-api/internal/application/prayersession"
	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	domainuser "prayer-api/internal/domain/user"
)

func TestUpdate_Success(t *testing.T) {
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

	var saved *domain.PrayerSession

	sessions := &mockPrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return existing, nil
		},
		updateFn: func(
			ctx context.Context,
			session *domain.PrayerSession,
		) error {
			saved = session
			return nil
		},
	}

	users := &mockUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*domainuser.User, error) {
			return user, nil
		},
	}

	service := app.NewService(sessions, users)

	result, err := service.Update(
		context.Background(),
		app.UpdateCommand{
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
	require.NotNil(t, saved)

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
}