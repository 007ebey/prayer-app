package role_test

import (
	"errors"
	"testing"

	"prayer-api/internal/domain/role"
)

func TestCreateMembersRole(t *testing.T) {
	r, err := role.New(
		role.ID("role_members"),
		"Members",
		"Standard access to prayer sessions.",
		[]role.Permission{
			role.PermissionViewPrayerSessions,
			role.PermissionJoinPrayerSessions,
		},
		true,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !r.HasPermission(role.PermissionViewPrayerSessions) {
		t.Fatal("expected Members role to view prayer sessions")
	}

	if !r.HasPermission(role.PermissionJoinPrayerSessions) {
		t.Fatal("expected Members role to join prayer sessions")
	}

	if r.HasPermission(role.PermissionHostPrayerSessions) {
		t.Fatal("Members role must not host prayer sessions")
	}

	if r.HasPermission(role.PermissionManageUsers) {
		t.Fatal("Members role must not manage users")
	}
}

func TestRoleRejectsUnknownPermission(t *testing.T) {
	_, err := role.New(
		role.ID("role_bad"),
		"Bad Role",
		"Invalid role",
		[]role.Permission{
			role.Permission("users:delete_everything"),
		},
		false,
	)

	if !errors.Is(err, role.ErrInvalidPermission) {
		t.Fatalf("expected ErrInvalidPermission, got %v", err)
	}
}

func TestRoleRemovesDuplicatePermissions(t *testing.T) {
	r, err := role.New(
		role.ID("role_test"),
		"Test",
		"Test role",
		[]role.Permission{
			role.PermissionViewPrayerSessions,
			role.PermissionViewPrayerSessions,
			role.PermissionJoinPrayerSessions,
		},
		false,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(r.Permissions) != 2 {
		t.Fatalf("expected 2 unique permissions, got %d", len(r.Permissions))
	}
}
