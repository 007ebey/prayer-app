package userprofile_test

import (
	"context"
	"errors"
	"testing"

	appauth "prayer-api/internal/application/auth"
	"prayer-api/internal/application/userprofile"
	"prayer-api/internal/domain/user"
	"prayer-api/internal/repository/memory"
)

type testEnvironment struct {
	login    *appauth.Service
	profiles *userprofile.Service
}

func newTestEnvironment() *testEnvironment {
	users := memory.NewUserRepository()
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

	loginResult, err := env.login.Login(ctx, appauth.LoginCommand{
		ExternalID: "clerk_profile_1",
		Name:       "Anna Mary",
	})

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

	if result.User.ID != loginResult.User.ID {
		t.Fatalf("expected user %s, got %s", loginResult.User.ID, result.User.ID)
	}

	if len(result.Roles) != 1 {
		t.Fatalf("expected one role, got %d", len(result.Roles))
	}

	if result.Roles[0].ID != "role_members" {
		t.Fatalf("expected role_members, got %s", result.Roles[0].ID)
	}

	if len(result.PrayerGroups) != 1 {
		t.Fatalf("expected one prayer group, got %d", len(result.PrayerGroups))
	}

	if result.PrayerGroups[0].ID != "visitor" {
		t.Fatalf("expected visitor group, got %s", result.PrayerGroups[0].ID)
	}
}

func TestGetOtherUserIsForbidden(t *testing.T) {
	env := newTestEnvironment()
	ctx := context.Background()

	actor, err := env.login.Login(ctx, appauth.LoginCommand{
		ExternalID: "clerk_actor",
		Name:       "Actor",
	})

	if err != nil {
		t.Fatal(err)
	}

	target, err := env.login.Login(ctx, appauth.LoginCommand{
		ExternalID: "clerk_target",
		Name:       "Target",
	})

	if err != nil {
		t.Fatal(err)
	}

	if actor.User.ID == target.User.ID {
		t.Fatal("expected different users")
	}

	_, err = env.profiles.Get(
		ctx,
		"clerk_actor",
		target.User.ID,
	)

	if !errors.Is(err, userprofile.ErrForbidden) {
		t.Fatalf("expected ErrForbidden, got %v", err)
	}
}

func TestGetUnknownActorReturnsNotFound(t *testing.T) {
	env := newTestEnvironment()

	_, err := env.profiles.Get(
		context.Background(),
		"clerk_unknown",
		user.ID("user_999"),
	)

	if !errors.Is(err, userprofile.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}
