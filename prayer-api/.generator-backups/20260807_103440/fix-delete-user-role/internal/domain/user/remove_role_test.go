package user_test

import (
	"testing"

	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

func TestRemoveRoleRemovesAssignedRole(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	adminRoleID := role.ID("role_admin")
	u.AssignRole(adminRoleID)

	removed := u.RemoveRole(adminRoleID)

	if !removed {
		t.Fatal("expected role to be removed")
	}

	for _, roleID := range u.RoleIDs {
		if roleID == adminRoleID {
			t.Fatal("expected role_admin to be removed")
		}
	}
}

func TestRemoveRoleReturnsFalseWhenRoleIsNotAssigned(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	removed := u.RemoveRole(role.ID("role_admin"))

	if removed {
		t.Fatal("expected false when role was not assigned")
	}
}

func TestRemoveRoleDoesNotRemoveOtherRoles(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	adminRoleID := role.ID("role_admin")
	u.AssignRole(adminRoleID)

	removed := u.RemoveRole(adminRoleID)
	if !removed {
		t.Fatal("expected role_admin to be removed")
	}

	foundMembers := false

	for _, roleID := range u.RoleIDs {
		if roleID == role.ID("role_members") {
			foundMembers = true
		}
	}

	if !foundMembers {
		t.Fatal("removing role_admin must not remove role_members")
	}
}

func newUserForRemoveRoleTest(t *testing.T) *user.User {
	t.Helper()

	u, err := user.New(
		user.ID("user_remove_test"),
		"clerk_remove_test",
		"Remove Test",
		role.ID("role_members"),
	)

	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	return u
}
