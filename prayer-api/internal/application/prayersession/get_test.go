package prayersession

import (
	"context"
	"errors"
	"testing"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	"prayer-api/internal/domain/user"
)

type getPrayerSessionRepository struct {
	findByIDFn func(
		ctx context.Context,
		id identity.PrayerSessionID,
	) (*domain.PrayerSession, error)
}

func (m *getPrayerSessionRepository) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}

func (m *getPrayerSessionRepository) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {
	return nil, nil
}

func (m *getPrayerSessionRepository) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (m *getPrayerSessionRepository) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (m *getPrayerSessionRepository) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	return nil
}

type getUserRepository struct {
	findByIDFn func(
		ctx context.Context,
		id identity.UserID,
	) (*user.User, error)
}

func (m *getUserRepository) FindByID(
	ctx context.Context,
	id identity.UserID,
) (*user.User, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}

func testGetSession() *domain.PrayerSession {
	return &domain.PrayerSession{
		ID:             identity.PrayerSessionID("session-1"),
		PrayerGroupID:  identity.PrayerGroupID("group-1"),
		Title:          "Morning Prayer",
		Time:           "06:00",
		Duration:       30,
		PrayerPointIDs: []identity.PrayerPointID{},
	}
}

func TestGetByIDSuccess(t *testing.T) {
	expectedSession := testGetSession()

	sessions := &getPrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			if id != identity.PrayerSessionID("session-1") {
				t.Fatalf(
					"expected session ID %q, got %q",
					"session-1",
					id,
				)
			}

			return expectedSession, nil
		},
	}

	users := &getUserRepository{
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

	result, err := service.GetByID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected session, got nil")
	}

	if result.ID != expectedSession.ID {
		t.Errorf(
			"expected ID %q, got %q",
			expectedSession.ID,
			result.ID,
		)
	}

	if result.PrayerGroupID != expectedSession.PrayerGroupID {
		t.Errorf(
			"expected group ID %q, got %q",
			expectedSession.PrayerGroupID,
			result.PrayerGroupID,
		)
	}

	if result.Title != expectedSession.Title {
		t.Errorf(
			"expected title %q, got %q",
			expectedSession.Title,
			result.Title,
		)
	}
}

func TestGetByIDSessionNotFound(t *testing.T) {
	sessions := &getPrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return nil, nil
		},
	}

	users := &getUserRepository{}

	service := NewService(sessions, users)

	result, err := service.GetByID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if result != nil {
		t.Fatalf("expected nil session, got %v", result)
	}

	if !errors.Is(err, ErrNotFound) {
		t.Fatalf(
			"expected ErrNotFound, got %v",
			err,
		)
	}
}

func TestGetByIDRepositoryError(t *testing.T) {
	repoErr := errors.New("find session failed")

	sessions := &getPrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return nil, repoErr
		},
	}

	users := &getUserRepository{}

	service := NewService(sessions, users)

	_, err := service.GetByID(
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

func TestGetByIDUserNotFound(t *testing.T) {
	sessions := &getPrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return testGetSession(), nil
		},
	}

	users := &getUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return nil, nil
		},
	}

	service := NewService(sessions, users)

	result, err := service.GetByID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if result != nil {
		t.Fatalf("expected nil session, got %v", result)
	}

	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}

func TestGetByIDUserRepositoryError(t *testing.T) {
	repoErr := errors.New("find user failed")

	sessions := &getPrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return testGetSession(), nil
		},
	}

	users := &getUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return nil, repoErr
		},
	}

	service := NewService(sessions, users)

	_, err := service.GetByID(
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

func TestGetByIDAccessDenied(t *testing.T) {
	sessions := &getPrayerSessionRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.PrayerSessionID,
		) (*domain.PrayerSession, error) {
			return testGetSession(), nil
		},
	}

	users := &getUserRepository{
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

	result, err := service.GetByID(
		context.Background(),
		identity.UserID("user-1"),
		identity.PrayerSessionID("session-1"),
	)

	if result != nil {
		t.Fatalf("expected nil session, got %v", result)
	}

	if !errors.Is(err, ErrForbidden) {
		t.Fatalf(
			"expected ErrForbidden, got %v",
			err,
		)
	}
}