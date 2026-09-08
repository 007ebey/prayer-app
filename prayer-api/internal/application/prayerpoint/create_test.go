package prayerpoint

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"prayer-api/internal/domain/identity"
	groupdomain "prayer-api/internal/domain/prayergroup"
	domain "prayer-api/internal/domain/prayerpoint"
)

func TestService_Create(t *testing.T) {
	ctx := context.Background()

	groupID := identity.PrayerGroupID("group-123")

	newService := func(
		repo *mockPrayerPointRepository,
		groups *mockPrayerGroupReader,
		access *mockUserAccessReader,
	) *Service {
		return NewService(
			repo,
			groups,
			access,
		)
	}

	newGroupReader := func() *mockPrayerGroupReader {
		return &mockPrayerGroupReader{
			findByIDFn: func(
				ctx context.Context,
				id identity.PrayerGroupID,
			) (*groupdomain.PrayerGroup, error) {
				return &groupdomain.PrayerGroup{
					ID: id,
				}, nil
			},
		}
	}

	newAccessReader := func() *mockUserAccessReader {
		return &mockUserAccessReader{
			hasPrayerGroupFn: func(
				ctx context.Context,
				userID identity.UserID,
				groupID identity.PrayerGroupID,
			) (bool, error) {
				return true, nil
			},
		}
	}

	t.Run("creates a prayer point successfully", func(t *testing.T) {
		var savedPoint *domain.PrayerPoint

		repo := &mockPrayerPointRepository{
			saveFn: func(
				ctx context.Context,
				point *domain.PrayerPoint,
			) error {
				savedPoint = point
				return nil
			},
		}

		service := newService(
			repo,
			newGroupReader(),
			newAccessReader(),
		)

		result, err := service.Create(
			ctx,
			CreateInput{
				GroupID: groupID,
				Title:   "Pray for healing",
				Content: "Pray for complete healing and strength",
			},
		)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, savedPoint)

		assert.NotEmpty(t, result.ID)
		assert.Equal(t, groupID, result.GroupID)
		assert.Equal(t, "Pray for healing", result.Title)
		assert.Equal(
			t,
			"Pray for complete healing and strength",
			result.Content,
		)
		assert.Equal(t, domain.StatusActive, result.Status)

		assert.Equal(t, result.ID, savedPoint.ID)
		assert.Equal(t, result.GroupID, savedPoint.GroupID)
		assert.Equal(t, result.Title, savedPoint.Title)
		assert.Equal(t, result.Content, savedPoint.Content)
	})

	t.Run("trims title and content before saving", func(t *testing.T) {
		var savedPoint *domain.PrayerPoint

		repo := &mockPrayerPointRepository{
			saveFn: func(
				ctx context.Context,
				point *domain.PrayerPoint,
			) error {
				savedPoint = point
				return nil
			},
		}

		service := newService(
			repo,
			newGroupReader(),
			newAccessReader(),
		)

		result, err := service.Create(
			ctx,
			CreateInput{
				GroupID: groupID,
				Title:   "  Pray for healing  ",
				Content: "  Pray for complete healing  ",
			},
		)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, savedPoint)

		assert.Equal(t, "Pray for healing", result.Title)
		assert.Equal(t, "Pray for complete healing", result.Content)
	})

	t.Run("returns validation error when group ID is missing", func(t *testing.T) {
		var saveCalled bool

		repo := &mockPrayerPointRepository{
			saveFn: func(
				ctx context.Context,
				point *domain.PrayerPoint,
			) error {
				saveCalled = true
				return nil
			},
		}

		service := newService(
			repo,
			newGroupReader(),
			newAccessReader(),
		)

		result, err := service.Create(
			ctx,
			CreateInput{
				Title:   "Pray for healing",
				Content: "Pray for complete healing",
			},
		)

		require.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrGroupIDRequired)
		assert.Nil(t, result)
		assert.False(t, saveCalled)
	})

	t.Run("returns validation error when title is empty", func(t *testing.T) {
		var saveCalled bool

		repo := &mockPrayerPointRepository{
			saveFn: func(
				ctx context.Context,
				point *domain.PrayerPoint,
			) error {
				saveCalled = true
				return nil
			},
		}

		service := newService(
			repo,
			newGroupReader(),
			newAccessReader(),
		)

		result, err := service.Create(
			ctx,
			CreateInput{
				GroupID: groupID,
				Title:   "   ",
				Content: "Pray for complete healing",
			},
		)

		require.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrTitleRequired)
		assert.Nil(t, result)
		assert.False(t, saveCalled)
	})

	t.Run("returns validation error when content is empty", func(t *testing.T) {
		var saveCalled bool

		repo := &mockPrayerPointRepository{
			saveFn: func(
				ctx context.Context,
				point *domain.PrayerPoint,
			) error {
				saveCalled = true
				return nil
			},
		}

		service := newService(
			repo,
			newGroupReader(),
			newAccessReader(),
		)

		result, err := service.Create(
			ctx,
			CreateInput{
				GroupID: groupID,
				Title:   "Pray for healing",
				Content: "   ",
			},
		)

		require.Error(t, err)
		assert.ErrorIs(t, err, domain.ErrContentRequired)
		assert.Nil(t, result)
		assert.False(t, saveCalled)
	})

	t.Run("returns repository error when save fails", func(t *testing.T) {
		expectedErr := errors.New("save failed")

		repo := &mockPrayerPointRepository{
			saveFn: func(
				ctx context.Context,
				point *domain.PrayerPoint,
			) error {
				return expectedErr
			},
		}

		service := newService(
			repo,
			newGroupReader(),
			newAccessReader(),
		)

		result, err := service.Create(
			ctx,
			CreateInput{
				GroupID: groupID,
				Title:   "Pray for healing",
				Content: "Pray for complete healing",
			},
		)

		require.Error(t, err)
		assert.ErrorIs(t, err, expectedErr)
		assert.Nil(t, result)
	})

	t.Run("creates every prayer point with active status", func(t *testing.T) {
		var savedPoint *domain.PrayerPoint

		repo := &mockPrayerPointRepository{
			saveFn: func(
				ctx context.Context,
				point *domain.PrayerPoint,
			) error {
				savedPoint = point
				return nil
			},
		}

		service := newService(
			repo,
			newGroupReader(),
			newAccessReader(),
		)

		result, err := service.Create(
			ctx,
			CreateInput{
				GroupID: groupID,
				Title:   "Pray for family",
				Content: "Pray for peace in the family",
			},
		)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, savedPoint)

		assert.Equal(t, domain.StatusActive, result.Status)
		assert.True(t, result.IsActive())
		assert.Equal(t, domain.StatusActive, savedPoint.Status)
	})

	t.Run("sets created and updated timestamps", func(t *testing.T) {
		var savedPoint *domain.PrayerPoint

		repo := &mockPrayerPointRepository{
			saveFn: func(
				ctx context.Context,
				point *domain.PrayerPoint,
			) error {
				savedPoint = point
				return nil
			},
		}

		service := newService(
			repo,
			newGroupReader(),
			newAccessReader(),
		)

		result, err := service.Create(
			ctx,
			CreateInput{
				GroupID: groupID,
				Title:   "Pray for wisdom",
				Content: "Pray for wisdom and guidance",
			},
		)

		require.NoError(t, err)
		require.NotNil(t, result)
		require.NotNil(t, savedPoint)

		assert.NotEmpty(t, result.ID)
		assert.False(t, result.CreatedAt.IsZero())
		assert.False(t, result.UpdatedAt.IsZero())
		assert.Equal(t, result.CreatedAt, savedPoint.CreatedAt)
		assert.Equal(t, result.UpdatedAt, savedPoint.UpdatedAt)
	})
}
