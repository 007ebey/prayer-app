package prayerpoint

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"prayer-api/internal/domain/identity"
	pointdomain "prayer-api/internal/domain/prayerpoint"
)

func TestService_Update_Success(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := identity.UserID("user-123")
	pointID := identity.PrayerPointID("point-123")
	groupID := identity.PrayerGroupID("group-123")

	point, err := pointdomain.New(
		pointID,
		groupID,
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	originalUpdatedAt := point.UpdatedAt

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
			updatedPoint *pointdomain.PrayerPoint,
		) error {
			assert.Equal(t, pointID, updatedPoint.ID)
			assert.Equal(t, groupID, updatedPoint.GroupID)
			assert.Equal(t, "New title", updatedPoint.Title)
			assert.Equal(t, "New content", updatedPoint.Content)

			return nil
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

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		userID,
		pointID,
		" New title ",
		" New content ",
	)

	require.NoError(t, err)
	require.NotNil(t, updatedPoint)

	assert.Equal(t, pointID, updatedPoint.ID)
	assert.Equal(t, groupID, updatedPoint.GroupID)
	assert.Equal(t, "New title", updatedPoint.Title)
	assert.Equal(t, "New content", updatedPoint.Content)
	assert.Equal(t, pointdomain.StatusActive, updatedPoint.Status)
	assert.True(t, updatedPoint.UpdatedAt.After(originalUpdatedAt))
}

func TestService_Update_FindByIDError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	expectedErr := errors.New("find prayer point failed")

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return nil, expectedErr
		},
	}

	access := &mockUserAccessReader{}

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		identity.PrayerPointID("point-123"),
		"New title",
		"New content",
	)

	assert.Nil(t, updatedPoint)
	assert.ErrorIs(t, err, expectedErr)
}

func TestService_UpdateNotFound(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return nil, nil
		},
	}

	access := &mockUserAccessReader{}

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		identity.PrayerPointID("point-123"),
		"New title",
		"New content",
	)

	assert.Nil(t, updatedPoint)
	assert.ErrorIs(t, err, ErrPrayerPointNotFound)
}

func TestService_UpdateAccessDenied(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	userID := identity.UserID("user-123")
	pointID := identity.PrayerPointID("point-123")
	groupID := identity.PrayerGroupID("group-123")

	point, err := pointdomain.New(
		pointID,
		groupID,
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	var updateCalled bool

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return point, nil
		},
		updateFn: func(
			ctx context.Context,
			point *pointdomain.PrayerPoint,
		) error {
			updateCalled = true
			return nil
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

			return false, nil
		},
	}

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		userID,
		pointID,
		"New title",
		"New content",
	)

	assert.Nil(t, updatedPoint)
	assert.ErrorIs(t, err, ErrAccessDenied)
	assert.False(t, updateCalled)
}

func TestService_UpdateAccessCheckError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	expectedErr := errors.New("access check failed")

	point, err := pointdomain.New(
		identity.PrayerPointID("point-123"),
		identity.PrayerGroupID("group-123"),
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return point, nil
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

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		identity.PrayerPointID("point-123"),
		"New title",
		"New content",
	)

	assert.Nil(t, updatedPoint)
	assert.ErrorIs(t, err, expectedErr)
}

func TestService_UpdateTitleValidationError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	point, err := pointdomain.New(
		identity.PrayerPointID("point-123"),
		identity.PrayerGroupID("group-123"),
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	var updateCalled bool

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return point, nil
		},
		updateFn: func(
			ctx context.Context,
			point *pointdomain.PrayerPoint,
		) error {
			updateCalled = true
			return nil
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

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		identity.PrayerPointID("point-123"),
		"   ",
		"New content",
	)

	assert.Nil(t, updatedPoint)
	assert.ErrorIs(t, err, pointdomain.ErrTitleRequired)
	assert.False(t, updateCalled)
}

