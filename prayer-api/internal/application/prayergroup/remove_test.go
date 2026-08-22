package prayergroup

import (
	"context"
	"errors"
	"testing"

	domainid "prayer-api/internal/domain/identity"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainrole "prayer-api/internal/domain/role"
	domainuser "prayer-api/internal/domain/user"
)

func TestRemovePrayerGroupService_Remove(t *testing.T) {
	roleID := domainid.RoleID("role-1")

	user, err := domainuser.New(
		domainid.UserID("user-1"),
		"external-id",
		"John",
		roleID,
	)
	if err != nil {
		t.Fatal(err)
	}

	group, err := domainprayergroup.New(
		domainid.PrayerGroupID("group-1"),
		"Youth",
		"",
		domainprayergroup.TypeRegular,
	)
	if err != nil {
		t.Fatal(err)
	}

	if err := user.AssignPrayerGroup(group.ID); err != nil {
		t.Fatal(err)
	}

	users := &removeUserRepositoryStub{
		user: user,
	}

	groups := &removePrayerGroupRepositoryStub{
		group: group,
	}

	service := NewRemovePrayerGroupService(
		users,
		groups,
	)

	err = service.Remove(
		context.Background(),
		RemoveCommand{
			UserID:  user.ID,
			GroupID: group.ID,
		},
	)

	if err != nil {
		t.Fatalf("expected nil, got %v", err)
	}

	if user.HasPrayerGroup(group.ID) {
		t.Fatal("expected prayer group to be removed")
	}

	if !users.updated {
		t.Fatal("expected user repository update")
	}
}

func TestRemovePrayerGroupService_UserNotFound(t *testing.T) {
	service := NewRemovePrayerGroupService(
		&removeUserRepositoryStub{
			getErr: domainuser.ErrUserNotFound,
		},
		&removePrayerGroupRepositoryStub{},
	)

	err := service.Remove(
		context.Background(),
		RemoveCommand{},
	)

	if !errors.Is(err, domainuser.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}

func TestRemovePrayerGroupService_PrayerGroupNotFound(t *testing.T) {
	roleID := domainid.RoleID("role")

	user, _ := domainuser.New(
		domainid.UserID("user"),
		"external",
		"John",
		roleID,
	)

	service := NewRemovePrayerGroupService(
		&removeUserRepositoryStub{
			user: user,
		},
		&removePrayerGroupRepositoryStub{
			getErr: domainprayergroup.ErrPrayerGroupNotFound,
		},
	)

	err := service.Remove(
		context.Background(),
		RemoveCommand{},
	)

	if !errors.Is(err, domainprayergroup.ErrPrayerGroupNotFound) {
		t.Fatalf("expected ErrPrayerGroupNotFound, got %v", err)
	}
}

func TestRemovePrayerGroupService_PrayerGroupNotAssigned(t *testing.T) {
	roleID := domainid.RoleID("role")

	user, _ := domainuser.New(
		domainid.UserID("user"),
		"external",
		"John",
		roleID,
	)

	group, _ := domainprayergroup.New(
		domainid.PrayerGroupID("group"),
		"Youth",
		"",
		domainprayergroup.TypeRegular,
	)

	service := NewRemovePrayerGroupService(
		&removeUserRepositoryStub{
			user: user,
		},
		&removePrayerGroupRepositoryStub{
			group: group,
		},
	)

	err := service.Remove(
		context.Background(),
		RemoveCommand{
			UserID:  user.ID,
			GroupID: group.ID,
		},
	)

	if !errors.Is(err, domainuser.ErrPrayerGroupNotAssigned) {
		t.Fatalf("expected ErrPrayerGroupNotAssigned, got %v", err)
	}
}

func TestRemovePrayerGroupService_UpdateFails(t *testing.T) {
	roleID := domainid.RoleID("role")

	user, _ := domainuser.New(
		domainid.UserID("user"),
		"external",
		"John",
		roleID,
	)

	group, _ := domainprayergroup.New(
		domainid.PrayerGroupID("group"),
		"Youth",
		"",
		domainprayergroup.TypeRegular,
	)

	_ = user.AssignPrayerGroup(group.ID)

	expected := errors.New("update failed")

	service := NewRemovePrayerGroupService(
		&removeUserRepositoryStub{
			user:      user,
			updateErr: expected,
		},
		&removePrayerGroupRepositoryStub{
			group: group,
		},
	)

	err := service.Remove(
		context.Background(),
		RemoveCommand{
			UserID:  user.ID,
			GroupID: group.ID,
		},
	)

	if !errors.Is(err, expected) {
		t.Fatalf("expected update error, got %v", err)
	}
}