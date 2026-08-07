package userrole_test

import (
	"context"
	"errors"
	"testing"

	"prayer-api/internal/application/userrole"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
	"prayer-api/internal/repository/memory"
)

func TestAdministratorCanRemoveRoleFromUser(t *testing.T) {
	service, users := newRemoveRoleService(t)

	actor := createRemoveTestUser(t, users, "actor", "clerk_admin")
	actor.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, actor)

	target := createRemoveTestUser(t, users, "target", "clerk_target")
	target.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, target)

	result, err := service.Remove(
		context.Background(),
		userrole.RemoveCommand{
			ActorExternalID: "clerk_admin",
			TargetUserID:    target.ID,
			RoleID:          role.ID("role_admin"),
		},
	)

	if err != nil {
		t.Fatalf("expected role removal to succeed: %v", err)
	}

	if !result.Removed {
		t.Fatal("expected removed=true")
	}

	if hasRemoveTestRole(result.User, role.ID("role_admin")) {
		t.Fatal("expected role_admin to be removed")
	}
}

func TestMemberCannotRemoveRole(t *testing.T) {
	service, users := newRemoveRoleService(t)

	actor := createRemoveTestUser(t, users, "member", "clerk_member")
	mustSaveRemoveTestUser(t, users, actor)

	target := createRemoveTestUser(t, users, "target", "clerk_target")
	target.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, target)

	_, err := service.Remove(
		context.Background(),
		userrole.RemoveCommand{
			ActorExternalID: "clerk_member",
			TargetUserID:    target.ID,
			RoleID:          role.ID("role_admin"),
		},
	)

	if !errors.Is(err, userrole.ErrForbidden) {
		t.Fatalf("expected ErrForbidden, got %v", err)
	}

	stored, err := users.FindByID(context.Background(), target.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !hasRemoveTestRole(stored, role.ID("role_admin")) {
		t.Fatal("forbidden removal must not mutate target roles")
	}
}

func TestUnknownActorCannotRemoveRole(t *testing.T) {
	service, users := newRemoveRoleService(t)

	target := createRemoveTestUser(t, users, "target", "clerk_target")
	target.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, target)

	_, err := service.Remove(
		context.Background(),
		userrole.RemoveCommand{
			ActorExternalID: "unknown_clerk_user",
			TargetUserID:    target.ID,
			RoleID:          role.ID("role_admin"),
		},
	)

	if !errors.Is(err, userrole.ErrForbidden) {
		t.Fatalf("expected ErrForbidden, got %v", err)
	}
}

func TestRemoveRoleTargetUserNotFound(t *testing.T) {
	service, users := newRemoveRoleService(t)

	actor := createRemoveTestUser(t, users, "actor", "clerk_admin")
	actor.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, actor)

	_, err := service.Remove(
		context.Background(),
		userrole.RemoveCommand{
			ActorExternalID: "clerk_admin",
			TargetUserID:    user.ID("missing_user"),
			RoleID:          role.ID("role_admin"),
		},
	)

	if !errors.Is(err, userrole.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}

func TestRemoveRoleRoleNotFound(t *testing.T) {
	service, users := newRemoveRoleService(t)

	actor := createRemoveTestUser(t, users, "actor", "clerk_admin")
	actor.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, actor)

	target := createRemoveTestUser(t, users, "target", "clerk_target")
	mustSaveRemoveTestUser(t, users, target)

	_, err := service.Remove(
		context.Background(),
		userrole.RemoveCommand{
			ActorExternalID: "clerk_admin",
			TargetUserID:    target.ID,
			RoleID:          role.ID("role_missing"),
		},
	)

	if !errors.Is(err, userrole.ErrRoleNotFound) {
		t.Fatalf("expected ErrRoleNotFound, got %v", err)
	}
}

func TestRemoveMembersRoleIsProtected(t *testing.T) {
	service, users := newRemoveRoleService(t)

	actor := createRemoveTestUser(t, users, "actor", "clerk_admin")
	actor.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, actor)

	target := createRemoveTestUser(t, users, "target", "clerk_target")
	mustSaveRemoveTestUser(t, users, target)

	_, err := service.Remove(
		context.Background(),
		userrole.RemoveCommand{
			ActorExternalID: "clerk_admin",
			TargetUserID:    target.ID,
			RoleID:          role.ID("role_members"),
		},
	)

	if !errors.Is(err, userrole.ErrProtectedRole) {
		t.Fatalf("expected ErrProtectedRole, got %v", err)
	}

	stored, err := users.FindByID(context.Background(), target.ID)
	if err != nil {
		t.Fatal(err)
	}

	if !hasRemoveTestRole(stored, role.ID("role_members")) {
		t.Fatal("protected Members role must remain assigned")
	}
}

func TestRemoveRoleIsIdempotent(t *testing.T) {
	service, users := newRemoveRoleService(t)

	actor := createRemoveTestUser(t, users, "actor", "clerk_admin")
	actor.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, actor)

	target := createRemoveTestUser(t, users, "target", "clerk_target")
	target.AssignRole(role.ID("role_admin"))
	mustSaveRemoveTestUser(t, users, target)

	command := userrole.RemoveCommand{
		ActorExternalID: "clerk_admin",
		TargetUserID:    target.ID,
		RoleID:          role.ID("role_admin"),
	}

	first, err := service.Remove(context.Background(), command)
	if err != nil {
		t.Fatalf("first removal failed: %v", err)
	}

	if !first.Removed {
		t.Fatal("expected first removal to return removed=true")
	}

	second, err := service.Remove(context.Background(), command)
	if err != nil {
		t.Fatalf("second removal failed: %v", err)
	}

	if second.Removed {
		t.Fatal("expected second removal to return removed=false")
	}
}

func newRemoveRoleService(t *testing.T) (*userrole.Service, *memory.UserRepository) {
	t.Helper()

	users := memory.NewUserRepository()
	roles := memory.NewRoleRepository()

	return userrole.NewService(users, roles), users
}

func createRemoveTestUser(
	t *testing.T,
	repository *memory.UserRepository,
	id string,
	externalID string,
) *user.User {
	t.Helper()

	u, err := user.New(
		user.ID(id),
		externalID,
		id,
		role.ID("role_members"),
	)

	if err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	mustSaveRemoveTestUser(t, repository, u)

	return u
}

func mustSaveRemoveTestUser(
	t *testing.T,
	repository *memory.UserRepository,
	u *user.User,
) {
	t.Helper()

	if err := repository.Save(context.Background(), u); err != nil {
		t.Fatalf("failed to save test user: %v", err)
	}
}

func hasRemoveTestRole(u *user.User, roleID role.ID) bool {
	for _, current := range u.RoleIDs {
		if current == roleID {
			return true
		}
	}

	return false
}
