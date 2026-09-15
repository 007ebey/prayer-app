package prayersession

import (
	"context"
	"errors"
	"testing"
	"time"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	"prayer-api/internal/domain/user"
)

type mockPrayerSessionRepository struct {
	saveFn func(
		ctx context.Context,
		session *domain.PrayerSession,
	) error
}

func (m *mockPrayerSessionRepository) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	if m.saveFn != nil {
		return m.saveFn(ctx, session)
	}

	return nil
}

func (m *mockPrayerSessionRepository) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	return nil, nil
}

func (m *mockPrayerSessionRepository) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {
	return nil, nil
}

func (m *mockPrayerSessionRepository) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (m *mockPrayerSessionRepository) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	return nil
}

type mockUserRepository struct {
	findByIDFn func(
		ctx context.Context,
		id identity.UserID,
	) (*user.User, error)
}

func (m *mockUserRepository) FindByID(
	ctx context.Context,
	id identity.UserID,
) (*user.User, error) {
	if m.findByIDFn != nil {
		return m.findByIDFn(ctx, id)
	}

	return nil, nil
}

func testUser(
	userID identity.UserID,
	groupID identity.PrayerGroupID,
) *user.User {
	return &user.User{
		ID:             userID,
		PrayerGroupIDs: []identity.PrayerGroupID{groupID},
	}
}

func testCreateCommand() CreateCommand {
	return CreateCommand{
		ActorID:       identity.UserID("user-1"),
		PrayerGroupID: identity.PrayerGroupID("group-1"),
		Title:         "  Morning Prayer  ",
		Date:          time.Date(2026, 9, 20, 0, 0, 0, 0, time.UTC),
		Time:          " 06:00 ",
		Duration:      30,
		PrayerPointIDs: []identity.PrayerPointID{
			identity.PrayerPointID("point-1"),
			identity.PrayerPointID("point-2"),
		},
	}
}

