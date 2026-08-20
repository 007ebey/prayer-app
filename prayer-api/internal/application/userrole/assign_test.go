package userrole_test

import (
	"context"
	"errors"
	"testing"

	appauth "prayer-api/internal/application/auth"
	"prayer-api/internal/application/userrole"
	"prayer-api/internal/domain/user"
	"prayer-api/internal/repository/memory"
	"prayer-api/internal/domain/identity"
)

type testEnvironment struct {
	users     *memory.UserRepository
	roles     *memory.RoleRepository
	login     *appauth.Service
	userRoles *userrole.Service
}

func newTestEnvironment() *testEnvironment {
	users := memory.NewUserRepository()
	roles := memory.NewRoleRepository()
	groups := memory.NewPrayerGroupRepository()
	ids := memory.NewIDGenerator()

	return &testEnvironment{
		users: users,
		roles: roles,
		login: appauth.NewService(
			users,
			roles,
			groups,
			ids,
		),
		userRoles: userrole.NewService(
			users,
			roles,
		),
	}
}

func createUser(
	t *testing.T,
	env *testEnvironment,
	externalID string,
	name string,
) *user.User {
	t.Helper()

	result, err := env.login.Login(
		context.Background(),
		appauth.LoginCommand{
			ExternalID: externalID,
			Name:       name,
		},
	)

	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	return result.User
}

func createAdministrator(
	t *testing.T,
	env *testEnvironment,
	externalID string,
) *user.User {
	t.Helper()

	admin := createUser(
		t,
		env,
		externalID,
		"Administrator",
	)

	if err := admin.AssignRole(identity.RoleID("role_admin")); err != nil {
		t.Fatalf("failed to assign Administrator role: %v", err)
	}

	if err := env.users.Save(context.Background(), admin); err != nil {
		t.Fatalf("failed to save Administrator: %v", err)
	}

	return admin
}

func TestAdministratorCanAssignRoleToUser(t *testing.T) {
	env := newTestEnvironment()

	admin := createAdministrator(
		t,
		env,
		"clerk_admin_assign",
	)

	target := createUser(
		t,
		env,
		"clerk_target_assign",
		"Target User",
	)

	result, err := env.userRoles.Assign(
		context.Background(),
		userrole.AssignCommand{
			ActorExternalID: admin.ExternalID,
			TargetUserID:    target.ID,
			RoleID:          identity.RoleID("role_admin"),
		},
	)

	if err != nil {
		t.Fatalf("expected role assignment to succeed: %v", err)
	}

	if result == nil {
		t.Fatal("expected assignment result")
	}

	if !result.Assigned {
		t.Fatal("expected assigned=true")
	}

	if result.Role == nil {
		t.Fatal("expected resolved role")
	}

	if result.Role.ID != identity.RoleID("role_admin") {
		t.Fatalf("expected role_admin, got %s", result.Role.ID)
	}

	found := false

	for _, roleID := range result.User.RoleIDs {
		if roleID == identity.RoleID("role_admin") {
			found = true
			break
		}
	}

	if !found {
		t.Fatal("expected target user to contain role_admin")
	}
}

func TestMemberCannotAssignRole(t *testing.T) {
	env := newTestEnvironment()

	member := createUser(
		t,
		env,
		"clerk_member_actor",
		"Member Actor",
	)

	target := createUser(
		t,
		env,
		"clerk_member_target",
		"Target User",
	)

	_, err := env.userRoles.Assign(
		context.Background(),
		userrole.AssignCommand{
			ActorExternalID: member.ExternalID,
			TargetUserID:    target.ID,
			RoleID:          identity.RoleID("role_admin"),
		},
	)

	if !errors.Is(err, userrole.ErrForbidden) {
		t.Fatalf("expected ErrForbidden, got %v", err)
	}
}

func TestUnknownActorCannotAssignRole(t *testing.T) {
	env := newTestEnvironment()

	target := createUser(
		t,
		env,
		"clerk_unknown_actor_target",
		"Target User",
	)

	_, err := env.userRoles.Assign(
		context.Background(),
		userrole.AssignCommand{
			ActorExternalID: "clerk_missing_actor",
			TargetUserID:    target.ID,
			RoleID:          identity.RoleID("role_admin"),
		},
	)

	if !errors.Is(err, userrole.ErrForbidden) {
		t.Fatalf("expected ErrForbidden, got %v", err)
	}
}

func TestAssignRoleTargetUserNotFound(t *testing.T) {
	env := newTestEnvironment()

	admin := createAdministrator(
		t,
		env,
		"clerk_admin_missing_target",
	)

	_, err := env.userRoles.Assign(
		context.Background(),
		userrole.AssignCommand{
			ActorExternalID: admin.ExternalID,
			TargetUserID:    identity.UserID("user_missing"),
			RoleID:          identity.RoleID("role_admin"),
		},
	)

	if !errors.Is(err, userrole.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}

func TestAssignRoleRoleNotFound(t *testing.T) {
	env := newTestEnvironment()

	admin := createAdministrator(
		t,
		env,
		"clerk_admin_missing_role",
	)

	target := createUser(
		t,
		env,
		"clerk_missing_role_target",
		"Target User",
	)

	_, err := env.userRoles.Assign(
		context.Background(),
		userrole.AssignCommand{
			ActorExternalID: admin.ExternalID,
			TargetUserID:    target.ID,
			RoleID:          identity.RoleID("role_does_not_exist"),
		},
	)

	if !errors.Is(err, userrole.ErrRoleNotFound) {
		t.Fatalf("expected ErrRoleNotFound, got %v", err)
	}
}

func TestAssignRoleIsIdempotent(t *testing.T) {
	env := newTestEnvironment()

	admin := createAdministrator(
		t,
		env,
		"clerk_admin_idempotent",
	)

	target := createUser(
		t,
		env,
		"clerk_idempotent_target",
		"Target User",
	)

	command := userrole.AssignCommand{
		ActorExternalID: admin.ExternalID,
		TargetUserID:    target.ID,
		RoleID:          identity.RoleID("role_admin"),
	}

	first, err := env.userRoles.Assign(
		context.Background(),
		command,
	)

	if err != nil {
		t.Fatalf("first assignment failed: %v", err)
	}

	if !first.Assigned {
		t.Fatal("expected first assignment assigned=true")
	}

	second, err := env.userRoles.Assign(
		context.Background(),
		command,
	)

	if err != nil {
		t.Fatalf("second assignment failed: %v", err)
	}

	if second.Assigned {
		t.Fatal("expected second assignment assigned=false")
	}

	count := 0

	for _, roleID := range second.User.RoleIDs {
		if roleID == identity.RoleID("role_admin") {
			count++
		}
	}

	if count != 1 {
		t.Fatalf("expected role_admin exactly once, got %d", count)
	}
}
