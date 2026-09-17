package prayersession

import (
	"testing"
	"time"

	"prayer-api/internal/domain/identity"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewPrayerSession(t *testing.T) {
	t.Parallel()

	validID := identity.PrayerSessionID("session-1")
	validGroupID := identity.PrayerGroupID("group-1")
	validTitle := "Morning Prayer"
	validDate := time.Date(2026, 9, 18, 0, 0, 0, 0, time.UTC)
	validTime := "07:00"
	validDuration := 60
	validPrayerPointIDs := []identity.PrayerPointID{
		"point-1",
		"point-2",
	}

	t.Run("creates a prayer session successfully", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			validTitle,
			validDate,
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		require.NoError(t, err)
		require.NotNil(t, session)

		assert.Equal(t, validID, session.ID)
		assert.Equal(t, validGroupID, session.PrayerGroupID)
		assert.Equal(t, validTitle, session.Title)
		assert.Equal(t, validDate, session.Date)
		assert.Equal(t, validTime, session.Time)
		assert.Equal(t, validDuration, session.Duration)
		assert.Equal(t, validPrayerPointIDs, session.PrayerPointIDs)
	})

	t.Run("requires an ID", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			"",
			validGroupID,
			validTitle,
			validDate,
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidID)
	})

	t.Run("rejects whitespace-only ID", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			"   ",
			validGroupID,
			validTitle,
			validDate,
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidID)
	})

	t.Run("requires a prayer group ID", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			"",
			validTitle,
			validDate,
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidPrayerGroupID)
	})

	t.Run("rejects whitespace-only prayer group ID", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			"   ",
			validTitle,
			validDate,
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidPrayerGroupID)
	})

	t.Run("requires a title", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			"",
			validDate,
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidTitle)
	})

	t.Run("rejects whitespace-only title", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			"   ",
			validDate,
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidTitle)
	})

	t.Run("trims title", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			"  Morning Prayer  ",
			validDate,
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		require.NoError(t, err)
		assert.Equal(t, "Morning Prayer", session.Title)
	})

	t.Run("requires a date", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			validTitle,
			time.Time{},
			validTime,
			validDuration,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidDate)
	})

	t.Run("requires a positive duration", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			validTitle,
			validDate,
			validTime,
			0,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidDuration)
	})

	t.Run("rejects negative duration", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			validTitle,
			validDate,
			validTime,
			-10,
			validPrayerPointIDs,
		)

		assert.Nil(t, session)
		assert.ErrorIs(t, err, ErrInvalidDuration)
	})

	t.Run("trims session time", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			validTitle,
			validDate,
			"  07:00  ",
			validDuration,
			validPrayerPointIDs,
		)

		require.NoError(t, err)
		assert.Equal(t, "07:00", session.Time)
	})

	t.Run("copies prayer point IDs", func(t *testing.T) {
		t.Parallel()

		input := []identity.PrayerPointID{
			"point-1",
			"point-2",
		}

		session, err := New(
			validID,
			validGroupID,
			validTitle,
			validDate,
			validTime,
			validDuration,
			input,
		)

		require.NoError(t, err)

		// Mutate the original input.
		input[0] = "changed"

		// Domain object should remain unchanged.
		assert.Equal(
			t,
			identity.PrayerPointID("point-1"),
			session.PrayerPointIDs[0],
		)
	})

	t.Run("allows nil prayer point IDs", func(t *testing.T) {
		t.Parallel()

		session, err := New(
			validID,
			validGroupID,
			validTitle,
			validDate,
			validTime,
			validDuration,
			nil,
		)

		require.NoError(t, err)
		require.NotNil(t, session)
		assert.Empty(t, session.PrayerPointIDs)
	})
}