func TestCreateSuccess(t *testing.T) {
	var savedSession *domain.PrayerSession

	sessions := &mockPrayerSessionRepository{
		saveFn: func(
			ctx context.Context,
			session *domain.PrayerSession,
		) error {
			savedSession = session
			return nil
		},
	}

	users := &mockUserRepository{
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

	result, err := service.Create(
		context.Background(),
		testCreateCommand(),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected session, got nil")
	}

	if savedSession == nil {
		t.Fatal("expected session to be saved")
	}

	if savedSession.ID != identity.PrayerSessionID("session-generated") {
		t.Errorf(
			"expected ID %q, got %q",
			"session-generated",
			savedSession.ID,
		)
	}

	if savedSession.PrayerGroupID != identity.PrayerGroupID("group-1") {
		t.Errorf(
			"expected group ID %q, got %q",
			"group-1",
			savedSession.PrayerGroupID,
		)
	}

	if savedSession.Title != "Morning Prayer" {
		t.Errorf(
			"expected trimmed title %q, got %q",
			"Morning Prayer",
			savedSession.Title,
		)
	}

	if savedSession.Time != "06:00" {
		t.Errorf(
			"expected trimmed time %q, got %q",
			"06:00",
			savedSession.Time,
		)
	}

	if savedSession.Duration != 30 {
		t.Errorf(
			"expected duration 30, got %d",
			savedSession.Duration,
		)
	}

	if len(savedSession.PrayerPointIDs) != 2 {
		t.Fatalf(
			"expected 2 prayer points, got %d",
			len(savedSession.PrayerPointIDs),
		)
	}
}

func TestCreateInvalidGroupID(t *testing.T) {
	cmd := testCreateCommand()
	cmd.PrayerGroupID = ""

	service := NewService(
		&mockPrayerSessionRepository{},
		&mockUserRepository{},
	)

	_, err := service.Create(context.Background(), cmd)

	if !errors.Is(err, ErrInvalidGroupID) {
		t.Fatalf(
			"expected ErrInvalidGroupID, got %v",
			err,
		)
	}
}

func TestCreateTitleRequired(t *testing.T) {
	cmd := testCreateCommand()
	cmd.Title = "   "

	service := NewService(
		&mockPrayerSessionRepository{},
		&mockUserRepository{},
	)

	_, err := service.Create(context.Background(), cmd)

	if !errors.Is(err, ErrTitleRequired) {
		t.Fatalf(
			"expected ErrTitleRequired, got %v",
			err,
		)
	}
}

func TestCreateInvalidDate(t *testing.T) {
	cmd := testCreateCommand()
	cmd.Date = time.Time{}

	service := NewService(
		&mockPrayerSessionRepository{},
		&mockUserRepository{},
	)

	_, err := service.Create(context.Background(), cmd)

	if !errors.Is(err, ErrInvalidDate) {
		t.Fatalf(
			"expected ErrInvalidDate, got %v",
			err,
		)
	}
}

func TestCreateTimeRequired(t *testing.T) {
	cmd := testCreateCommand()
	cmd.Time = "   "

	service := NewService(
		&mockPrayerSessionRepository{},
		&mockUserRepository{},
	)

	_, err := service.Create(context.Background(), cmd)

	if !errors.Is(err, ErrTimeRequired) {
		t.Fatalf(
			"expected ErrTimeRequired, got %v",
			err,
		)
	}
}

func TestCreateInvalidDuration(t *testing.T) {
	testCases := []struct {
		name     string
		duration int
	}{
		{
			name:     "zero duration",
			duration: 0,
		},
		{
			name:     "negative duration",
			duration: -10,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cmd := testCreateCommand()
			cmd.Duration = tc.duration

			service := NewService(
				&mockPrayerSessionRepository{},
				&mockUserRepository{},
			)

			_, err := service.Create(
				context.Background(),
				cmd,
			)

			if !errors.Is(err, ErrInvalidDuration) {
				t.Fatalf(
					"expected ErrInvalidDuration, got %v",
					err,
				)
			}
		})
	}
}

func TestCreateUserNotFound(t *testing.T) {
	users := &mockUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return nil, nil
		},
	}

	service := NewService(
		&mockPrayerSessionRepository{},
		users,
	)

	_, err := service.Create(
		context.Background(),
		testCreateCommand(),
	)

	if !errors.Is(err, ErrUserNotFound) {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}

func TestCreateUserRepositoryError(t *testing.T) {
	repoErr := errors.New("user repository failed")

	users := &mockUserRepository{
		findByIDFn: func(
			ctx context.Context,
			id identity.UserID,
		) (*user.User, error) {
			return nil, repoErr
		},
	}

	service := NewService(
		&mockPrayerSessionRepository{},
		users,
	)

	_, err := service.Create(
		context.Background(),
		testCreateCommand(),
	)

	if !errors.Is(err, repoErr) {
		t.Fatalf(
			"expected repository error, got %v",
			err,
		)
	}
}

func TestCreateAccessDenied(t *testing.T) {
	users := &mockUserRepository{
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

	service := NewService(
		&mockPrayerSessionRepository{},
		users,
	)

	_, err := service.Create(
		context.Background(),
		testCreateCommand(),
	)

	if !errors.Is(err, ErrForbidden) {
		t.Fatalf(
			"expected ErrForbidden, got %v",
			err,
		)
	}
}

func TestCreateRepositoryError(t *testing.T) {
	repoErr := errors.New("save failed")

	sessions := &mockPrayerSessionRepository{
		saveFn: func(
			ctx context.Context,
			session *domain.PrayerSession,
		) error {
			return repoErr
		},
	}

	users := &mockUserRepository{
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

	_, err := service.Create(
		context.Background(),
		testCreateCommand(),
	)

	if !errors.Is(err, repoErr) {
		t.Fatalf(
			"expected repository error, got %v",
			err,
		)
	}
}