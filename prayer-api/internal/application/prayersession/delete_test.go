package prayersession

import (
	"context"
	"errors"
	"testing"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	"prayer-api/internal/domain/user"
)

type deletePrayerSessionRepository struct {
	findByIDFn func(
		ctx context.Context,
		id identity.PrayerSessionID,
	) (*domain.PrayerSession, error)

	deleteFn func(
		ctx context.Context,
		id identity.PrayerSessionID,
	) error
}

func (m *deletePrayerSessionRepository) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}

func (m *deletePrayerSessionRepository) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {
	return nil, nil
}

func (m *deletePrayerSessionRepository) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (m *deletePrayerSessionRepository) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (m *deletePrayerSessionRepository) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	if m.deleteFn != nil {
		return m.deleteFn(ctx, id)
	}

	return nil
}

type deleteUserRepository struct {
	findByIDFn func(
		ctx context.Context,
		id identity.UserID,
	) (*user.User, error)
}

func (m *deleteUserRepository) FindByID(
	ctx context.Context,
	id identity.UserID,
) (*user.User, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}

func testSession() *domain.PrayerSession {
	return &domain.PrayerSession{
		ID:            identity.PrayerSessionID("session-1"),
		PrayerGroupID: identity.PrayerGroupID("group-1"),
		Title:         "Morning Prayer",
	}
}

func TestDeleteSuccess(t *testing.T) {
	var deletedID identity.PrayerSessionID

	sessions := &deletePrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return testSession(), nil
		},
		deleteFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) error {
			deletedID = id
			return nil
		},
	}

	users := &deleteUserRepository{
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

	err := service.Delete(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if deletedID != identity.PrayerSessionID("session-1") {
		t.Fatalf(
			"expected session %q to be deleted, got %q",
			"session-1",
			deletedID,
		)
	}
}

func TestDeleteSessionNotFound(t *testing.T) {
	sessions := &deletePrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return nil, nil
		},
	}

	users := &deleteUserRepository{}

	service := NewService(sessions, users)

	err := service.Delete(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if !errors.Is(err, ErrNotFound) {
		t.Fatalf(
			"expected ErrNotFound, got %v",
			err,
		)
	}
}

func TestDeleteFindSessionRepositoryError(t *testing.T) {
	repoErr := errors.New("find session failed")

	sessions := &deletePrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return nil, repoErr
		},
	}

	users := &deleteUserRepository{}

	service := NewService(sessions, users)

	err := service.Delete(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if !errors.Is(err, repoErr) {
		t.Fatalf(
			"expected repository error, got %v",
			err,
		)
	}
}

func TestDeleteUserNotFound(t *testing.T) {
	sessions := &deletePrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return testSession(), nil
		},
	}

	users := &deleteUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return nil, nil
		},
	}

	service := NewService(sessions, users)

	err := service.Delete(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}

func TestDeleteUserRepositoryError(t *testing.T) {
	repoErr := errors.New("find user failed")

	sessions := &deletePrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return testSession(), nil
		},
	}

	users := &deleteUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return nil, repoErr
		},
	}

	service := NewService(sessions, users)

	err := service.Delete(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if !errors.Is(err, repoErr) {
		t.Fatalf(
			"expected repository error, got %v",
			err,
		)
	}
}

func TestDeleteAccessDenied(t *testing.T) {
	deleteCalled := false

	sessions := &deletePrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return testSession(), nil
		},
		deleteFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) error {
			deleteCalled = true
			return nil
		},
	}

	users := &deleteUserRepository{
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

	err := service.Delete(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if !errors.Is(err, ErrForbidden) {
		t.Fatalf(
			"expected ErrForbidden, got %v",
			err,
		)
	}

	if deleteCalled {
		t.Fatal("expected Delete not to be called when access is denied")
	}
}

func TestDeleteRepositoryError(t *testing.T) {
	repoErr := errors.New("delete failed")

	sessions := &deletePrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return testSession(), nil
		},
		deleteFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) error {
			return repoErr
		},
	}

	users := &deleteUserRepository{
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

	err := service.Delete(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if !errors.Is(err, repoErr) {
		t.Fatalf(
			"expected repository error, got %v",
			err,
		)
	}
}