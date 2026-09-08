
package prayerpoint

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"prayer-api/internal/domain/identity"
	pointdomain "prayer-api/internal/domain/prayerpoint"
)

func TestService_Delete(t *testing.T) {
	ctx := context.Background()

	userID := identity.UserID("user-123")
	pointID := identity.PrayerPointID("point-123")
	groupID := identity.PrayerGroupID("group-123")

	newService := func(
		points *mockPrayerPointRepository,
		access *mockUserAccessReader,
	) *Service {
		return NewService(
			points,
			&mockPrayerGroupReader{},
			access,
		)
	}

	t.Run("deletes prayer point when user has access", func(t *testing.T) {
		point, err := pointdomain.New(
			pointID,
			groupID,
			"Pray for healing",
			"Pray for complete healing",
		)
		require.NoError(t, err)

		var deletedPointID identity.PrayerPointID
		var accessCheckedUserID identity.UserID
		var accessCheckedGroupID identity.PrayerGroupID

		points := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				assert.Equal(t, pointID, id)
				return point, nil
			},
			deleteFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) error {
				deletedPointID = id
				return nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				checkedUserID identity.UserID,
				checkedGroupID identity.PrayerGroupID,
			) (bool, error) {
				accessCheckedUserID = checkedUserID
				accessCheckedGroupID = checkedGroupID
				return true, nil
			},
		}

		service := newService(points, access)

		err = service.Delete(ctx, userID, pointID)

		require.NoError(t, err)
		assert.Equal(t, pointID, deletedPointID)
		assert.Equal(t, userID, accessCheckedUserID)
		assert.Equal(t, groupID, accessCheckedGroupID)
	})

	t.Run("returns error when finding prayer point fails", func(t *testing.T) {
		expectedErr := errors.New("find prayer point failed")

		var accessCalled bool
		var deleteCalled bool

		points := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return nil, expectedErr
			},
			deleteFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) error {
				deleteCalled = true
				return nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				userID identity.UserID,
				groupID identity.PrayerGroupID,
			) (bool, error) {
				accessCalled = true
				return true, nil
			},
		}

		service := newService(points, access)

		err := service.Delete(ctx, userID, pointID)

		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
		assert.False(t, accessCalled)
		assert.False(t, deleteCalled)
	})

	t.Run("returns prayer point not found when point does not exist", func(t *testing.T) {
		var accessCalled bool
		var deleteCalled bool

		points := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return nil, nil
			},
			deleteFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) error {
				deleteCalled = true
				return nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				userID identity.UserID,
				groupID identity.PrayerGroupID,
			) (bool, error) {
				accessCalled = true
				return true, nil
			},
		}

		service := newService(points, access)

		err := service.Delete(ctx, userID, pointID)

		require.Error(t, err)
		assert.ErrorIs(t, err, ErrPrayerPointNotFound)
		assert.False(t, accessCalled)
		assert.False(t, deleteCalled)
	})

	t.Run("returns access denied when user has no group access", func(t *testing.T) {
		point, err := pointdomain.New(
			pointID,
			groupID,
			"Pray for healing",
			"Pray for complete healing",
		)
		require.NoError(t, err)

		var deleteCalled bool
		var accessCheckedUserID identity.UserID
		var accessCheckedGroupID identity.PrayerGroupID

		points := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return point, nil
			},
			deleteFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) error {
				deleteCalled = true
				return nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				checkedUserID identity.UserID,
				checkedGroupID identity.PrayerGroupID,
			) (bool, error) {
				accessCheckedUserID = checkedUserID
				accessCheckedGroupID = checkedGroupID
				return false, nil
			},
		}

		service := newService(points, access)

		err = service.Delete(ctx, userID, pointID)

		require.Error(t, err)
		assert.ErrorIs(t, err, ErrAccessDenied)
		assert.Equal(t, userID, accessCheckedUserID)
		assert.Equal(t, groupID, accessCheckedGroupID)
		assert.False(t, deleteCalled)
	})

	t.Run("returns error when access check fails", func(t *testing.T) {
		expectedErr := errors.New("access check failed")

		point, err := pointdomain.New(
			pointID,
			groupID,
			"Pray for healing",
			"Pray for complete healing",
		)
		require.NoError(t, err)

		var deleteCalled bool

		points := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return point, nil
			},
			deleteFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) error {
				deleteCalled = true
				return nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				userID identity.UserID,
				groupID identity.PrayerGroupID,
			) (bool, error) {
				return false, expectedErr
			},
		}

		service := newService(points, access)

		err = service.Delete(ctx, userID, pointID)

		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
		assert.False(t, deleteCalled)
	})

	t.Run("returns error when repository delete fails", func(t *testing.T) {
		expectedErr := errors.New("delete failed")

		point, err := pointdomain.New(
			pointID,
			groupID,
			"Pray for healing",
			"Pray for complete healing",
		)
		require.NoError(t, err)

		points := &mockPrayerPointRepository{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) (*pointdomain.PrayerPoint, error) {
				return point, nil
			},
			deleteFn: func(
				ctx context.Context,
				id identity.PrayerPointID,
			) error {
				assert.Equal(t, pointID, id)
				return expectedErr
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

		service := newService(points, access)

		err = service.Delete(ctx, userID, pointID)

		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
	})
}
