package prayergroup

import (
	"context"
	"errors"
	"testing"

	domainidentity "prayer-api/internal/domain/identity"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainrole "prayer-api/internal/domain/role"
	domainuser "prayer-api/internal/domain/user"
)

type assignUserRepositoryStub struct {
	actor      *domainuser.User
	user       *domainuser.User
	findExtErr error
	findIDErr  error
	updateErr  error
	updated    *domainuser.User
}

func (r *assignUserRepositoryStub) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {
	return r.actor, r.findExtErr
}

func (r *assignUserRepositoryStub) FindByID(
	ctx context.Context,
	id domainidentity.UserID,
) (*domainuser.User, error) {
	return r.user, r.findIDErr
}

func (r *assignUserRepositoryStub) Update(
	ctx context.Context,
	user *domainuser.User,
) error {
	r.updated = user
	return r.updateErr
}

type assignPrayerGroupRepositoryStub struct {
	group *domainprayergroup.PrayerGroup
	err   error
}

func (r *assignPrayerGroupRepositoryStub) FindByID(
	ctx context.Context,
	id domainidentity.PrayerGroupID,
) (*domainprayergroup.PrayerGroup, error) {
	return r.group, r.err
}

func (r *assignPrayerGroupRepositoryStub) Create(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	return nil
}

func (r *assignPrayerGroupRepositoryStub) Update(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	return nil
}

func (r *assignPrayerGroupRepositoryStub) Delete(
	ctx context.Context,
	id domainidentity.PrayerGroupID,
) error {
	return nil
}

func (r *assignPrayerGroupRepositoryStub) List(
	ctx context.Context,
	actorID domainidentity.UserID,
	) ([]domainprayergroup.PrayerGroup, error) {
    return nil, nil
}

func (r *assignPrayerGroupRepositoryStub) Save(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	return nil
}

type assignRoleRepositoryStub struct {
	role *domainrole.Role
	err  error
}

func (r *assignRoleRepositoryStub) FindByID(
	ctx context.Context,
	id domainidentity.RoleID,
) (*domainrole.Role, error) {
	return r.role, r.err
}

func adminRole(t *testing.T) *domainrole.Role {
	t.Helper()

	role, err := domainrole.New(
		domainidentity.RoleID("admin"),
		"Administrator",
		"",
		[]domainrole.Permission{
			domainrole.PermissionManagePrayerGroups,
		},
		true,
	)
	if err != nil {
		t.Fatal(err)
	}

	return role
}

func memberRole(t *testing.T) *domainrole.Role {
	t.Helper()

	role, err := domainrole.New(
		domainidentity.RoleID("member"),
		"Member",
		"",
		[]domainrole.Permission{
			domainrole.PermissionJoinPrayerSessions,
		},
		true,
	)
	if err != nil {
		t.Fatal(err)
	}

	return role
}

func TestAssignRequiresAuthenticatedActor(t *testing.T) {
	service := NewAssignPrayerGroupService(
		&assignUserRepositoryStub{},
		&assignPrayerGroupRepositoryStub{},
		&assignRoleRepositoryStub{},
	)

	err := service.Assign(context.Background(), AssignCommand{})

	if !errors.Is(err, ErrUnauthorized) {
		t.Fatalf("expected ErrUnauthorized, got %v", err)
	}
}

func TestAssignReturnsActorNotFound(t *testing.T) {
	service := NewAssignPrayerGroupService(
		&assignUserRepositoryStub{},
		&assignPrayerGroupRepositoryStub{},
		&assignRoleRepositoryStub{},
	)

	err := service.Assign(context.Background(), AssignCommand{
		ActorExternalID: "actor",
	})

	if !errors.Is(err, ErrActorNotFound) {
		t.Fatalf("expected ErrActorNotFound, got %v", err)
	}
}

func TestAssignReturnsRoleRepositoryError(t *testing.T) {
	expected := errors.New("role lookup failed")

	actor := &domainuser.User{
		RoleIDs: []domainidentity.RoleID{
			domainidentity.RoleID("admin"),
		},
	}

	service := NewAssignPrayerGroupService(
		&assignUserRepositoryStub{
			actor: actor,
		},
		&assignPrayerGroupRepositoryStub{},
		&assignRoleRepositoryStub{
			err: expected,
		},
	)

	err := service.Assign(context.Background(), AssignCommand{
		ActorExternalID: "actor",
	})

	if !errors.Is(err, expected) {
		t.Fatalf("expected %v, got %v", expected, err)
	}
}

