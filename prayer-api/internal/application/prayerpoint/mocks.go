package prayerpoint

import (
	"context"

	"prayer-api/internal/domain/identity"
	groupdomain "prayer-api/internal/domain/prayergroup"
	pointdomain "prayer-api/internal/domain/prayerpoint"
)

type mockPrayerPointRepository struct {
	findByIDFn func(
		ctx context.Context,
		id identity.PrayerPointID,
	) (*pointdomain.PrayerPoint, error)

	listByGroupIDFn func(
		ctx context.Context,
		groupID identity.PrayerGroupID,
	) ([]pointdomain.PrayerPoint, error)

	saveFn func(
		ctx context.Context,
		point *pointdomain.PrayerPoint,
	) error

	updateFn func(
		ctx context.Context,
		point *pointdomain.PrayerPoint,
	) error

	deleteFn func(
		ctx context.Context,
		id identity.PrayerPointID,
	) error
}

func (m *mockPrayerPointRepository) FindByID(
	ctx context.Context,
	id identity.PrayerPointID,
) (*pointdomain.PrayerPoint, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}

func (m *mockPrayerPointRepository) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]pointdomain.PrayerPoint, error) {
	if m.listByGroupIDFn != nil {
		return m.listByGroupIDFn(ctx, groupID)
	}

	return nil, nil
}

func (m *mockPrayerPointRepository) Save(
	ctx context.Context,
	point *pointdomain.PrayerPoint,
) error {
	if m.saveFn != nil {
		return m.saveFn(ctx, point)
	}

	return nil
}

func (m *mockPrayerPointRepository) Update(
	ctx context.Context,
	point *pointdomain.PrayerPoint,
) error {
	if m.updateFn != nil {
		return m.updateFn(ctx, point)
	}

	return nil
}

func (m *mockPrayerPointRepository) Delete(
	ctx context.Context,
	id identity.PrayerPointID,
) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}

	return nil
}

type mockPrayerGroupReader struct {
	findByIDFn func(
		ctx context.Context,
		id identity.PrayerGroupID,
	) (*groupdomain.PrayerGroup, error)
}

func (m *mockPrayerGroupReader) FindByID(
	ctx context.Context,
	id identity.PrayerGroupID,
) (*groupdomain.PrayerGroup, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}

type mockUserAccessReader struct {
	hasPrayerGroupFn func(
		ctx context.Context,
		userID identity.UserID,
		groupID identity.PrayerGroupID,
	) (bool, error)
}

func (m *mockUserAccessReader) HasPrayerGroup(
	ctx context.Context,
	userID identity.UserID,
	groupID identity.PrayerGroupID,
) (bool, error) {
	if m.hasPrayerGroupFn != nil {
		return m.hasPrayerGroupFn(ctx, userID, groupID)
	}

	return false, nil
}
