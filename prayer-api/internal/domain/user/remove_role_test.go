package user_test

import (
	"errors"
	"testing"

	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

func TestRemoveRoleRemovesAssignedRole(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	adminRoleID := role.ID("role_admin")

	u.AssignRole(adminRoleID)

	if !hasRoleForRemoveTest(u, adminRoleID) {
		t.Fatal("expected role_admin to be assigned before removal")
	}

	err := u.RemoveRole(adminRoleID)

	if err != nil {
		t.Fatalf("expected role removal to succeed: %v", err)
	}

	if hasRoleForRemoveTest(u, adminRoleID) {
		t.Fatal("expected role_admin to be removed")
	}
}

func TestRemoveRoleNotAssignedDoesNotMutateRoles(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	before := append(
		[]role.ID(nil),
		u.RoleIDs...,
	)

	err := u.RemoveRole(
		role.ID("role_admin"),
	)

	if !errors.Is(
		err,
		user.ErrRoleNotAssigned,
	) {
		t.Fatalf(
			"expected ErrRoleNotAssigned, got %v",
			err,
		)
	}

	if len(u.RoleIDs) != len(before) {
		t.Fatalf(
			"expected role count to remain %d, got %d",
			len(before),
			len(u.RoleIDs),
		)
	}

	for index := range before {
		if u.RoleIDs[index] != before[index] {
			t.Fatalf(
				"roles mutated at index %d: expected %s, got %s",
				index,
				before[index],
				u.RoleIDs[index],
			)
		}
	}
}

func TestRemoveRoleDoesNotRemoveOtherRoles(t *testing.T) {
	u := newUserForRemoveRoleTest(t)

	adminRoleID := role.ID("role_admin")
	memberRoleID := role.ID("role_members")

	u.AssignRole(adminRoleID)

	err := u.RemoveRole(adminRoleID)

	if err != nil {
		t.Fatalf(
			"expected role_admin removal to succeed: %v",
			err,
		)
	}

	if hasRoleForRemoveTest(
		u,
		adminRoleID,
	) {
		t.Fatal(
			"expected role_admin to be removed",
		)
	}

	if !hasRoleForRemoveTest(
		u,
		memberRoleID,
	) {
		t.Fatal(
			"removing role_admin must not remove role_members",
		)
	}
}

func newUserForRemoveRoleTest(
	t *testing.T,
) *user.User {
	t.Helper()

	u, err := user.New(
		user.ID("user_remove_test"),
		"clerk_remove_test",
		"Remove Test",
		role.ID("role_members"),
	)

	if err != nil {
		t.Fatalf(
			"failed to create user: %v",
			err,
		)
	}

	return u
}

func hasRoleForRemoveTest(
	u *user.User,
	roleID role.ID,
) bool {
	for _, current := range u.RoleIDs {
		if current == roleID {
			return true
		}
	}

	return false
}