func TestAssignRequiresPermission(t *testing.T) {
	actor := &domainuser.User{
		RoleIDs: []domainidentity.RoleID{
			domainidentity.RoleID("member"),
		},
	}

	service := NewAssignPrayerGroupService(
		&assignUserRepositoryStub{
			actor: actor,
		},
		&assignPrayerGroupRepositoryStub{},
		&assignRoleRepositoryStub{
			role: memberRole(t),
		},
	)

	err := service.Assign(context.Background(), AssignCommand{
		ActorExternalID: "actor",
	})

	if !errors.Is(err, ErrForbidden) {
		t.Fatalf("expected ErrForbidden, got %v", err)
	}
}

func TestAssignReturnsUserNotFound(t *testing.T) {
	actor := &domainuser.User{
		RoleIDs: []domainidentity.RoleID{
			domainidentity.RoleID("admin"),
		},
	}

	service := NewAssignPrayerGroupService(
		&assignUserRepositoryStub{
			actor: actor,
		},
		&assignPrayerGroupRepositoryStub{},
		&assignRoleRepositoryStub{
			role: adminRole(t),
		},
	)

	err := service.Assign(context.Background(), AssignCommand{
		ActorExternalID: "actor",
		UserID:          domainidentity.UserID("user"),
	})

	if !errors.Is(err, domainuser.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}
}

func TestAssignReturnsPrayerGroupNotFound(t *testing.T) {
	actor := &domainuser.User{
		RoleIDs: []domainidentity.RoleID{
			domainidentity.RoleID("admin"),
		},
	}

	user := &domainuser.User{}

	service := NewAssignPrayerGroupService(
		&assignUserRepositoryStub{
			actor: actor,
			user:  user,
		},
		&assignPrayerGroupRepositoryStub{},
		&assignRoleRepositoryStub{
			role: adminRole(t),
		},
	)

	err := service.Assign(context.Background(), AssignCommand{
		ActorExternalID: "actor",
		UserID:          domainidentity.UserID("user"),
		GroupID:         domainidentity.PrayerGroupID("group"),
	})

	if !errors.Is(err, ErrPrayerGroupNotFound) {
		t.Fatalf("expected ErrPrayerGroupNotFound, got %v", err)
	}
}

func TestAssignUpdatesUser(t *testing.T) {
	actor := &domainuser.User{
		RoleIDs: []domainidentity.RoleID{
			domainidentity.RoleID("admin"),
		},
	}

	user := &domainuser.User{}

	group := &domainprayergroup.PrayerGroup{
		ID: domainidentity.PrayerGroupID("group"),
	}

	users := &assignUserRepositoryStub{
		actor: actor,
		user:  user,
	}

	service := NewAssignPrayerGroupService(
		users,
		&assignPrayerGroupRepositoryStub{
			group: group,
		},
		&assignRoleRepositoryStub{
			role: adminRole(t),
		},
	)

	err := service.Assign(context.Background(), AssignCommand{
		ActorExternalID: "actor",
		UserID:          domainidentity.UserID("user"),
		GroupID:         group.ID,
	})

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if users.updated == nil {
		t.Fatal("expected Update to be called")
	}
}

func TestAssignPropagatesUpdateError(t *testing.T) {
	expected := errors.New("update failed")

	actor := &domainuser.User{
		RoleIDs: []domainidentity.RoleID{
			domainidentity.RoleID("admin"),
		},
	}

	user := &domainuser.User{}

	group := &domainprayergroup.PrayerGroup{
		ID: domainidentity.PrayerGroupID("group"),
	}

	service := NewAssignPrayerGroupService(
		&assignUserRepositoryStub{
			actor:     actor,
			user:      user,
			updateErr: expected,
		},
		&assignPrayerGroupRepositoryStub{
			group: group,
		},
		&assignRoleRepositoryStub{
			role: adminRole(t),
		},
	)

	err := service.Assign(context.Background(), AssignCommand{
		ActorExternalID: "actor",
		UserID:          domainidentity.UserID("user"),
		GroupID:         group.ID,
	})

	if !errors.Is(err, expected) {
		t.Fatalf("expected %v, got %v", expected, err)
	}
}