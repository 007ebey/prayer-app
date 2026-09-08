package prayerpoint

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	groupdomain "prayer-api/internal/domain/prayergroup"
	"prayer-api/internal/domain/identity"
	pointdomain "prayer-api/internal/domain/prayerpoint"
)

func TestService_Block(t *testing.T) {
	ctx := context.Background()

	pointID := identity.PrayerPointID("point-123")
	groupID := identity.PrayerGroupID("group-123")
	userID := identity.UserID("user-123")

	newPoint := func(t *testing.T) *pointdomain.PrayerPoint {
		t.Helper()

		point, err := pointdomain.New(
			pointID,
			groupID,
			"Pray for healing",
			"Pray for complete healing",
		)
		require.NoError(t, err)

		return point
	}

	newGroup := func() *groupdomain.PrayerGroup {
		return &groupdomain.PrayerGroup{
			ID: groupID,
		}
	}

	t.Run("blocks an active prayer point", func(t *testing.T) {
		point := newPoint(t)

		var updatedPoint *pointdomain.PrayerPoint

		repo := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				assert.Equal(t, pointID, id)
				return point, nil
			},
			updateFn: func(
				ctx context.Context,
				p *pointdomain.PrayerPoint,
			) error {
				updatedPoint = p
				return nil
			},
		}

		groups := &mockPrayerGroupReader{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerGroupID,
			) (*groupdomain.PrayerGroup, error) {
				assert.Equal(t, groupID, id)
				return newGroup(), nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				checkedUserID identity.UserID,
				checkedGroupID identity.PrayerGroupID,
			) (bool, error) {
				assert.Equal(t, userID, checkedUserID)
				assert.Equal(t, groupID, checkedGroupID)
				return true, nil
			},
		}

		service := NewService(repo, groups, access)

		err := service.Block(ctx, userID, pointID)

		require.NoError(t, err)
		require.NotNil(t, updatedPoint)
		assert.Equal(t, pointdomain.StatusBlocked, updatedPoint.Status)
		assert.False(t, updatedPoint.IsActive())
	})

	t.Run("returns repository error when finding prayer point fails", func(t *testing.T) {
		expectedErr := errors.New("database unavailable")

		repo := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return nil, expectedErr
			},
		}

		service := NewService(
			repo,
			&mockPrayerGroupReader{},
			&mockUserAccessReader{},
		)

		err := service.Block(ctx, userID, pointID)

		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("returns not found when prayer point does not exist", func(t *testing.T) {
		repo := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return nil, nil
			},
		}

		service := NewService(
			repo,
			&mockPrayerGroupReader{},
			&mockUserAccessReader{},
		)

		err := service.Block(ctx, userID, pointID)

		require.Error(t, err)
		assert.ErrorIs(t, err, pointdomain.ErrNotFound)
	})

	t.Run("returns repository error when update fails", func(t *testing.T) {
		expectedErr := errors.New("update failed")
		point := newPoint(t)

		repo := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return point, nil
			},
			updateFn: func(
				ctx context.Context,
				p *pointdomain.PrayerPoint,
			) error {
				return expectedErr
			},
		}

		groups := &mockPrayerGroupReader{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerGroupID,
			) (*groupdomain.PrayerGroup, error) {
				return newGroup(), nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				userID identity.UserID,
				groupID identity.PrayerGroupID,
			) (bool, error) {
				return true, nil
			},
		}

		service := NewService(repo, groups, access)

		err := service.Block(ctx, userID, pointID)

		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
	})

	t.Run("blocking an already blocked prayer point remains blocked", func(t *testing.T) {
		point := newPoint(t)
		point.Block()

		var updateCount int

		repo := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return point, nil
			},
			updateFn: func(
				ctx context.Context,
				p *pointdomain.PrayerPoint,
			) error {
				updateCount++
				assert.Equal(t, pointdomain.StatusBlocked, p.Status)
				return nil
			},
		}

		groups := &mockPrayerGroupReader{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerGroupID,
			) (*groupdomain.PrayerGroup, error) {
				return newGroup(), nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				userID identity.UserID,
				groupID identity.PrayerGroupID,
			) (bool, error) {
				return true, nil
			},
		}

		service := NewService(repo, groups, access)

		err := service.Block(ctx, userID, pointID)

		require.NoError(t, err)
		assert.Equal(t, 1, updateCount)
		assert.Equal(t, pointdomain.StatusBlocked, point.Status)
	})
}
