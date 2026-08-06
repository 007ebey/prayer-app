package auth_test

import (
	"context"
	"errors"
	"testing"

	appauth "prayer-api/internal/application/auth"
	"prayer-api/internal/domain/prayergroup"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
	"prayer-api/internal/repository/memory"
)

func newTestService() (*appauth.Service, *memory.UserRepository, *memory.PrayerGroupRepository) {
	users := memory.NewUserRepository()
	roles := memory.NewRoleRepository()
	groups := memory.NewPrayerGroupRepository()
	ids := memory.NewIDGenerator()

	service := appauth.NewService(
		users,
		roles,
		groups,
		ids,
	)

	return service, users, groups
}

func TestLoginCreatesNewUser(t *testing.T) {
	service, users, _ := newTestService()
	ctx := context.Background()

	result, err := service.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_test_123",
			Name:       "Anna Mary",
		},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !result.Created {
		t.Fatal("expected user to be created")
	}

	if result.User == nil {
		t.Fatal("expected user in result")
	}

	if result.User.ExternalID != "clerk_test_123" {
		t.Fatalf("expected external ID clerk_test_123, got %s", result.User.ExternalID)
	}

	if result.User.Name != "Anna Mary" {
		t.Fatalf("expected name Anna Mary, got %s", result.User.Name)
	}

	if result.User.Status != user.StatusActive {
		t.Fatalf("expected active user, got %s", result.User.Status)
	}

	stored, err := users.FindByExternalID(ctx, "clerk_test_123")
	if err != nil {
		t.Fatalf("failed to retrieve stored user: %v", err)
	}

	if stored == nil {
		t.Fatal("expected user to be persisted")
	}
}

func TestNewUserGetsMembersRole(t *testing.T) {
	service, _, _ := newTestService()

	result, err := service.Login(
		context.Background(),
		appauth.LoginCommand{
			ExternalID: "clerk_member_test",
			Name:       "John Samuel",
		},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.User.RoleIDs) != 1 {
		t.Fatalf("expected exactly one default role, got %d", len(result.User.RoleIDs))
	}

	if result.User.RoleIDs[0] != role.ID("role_members") {
		t.Fatalf("expected role_members, got %s", result.User.RoleIDs[0])
	}
}

func TestNewUserGetsVisitorGroupAccess(t *testing.T) {
	service, _, _ := newTestService()

	result, err := service.Login(
		context.Background(),
		appauth.LoginCommand{
			ExternalID: "clerk_visitor_test",
			Name:       "Robert K",
		},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.PrayerGroupAccess) != 1 {
		t.Fatalf("expected one prayer group access, got %d", len(result.PrayerGroupAccess))
	}

	access := result.PrayerGroupAccess[0]

	if access.GroupID != prayergroup.ID("visitor") {
		t.Fatalf("expected visitor group, got %s", access.GroupID)
	}

	if access.Status != prayergroup.AccessActive {
		t.Fatalf("expected active visitor access, got %s", access.Status)
	}

	if access.UserID != result.User.ID {
		t.Fatalf("expected access user ID %s, got %s", result.User.ID, access.UserID)
	}
}

func TestLoginExistingUserDoesNotCreateAnotherUser(t *testing.T) {
	service, _, _ := newTestService()
	ctx := context.Background()

	first, err := service.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_existing_test",
			Name:       "David Thomas",
		},
	)

	if err != nil {
		t.Fatalf("first login failed: %v", err)
	}

	second, err := service.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_existing_test",
			Name:       "David Thomas",
		},
	)

	if err != nil {
		t.Fatalf("second login failed: %v", err)
	}

	if !first.Created {
		t.Fatal("expected first login to create user")
	}

	if second.Created {
		t.Fatal("expected second login not to create user")
	}

	if first.User.ID != second.User.ID {
		t.Fatalf("expected same user ID, first=%s second=%s", first.User.ID, second.User.ID)
	}
}

func TestBlockedUserCannotLogin(t *testing.T) {
	service, _, _ := newTestService()
	ctx := context.Background()

	first, err := service.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_blocked_test",
			Name:       "Blocked User",
		},
	)

	if err != nil {
		t.Fatalf("initial login failed: %v", err)
	}

	first.User.Block()

	_, err = service.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_blocked_test",
			Name:       "Blocked User",
		},
	)

	if !errors.Is(err, user.ErrUserBlocked) {
		t.Fatalf("expected ErrUserBlocked, got %v", err)
	}
}

func TestLoginRequiresAuthenticatedIdentity(t *testing.T) {
	service, _, _ := newTestService()

	_, err := service.Login(
		context.Background(),
		appauth.LoginCommand{
			ExternalID: "",
			Name:       "Anonymous",
		},
	)

	if !errors.Is(err, appauth.ErrIdentityRequired) {
		t.Fatalf("expected ErrIdentityRequired, got %v", err)
	}
}

func TestExistingUserVisitorAccessIsNotDuplicated(t *testing.T) {
	service, _, _ := newTestService()
	ctx := context.Background()

	_, err := service.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_access_test",
			Name:       "Anna Mary",
		},
	)

	if err != nil {
		t.Fatalf("first login failed: %v", err)
	}

	result, err := service.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_access_test",
			Name:       "Anna Mary",
		},
	)

	if err != nil {
		t.Fatalf("second login failed: %v", err)
	}

	if len(result.PrayerGroupAccess) != 1 {
		t.Fatalf("expected exactly one visitor access, got %d", len(result.PrayerGroupAccess))
	}
}
