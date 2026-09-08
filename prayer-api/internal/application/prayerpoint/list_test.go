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

func TestService_ListByGroupID(t *testing.T) {
	ctx := context.Background()

	userID := identity.UserID("user-123")
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

	t.Run("lists prayer points when user has access", func(t *testing.T) {
		points := []pointdomain.PrayerPoint{
			{
				ID:      identity.PrayerPointID("point-1"),
				GroupID: groupID,
				Title:   "Pray for healing",
				Content: "Pray for complete healing",
				Status:  pointdomain.StatusActive,
			},
			{
				ID:      identity.PrayerPointID("point-2"),
				GroupID: groupID,
				Title:   "Pray for wisdom",
				Content: "Pray for wisdom and guidance",
				Status:  pointdomain.StatusActive,
			},
		}

		var checkedUserID identity.UserID
		var checkedGroupID identity.PrayerGroupID
		var listedGroupID identity.PrayerGroupID

		repo := &mockPrayerPointRepository{
			listByGroupIDFn: func(
				ctx context.Context,
				requestedGroupID identity.PrayerGroupID,
			) ([]pointdomain.PrayerPoint, error) {
				listedGroupID = requestedGroupID
				return points, nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				requestedUserID identity.UserID,
				requestedGroupID identity.PrayerGroupID,
			) (bool, error) {
				checkedUserID = requestedUserID
				checkedGroupID = requestedGroupID
				return true, nil
			},
		}

		service := newService(repo, access)

		result, err := service.ListByGroupID(
			ctx,
			ListInput{
				UserID:  userID,
				GroupID: groupID,
			},
		)

		require.NoError(t, err)
		require.Len(t, result, 2)

		assert.Equal(t, userID, checkedUserID)
		assert.Equal(t, groupID, checkedGroupID)
		assert.Equal(t, groupID, listedGroupID)

		assert.Equal(t, points[0], result[0])
		assert.Equal(t, points[1], result[1])
	})

	t.Run("returns empty list when group has no prayer points", func(t *testing.T) {
		repo := &mockPrayerPointRepository{
			listByGroupIDFn: func(
				ctx context.Context,
				requestedGroupID identity.PrayerGroupID,
			) ([]pointdomain.PrayerPoint, error) {
				return []pointdomain.PrayerPoint{}, nil
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

		service := newService(repo, access)

		result, err := service.ListByGroupID(
			ctx,
			ListInput{
				UserID:  userID,
				GroupID: groupID,
			},
		)

		require.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("returns nil or empty list when repository returns nil", func(t *testing.T) {
		repo := &mockPrayerPointRepository{
			listByGroupIDFn: func(
				ctx context.Context,
				requestedGroupID identity.PrayerGroupID,
			) ([]pointdomain.PrayerPoint, error) {
				return nil, nil
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

		service := newService(repo, access)

		result, err := service.ListByGroupID(
			ctx,
			ListInput{
				UserID:  userID,
				GroupID: groupID,
			},
		)

		require.NoError(t, err)
		assert.Empty(t, result)
	})

	t.Run("returns access denied when user has no access", func(t *testing.T) {
		var listCalled bool

		repo := &mockPrayerPointRepository{
			listByGroupIDFn: func(
				ctx context.Context,
				requestedGroupID identity.PrayerGroupID,
			) ([]pointdomain.PrayerPoint, error) {
				listCalled = true
				return nil, nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				requestedUserID identity.UserID,
				requestedGroupID identity.PrayerGroupID,
			) (bool, error) {
				assert.Equal(t, userID, requestedUserID)
				assert.Equal(t, groupID, requestedGroupID)
				return false, nil
			},
		}

		service := newService(repo, access)

		result, err := service.ListByGroupID(
			ctx,
			ListInput{
				UserID:  userID,
				GroupID: groupID,
			},
		)

		require.Error(t, err)
		assert.ErrorIs(t, err, ErrAccessDenied)
		assert.Nil(t, result)
		assert.False(t, listCalled)
	})

	t.Run("returns error when access check fails", func(t *testing.T) {
		expectedErr := errors.New("access check failed")
		var listCalled bool

		repo := &mockPrayerPointRepository{
			listByGroupIDFn: func(
				ctx context.Context,
				requestedGroupID identity.PrayerGroupID,
			) ([]pointdomain.PrayerPoint, error) {
				listCalled = true
				return nil, nil
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				requestedUserID identity.UserID,
				requestedGroupID identity.PrayerGroupID,
			) (bool, error) {
				return false, expectedErr
			},
		}

		service := newService(repo, access)

		result, err := service.ListByGroupID(
			ctx,
			ListInput{
				UserID:  userID,
				GroupID: groupID,
			},
		)

		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, result)
		assert.False(t, listCalled)
	})

	t.Run("returns repository error when listing fails", func(t *testing.T) {
		expectedErr := errors.New("list prayer points failed")

		repo := &mockPrayerPointRepository{
			listByGroupIDFn: func(
				ctx context.Context,
				requestedGroupID identity.PrayerGroupID,
			) ([]pointdomain.PrayerPoint, error) {
				assert.Equal(t, groupID, requestedGroupID)
				return nil, expectedErr
			},
		}

		access := &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				requestedUserID identity.UserID,
				requestedGroupID identity.PrayerGroupID,
			) (bool, error) {
				return true, nil
			},
		}

		service := newService(repo, access)

		result, err := service.ListByGroupID(
			ctx,
			ListInput{
				UserID:  userID,
				GroupID: groupID,
			},
		)

		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, result)
	})

	t.Run("passes the requested group ID to the repository", func(t *testing.T) {
		var receivedGroupID identity.PrayerGroupID

		repo := &mockPrayerPointRepository{
			listByGroupIDFn: func(
				ctx context.Context,
				requestedGroupID identity.PrayerGroupID,
			) ([]pointdomain.PrayerPoint, error) {
				receivedGroupID = requestedGroupID
				return []pointdomain.PrayerPoint{}, nil
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

		service := newService(repo, access)

		_, err := service.ListByGroupID(
			ctx,
			ListInput{
				UserID:  userID,
				GroupID: groupID,
			},
		)

		require.NoError(t, err)
		assert.Equal(t, groupID, receivedGroupID)
	})
}