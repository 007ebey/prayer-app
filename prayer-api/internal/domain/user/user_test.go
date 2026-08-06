package user_test

import (
	"errors"
	"testing"

	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

func TestNewUserIsActiveAndGetsDefaultRole(t *testing.T) {
	u, err := user.New(
		user.ID("user_1"),
		"clerk_123",
		"Anna Mary",
		role.ID("role_members"),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !u.IsActive() {
		t.Fatal("expected new user to be active")
	}

	if len(u.RoleIDs) != 1 {
		t.Fatalf("expected one role, got %d", len(u.RoleIDs))
	}

	if u.RoleIDs[0] != role.ID("role_members") {
		t.Fatalf("expected role_members, got %s", u.RoleIDs[0])
	}
}

func TestUserCanBeBlockedAndUnblocked(t *testing.T) {
	u, err := user.New(
		user.ID("user_1"),
		"clerk_123",
		"Anna Mary",
		role.ID("role_members"),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	u.Block()

	if u.IsActive() {
		t.Fatal("expected blocked user to be inactive")
	}

	if u.Status != user.StatusBlocked {
		t.Fatalf("expected blocked status, got %s", u.Status)
	}

	u.Unblock()

	if !u.IsActive() {
		t.Fatal("expected user to be active after unblock")
	}
}

func TestDuplicateRoleCannotBeAssigned(t *testing.T) {
	u, err := user.New(
		user.ID("user_1"),
		"clerk_123",
		"Anna Mary",
		role.ID("role_members"),
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = u.AssignRole(role.ID("role_members"))

	if !errors.Is(err, user.ErrRoleAlreadyAssigned) {
		t.Fatalf("expected ErrRoleAlreadyAssigned, got %v", err)
	}
}
