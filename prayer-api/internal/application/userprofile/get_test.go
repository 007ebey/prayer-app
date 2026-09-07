package userprofile_test

import (
	"context"
	"errors"
	"testing"

	appauth "prayer-api/internal/application/auth"
	"prayer-api/internal/application/userprofile"
	"prayer-api/internal/domain/identity"
	"prayer-api/internal/repository/memory"
	"prayer-api/internal/config"
)

type testEnvironment struct {
	login    *appauth.Service
	profiles *userprofile.Service
}

func newTestEnvironment() *testEnvironment {
	users := memory.NewUserRepository(config.Config{})
	roles := memory.NewRoleRepository()
	groups := memory.NewPrayerGroupRepository()
	ids := memory.NewIDGenerator()

	return &testEnvironment{
		login: appauth.NewService(
			users,
			roles,
			groups,
			ids,
		),
		profiles: userprofile.NewService(
			users,
			roles,
			groups,
		),
	}
}

func TestGetOwnProfile(t *testing.T) {
	env := newTestEnvironment()
	ctx := context.Background()

	loginResult, err := env.login.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_profile_1",
			Email:      identity.Email("anna.mary@example.com"),
			Name:       "Anna Mary",
		},
	)

	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	result, err := env.profiles.Get(
		ctx,
		"clerk_profile_1",
		loginResult.User.ID,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("expected profile result")
	}

	if result.User == nil {
		t.Fatal("expected user in profile result")
	}

	if result.User.ID != loginResult.User.ID {
		t.Fatalf(
			"expected user %s, got %s",
			loginResult.User.ID,
			result.User.ID,
		)
	}

	if result.User.Name != "Anna Mary" {
		t.Fatalf(
			"expected Anna Mary, got %s",
			result.User.Name,
		)
	}

	if result.User.Email != "anna.mary@example.com" {
		t.Fatalf(
			"expected anna.mary@example.com, got %s",
			result.User.Email,
		)
	}
}

func TestGetOwnProfileReturnsMembersRole(t *testing.T) {
	env := newTestEnvironment()
	ctx := context.Background()

	loginResult, err := env.login.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_role_profile",
			Email:      identity.Email("role.user@example.com"),
			Name:       "Role User",
		},
	)

	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	result, err := env.profiles.Get(
		ctx,
		"clerk_role_profile",
		loginResult.User.ID,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.Roles) != 1 {
		t.Fatalf(
			"expected one role, got %d",
			len(result.Roles),
		)
	}

	resolvedRole := result.Roles[0]

	if resolvedRole.ID != "role_members" {
		t.Fatalf(
			"expected role_members, got %s",
			resolvedRole.ID,
		)
	}

	if resolvedRole.Name != "Members" {
		t.Fatalf(
			"expected Members, got %s",
			resolvedRole.Name,
		)
	}

	if len(resolvedRole.Permissions) == 0 {
		t.Fatal("expected Members role permissions")
	}
}

func TestGetOwnProfileReturnsVisitorGroup(t *testing.T) {
	env := newTestEnvironment()
	ctx := context.Background()

	loginResult, err := env.login.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_group_profile",
			Email:      identity.Email("group.user@example.com"),
			Name:       "Group User",
		},
	)

	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	result, err := env.profiles.Get(
		ctx,
		"clerk_group_profile",
		loginResult.User.ID,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(result.PrayerGroups) != 1 {
		t.Fatalf(
			"expected one prayer group, got %d",
			len(result.PrayerGroups),
		)
	}

	group := result.PrayerGroups[0]

	if group.ID != "visitor" {
		t.Fatalf(
			"expected visitor group, got %s",
			group.ID,
		)
	}

	if group.Name != "Visitor" {
		t.Fatalf(
			"expected Visitor, got %s",
			group.Name,
		)
	}

	if group.AccessStatus != "active" {
		t.Fatalf(
			"expected active access, got %s",
			group.AccessStatus,
		)
	}
}

func TestGetOtherUserIsForbidden(t *testing.T) {
	env := newTestEnvironment()
	ctx := context.Background()

	actor, err := env.login.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_actor",
			Email:      identity.Email("actor@example.com"),
			Name:       "Actor",
		},
	)

	if err != nil {
		t.Fatalf("actor login failed: %v", err)
	}

	target, err := env.login.Login(
		ctx,
		appauth.LoginCommand{
			ExternalID: "clerk_target",
			Email:      identity.Email("target@example.com"),
			Name:       "Target",
		},
	)

	if err != nil {
		t.Fatalf("target login failed: %v", err)
	}

	if actor.User.ID == target.User.ID {
		t.Fatal("expected actor and target to be different users")
	}

	_, err = env.profiles.Get(
		ctx,
		actor.User.ExternalID,
		target.User.ID,
	)

	if !errors.Is(err, userprofile.ErrForbidden) {
		t.Fatalf(
			"expected ErrForbidden, got %v",
			err,
		)
	}
}

func TestGetUnknownActorReturnsNotFound(t *testing.T) {
	env := newTestEnvironment()

	_, err := env.profiles.Get(
		context.Background(),
		"clerk_unknown",
		identity.UserID("user_999"),
	)

	if !errors.Is(err, userprofile.ErrUserNotFound) {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}