func TestService_UpdateContentValidationError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	point, err := pointdomain.New(
		identity.PrayerPointID("point-123"),
		identity.PrayerGroupID("group-123"),
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	var updateCalled bool

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return point, nil
		},
		updateFn: func(
			ctx context.Context,
			point *pointdomain.PrayerPoint,
		) error {
			updateCalled = true
			return nil
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

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		identity.PrayerPointID("point-123"),
		"New title",
		"   ",
	)

	assert.Nil(t, updatedPoint)
	assert.ErrorIs(t, err, pointdomain.ErrContentRequired)
	assert.False(t, updateCalled)
}

func TestService_UpdateRepositoryError(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	expectedErr := errors.New("update prayer point failed")

	point, err := pointdomain.New(
		identity.PrayerPointID("point-123"),
		identity.PrayerGroupID("group-123"),
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return point, nil
		},
		updateFn: func(
			ctx context.Context,
			point *pointdomain.PrayerPoint,
		) error {
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

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		identity.PrayerPointID("point-123"),
		"New title",
		"New content",
	)

	assert.Nil(t, updatedPoint)
	assert.ErrorIs(t, err, expectedErr)
}

func TestService_UpdateDoesNotChangeGroupID(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	pointID := identity.PrayerPointID("point-123")
	groupID := identity.PrayerGroupID("group-123")

	point, err := pointdomain.New(
		pointID,
		groupID,
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return point, nil
		},
		updateFn: func(
			ctx context.Context,
			updatedPoint *pointdomain.PrayerPoint,
		) error {
			assert.Equal(t, groupID, updatedPoint.GroupID)
			return nil
		},
	}

	access := &mockUserAccessReader{
		hasPrayerGroupFn: func(
			ctx context.Context,
			userID identity.UserID,
			checkedGroupID identity.PrayerGroupID,
		) (bool, error) {
			assert.Equal(t, groupID, checkedGroupID)
			return true, nil
		},
	}

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		pointID,
		"New title",
		"New content",
	)

	require.NoError(t, err)
	require.NotNil(t, updatedPoint)

	assert.Equal(t, groupID, updatedPoint.GroupID)
	assert.NotEqual(t, identity.PrayerGroupID("another-group"), updatedPoint.GroupID)
}

func TestService_UpdatePreservesCreatedAt(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	point, err := pointdomain.New(
		identity.PrayerPointID("point-123"),
		identity.PrayerGroupID("group-123"),
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	createdAt := point.CreatedAt

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return point, nil
		},
		updateFn: func(
			ctx context.Context,
			point *pointdomain.PrayerPoint,
		) error {
			return nil
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

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		identity.PrayerPointID("point-123"),
		"New title",
		"New content",
	)

	require.NoError(t, err)
	require.NotNil(t, updatedPoint)

	assert.Equal(t, createdAt, updatedPoint.CreatedAt)
	assert.True(t, updatedPoint.UpdatedAt.After(createdAt))
}

func TestService_UpdateAlreadyBlockedPoint(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	point, err := pointdomain.New(
		identity.PrayerPointID("point-123"),
		identity.PrayerGroupID("group-123"),
		"Old title",
		"Old content",
	)
	require.NoError(t, err)

	point.Block()

	repo := &mockPrayerPointRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerPointID,
		) (*pointdomain.PrayerPoint, error) {
			return point, nil
		},
		updateFn: func(
			ctx context.Context,
			point *pointdomain.PrayerPoint,
		) error {
			assert.Equal(t, pointdomain.StatusBlocked, point.Status)
			return nil
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

	service := NewService(
		repo,
		&mockPrayerGroupReader{},
		access,
	)

	updatedPoint, err := service.Update(
		ctx,
		identity.UserID("user-123"),
		identity.PrayerPointID("point-123"),
		"New title",
		"New content",
	)

	require.NoError(t, err)
	require.NotNil(t, updatedPoint)

	assert.Equal(t, pointdomain.StatusBlocked, updatedPoint.Status)
}


