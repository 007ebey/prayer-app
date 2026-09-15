package prayersession

import (
	"context"
	"errors"
	"testing"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	"prayer-api/internal/domain/user"
)

type listPrayerSessionRepository struct {
	listByGroupIDFn func(
		ctx context.Context,
		groupID identity.PrayerGroupID,
	) ([]domain.PrayerSession, error)
}

func (m *listPrayerSessionRepository) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	return nil, nil
}

func (m *listPrayerSessionRepository) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {
	if m.listByGroupIDFn != nil {
		return m.listByGroupIDFn(ctx, groupID)
	}

	return nil, nil
}

func (m *listPrayerSessionRepository) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (m *listPrayerSessionRepository) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (m *listPrayerSessionRepository) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	return nil
}

type listUserRepository struct {
	findByIDFn func(
		ctx context.Context,
		id identity.UserID,
	) (*user.User, error)
}

func (m *listUserRepository) FindByID(
	ctx context.Context,
	id identity.UserID,
) (*user.User, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}

func testListSessions() []domain.PrayerSession {
	return []domain.PrayerSession{
		{
			ID:             identity.PrayerSessionID("session-1"),
			PrayerGroupID:  identity.PrayerGroupID("group-1"),
			Title:          "Morning Prayer",
			Time:           "06:00",
			Duration:       30,
			PrayerPointIDs: []identity.PrayerPointID{},
		},
		{
			ID:             identity.PrayerSessionID("session-2"),
			PrayerGroupID:  identity.PrayerGroupID("group-1"),
			Title:          "Evening Prayer",
			Time:           "19:00",
			Duration:       45,
			PrayerPointIDs: []identity.PrayerPointID{},
		},
	}
}

func TestListByGroupIDSuccess(t *testing.T) {
	expectedSessions := testListSessions()

	sessions := &listPrayerSessionRepository{
		listByGroupIDFn: func(
			ctx context.Context,
			groupID identity.PrayerGroupID,
		) ([]domain.PrayerSession, error) {

			if groupID != identity.PrayerGroupID("group-1") {
				t.Fatalf(
					"expected group ID %q, got %q",
					"group-1",
					groupID,
				)
			}

			return expectedSessions, nil
		},
	}

	users := &listUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return testUser(
				identity.UserID("user-1"),
				identity.PrayerGroupID("group-1"),
			), nil
		},
	}

	service := NewService(sessions, users)

	result, err := service.ListByGroupID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerGroupID("group-1"),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result) != 2 {
		t.Fatalf(
			"expected 2 sessions, got %d",
			len(result),
		)
	}

	if result[0].ID != identity.PrayerSessionID("session-1") {
		t.Errorf(
			"expected first session %q, got %q",
			"session-1",
			result[0].ID,
		)
	}

	if result[1].ID != identity.PrayerSessionID("session-2") {
		t.Errorf(
			"expected second session %q, got %q",
			"session-2",
			result[1].ID,
		)
	}
}

func TestListByGroupIDEmpty(t *testing.T) {
	sessions := &listPrayerSessionRepository{
		listByGroupIDFn: func(
			ctx context.Context,
			groupID identity.PrayerGroupID,
		) ([]domain.PrayerSession, error) {
			return []domain.PrayerSession{}, nil
		},
	}

	users := &listUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return testUser(
				identity.UserID("user-1"),
				identity.PrayerGroupID("group-1"),
			), nil
		},
	}

	service := NewService(sessions, users)

	result, err := service.ListByGroupID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerGroupID("group-1"),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected empty slice, got nil")
	}

	if len(result) != 0 {
		t.Fatalf(
			"expected 0 sessions, got %d",
			len(result),
		)
	}
}

func TestListByGroupIDAccessDenied(t *testing.T) {
	repositoryCalled := false

	sessions := &listPrayerSessionRepository{
		listByGroupIDFn: func(
			ctx context.Context,
			groupID identity.PrayerGroupID,
		) ([]domain.PrayerSession, error) {
			repositoryCalled = true
			return nil, nil
		},
	}

	users := &listUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return testUser(
				identity.UserID("user-1"),
				identity.PrayerGroupID("different-group"),
			), nil
		},
	}

	service := NewService(sessions, users)

	result, err := service.ListByGroupID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerGroupID("group-1"),
	)

	if result != nil {
		t.Fatalf(
			"expected nil result, got %v",
			result,
		)
	}

	if !errors.Is(err, ErrForbidden) {
		t.Fatalf(
			"expected ErrForbidden, got %v",
			err,
		)
	}

	if repositoryCalled {
		t.Fatal(
			"expected session repository not to be called when access is denied",
		)
	}
}

func TestListByGroupIDUserNotFound(t *testing.T) {
	repositoryCalled := false

	sessions := &listPrayerSessionRepository{
		listByGroupIDFn: func(
			ctx context.Context,
			groupID identity.PrayerGroupID,
		) ([]domain.PrayerSession, error) {
			repositoryCalled = true
			return nil, nil
		},
	}

	users := &listUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return nil, nil
		},
	}

	service := NewService(sessions, users)

	_, err := service.ListByGroupID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerGroupID("group-1"),
	)

	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}

	if repositoryCalled {
		t.Fatal(
			"expected session repository not to be called when user is not found",
		)
	}
}

func TestListByGroupIDUserRepositoryError(t *testing.T) {
	repoErr := errors.New("find user failed")

	sessions := &listPrayerSessionRepository{}

	users := &listUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return nil, repoErr
		},
	}

	service := NewService(sessions, users)

	_, err := service.ListByGroupID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerGroupID("group-1"),
	)

	if !errors.Is(err, repoErr) {
		t.Fatalf(
			"expected repository error, got %v",
			err,
		)
	}
}

func TestListByGroupIDRepositoryError(t *testing.T) {
	repoErr := errors.New("list sessions failed")

	sessions := &listPrayerSessionRepository{
		listByGroupIDFn: func(
			ctx context.Context,
			groupID identity.PrayerGroupID,
		) ([]domain.PrayerSession, error) {
			return nil, repoErr
		},
	}

	users := &listUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return testUser(
				identity.UserID("user-1"),
				identity.PrayerGroupID("group-1"),
			), nil
		},
	}

	service := NewService(sessions, users)

	_, err := service.ListByGroupID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerGroupID("group-1"),
	)

	if !errors.Is(err, repoErr) {
		t.Fatalf(
			"expected repository error, got %v",
			err,
		)
	}
}