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

// --------------------
// Stub repositories
// --------------------

type stubUserRepository struct {
	usersByID       map[domainid.UserID]*domainuser.User
	usersByExternal map[string]*domainuser.User

	findByIDErr       error
	findByExternalErr error
}

func newStubUserRepository() *stubUserRepository {
	return &stubUserRepository{
		usersByID:       make(map[domainid.UserID]*domainuser.User),
		usersByExternal: make(map[string]*domainuser.User),
	}
}

func (r *stubUserRepository) FindByID(
	ctx context.Context,
	id domainid.UserID,
) (*domainuser.User, error) {
	if r.findByIDErr != nil {
		return nil, r.findByIDErr
	}

	return r.usersByID[id], nil
}

func (r *stubUserRepository) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {
	if r.findByExternalErr != nil {
		return nil, r.findByExternalErr
	}

	return r.usersByExternal[externalID], nil
}

func (r *stubUserRepository) Save(
	ctx context.Context,
	user domainuser.User,
) error {
	r.usersByID[user.ID] = &user
	r.usersByExternal[user.ExternalID] = &user
	return nil
}

func (r *stubUserRepository) Update(
	ctx context.Context,
	user *domainuser.User,
) error {
	r.usersByID[user.ID] = user
	r.usersByExternal[user.ExternalID] = user
	return nil
}

// --------------------

type stubRoleRepository struct {
	roles map[domainid.RoleID]*domainrole.Role

	findByIDErr error
}

func newStubRoleRepository() *stubRoleRepository {
	return &stubRoleRepository{
		roles: make(map[domainid.RoleID]*domainrole.Role),
	}
}

func (r *stubRoleRepository) FindByID(
	ctx context.Context,
	id domainid.RoleID,
) (*domainrole.Role, error) {
	if r.findByIDErr != nil {
		return nil, r.findByIDErr
	}

	return r.roles[id], nil
}

// --------------------

type stubPrayerGroupRepository struct {
	groups map[domainid.PrayerGroupID]*domainprayergroup.PrayerGroup
	access map[string]domainprayergroup.Access

	findByIDErr   error
	saveErr       error
	listErr       error
	updateErr     error
	deleteErr     error
	findAccessErr error
	saveAccessErr error
}

func newStubPrayerGroupRepository() *stubPrayerGroupRepository {
	return &stubPrayerGroupRepository{
		groups: make(
			map[domainid.PrayerGroupID]*domainprayergroup.PrayerGroup,
		),
		access: make(map[string]domainprayergroup.Access),
	}
}

func (r *stubPrayerGroupRepository) FindByID(
	ctx context.Context,
	id domainid.PrayerGroupID,
) (*domainprayergroup.PrayerGroup, error) {
	if r.findByIDErr != nil {
		return nil, r.findByIDErr
	}

	return r.groups[id], nil
}

func (r *stubPrayerGroupRepository) Save(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	if r.saveErr != nil {
		return r.saveErr
	}

	r.groups[group.ID] = group

	return nil
}

func (r *stubPrayerGroupRepository) List(
	ctx context.Context,
	actorID domainid.UserID,
) ([]domainprayergroup.PrayerGroup, error) {
	if r.listErr != nil {
		return nil, r.listErr
	}

	groups := make(
		[]domainprayergroup.PrayerGroup,
		0,
		len(r.groups),
	)

	for _, group := range r.groups {
		groups = append(groups, *group)
	}

	return groups, nil
}

func (r *stubPrayerGroupRepository) Update(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	if r.updateErr != nil {
		return r.updateErr
	}

	r.groups[group.ID] = group

	return nil
}

func (r *stubPrayerGroupRepository) Delete(
	ctx context.Context,
	id domainid.PrayerGroupID,
) error {
	if r.deleteErr != nil {
		return r.deleteErr
	}

	delete(r.groups, id)

	return nil
}

func (r *stubPrayerGroupRepository) FindAccess(
	ctx context.Context,
	userID domainid.UserID,
	groupID domainid.PrayerGroupID,
) (*domainprayergroup.Access, error) {
	if r.findAccessErr != nil {
		return nil, r.findAccessErr
	}

	key := stubAccessKey(
		userID,
		groupID,
	)

	access, ok := r.access[key]
	if !ok {
		return nil, nil
	}

	return &access, nil
}

func (r *stubPrayerGroupRepository) SaveAccess(
	ctx context.Context,
	access domainprayergroup.Access,
) error {
	if r.saveAccessErr != nil {
		return r.saveAccessErr
	}

	key := stubAccessKey(
		access.UserID,
		access.GroupID,
	)

	r.access[key] = access

	return nil
}

func stubAccessKey(
	userID domainid.UserID,
	groupID domainid.PrayerGroupID,
) string {
	return string(userID) + ":" + string(groupID)
}

// --------------------
// Tests
// --------------------

func TestBlockPrayerGroupService_Block(t *testing.T) {
	ctx := context.Background()

	users := newStubUserRepository()
	groups := newStubPrayerGroupRepository()
	roles := newStubRoleRepository()

	actorID := domainid.UserID("actor-1")
	userID := domainid.UserID("user-1")
	groupID := domainid.PrayerGroupID("group-1")
	roleID := domainid.RoleID("admin-role")

	actor := &domainuser.User{
		ID:         actorID,
		ExternalID: "actor-external-id",
		RoleIDs:    []domainid.RoleID{roleID},
	}

	targetUser := &domainuser.User{
		ID: userID,
	}

	role := &domainrole.Role{
		ID: roleID,
		Permissions: []domainrole.Permission{
			domainrole.PermissionManagePrayerGroups,
		},
	}

	group := &domainprayergroup.PrayerGroup{
		ID: groupID,
	}

	access := domainprayergroup.NewAccess(
		userID,
		groupID,
	)

	users.usersByID[actorID] = actor
	users.usersByExternal[actor.ExternalID] = actor
	users.usersByID[userID] = targetUser

	roles.roles[roleID] = role
	groups.groups[groupID] = group

	groups.access[stubAccessKey(
		userID,
		groupID,
	)] = access

	service := NewBlockPrayerGroupService(
		users,
		groups,
		roles,
	)

	err := service.Block(
		ctx,
		BlockCommand{
			ActorExternalID: actor.ExternalID,
			UserID:          userID,
			GroupID:         groupID,
		},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	savedAccess, err := groups.FindAccess(
		ctx,
		userID,
		groupID,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if savedAccess == nil {
		t.Fatal("expected access to exist")
	}

    if savedAccess.Status != domainprayergroup.AccessBlocked {
	    t.Fatalf(
		  "expected access status %q, got %q",
		  domainprayergroup.AccessBlocked,
		  savedAccess.Status,
	    )
    }
}

func TestBlockPrayerGroupService_Unauthorized(t *testing.T) {
	service := NewBlockPrayerGroupService(
		newStubUserRepository(),
		newStubPrayerGroupRepository(),
		newStubRoleRepository(),
	)

	err := service.Block(
		context.Background(),
		BlockCommand{
			UserID:  domainid.UserID("user-1"),
			GroupID: domainid.PrayerGroupID("group-1"),
		},
	)

	if !errors.Is(err, ErrUnauthorized) {
		t.Fatalf(
			"expected ErrUnauthorized, got %v",
			err,
		)
	}
}

func TestBlockPrayerGroupService_ActorNotFound(t *testing.T) {
	service := NewBlockPrayerGroupService(
		newStubUserRepository(),
		newStubPrayerGroupRepository(),
		newStubRoleRepository(),
	)

	err := service.Block(
		context.Background(),
		BlockCommand{
			ActorExternalID: "missing-actor",
			UserID:          domainid.UserID("user-1"),
			GroupID:         domainid.PrayerGroupID("group-1"),
		},
	)

	if !errors.Is(err, ErrActorNotFound) {
		t.Fatalf(
			"expected ErrActorNotFound, got %v",
			err,
		)
	}
}

func TestBlockPrayerGroupService_Forbidden(t *testing.T) {
	ctx := context.Background()

	users := newStubUserRepository()
	groups := newStubPrayerGroupRepository()
	roles := newStubRoleRepository()

	roleID := domainid.RoleID("member-role")

	actor := &domainuser.User{
		ID:         domainid.UserID("actor-1"),
		ExternalID: "actor-external-id",
		RoleIDs:    []domainid.RoleID{roleID},
	}

	role := &domainrole.Role{
		ID: roleID,
	}

	users.usersByID[actor.ID] = actor
	users.usersByExternal[actor.ExternalID] = actor
	roles.roles[roleID] = role

	service := NewBlockPrayerGroupService(
		users,
		groups,
		roles,
	)

	err := service.Block(
		ctx,
		BlockCommand{
			ActorExternalID: actor.ExternalID,
			UserID:          domainid.UserID("user-1"),
			GroupID:         domainid.PrayerGroupID("group-1"),
		},
	)

	if !errors.Is(err, ErrForbidden) {
		t.Fatalf(
			"expected ErrForbidden, got %v",
			err,
		)
	}
}

func TestBlockPrayerGroupService_UserNotFound(t *testing.T) {
	ctx := context.Background()

	users, groups, roles, actor := setupBlockAdmin()

	service := NewBlockPrayerGroupService(
		users,
		groups,
		roles,
	)

	err := service.Block(
		ctx,
		BlockCommand{
			ActorExternalID: actor.ExternalID,
			UserID:          domainid.UserID("missing-user"),
			GroupID:         domainid.PrayerGroupID("group-1"),
		},
	)

	if !errors.Is(err, domainuser.ErrUserNotFound) {
		t.Fatalf(
			"expected ErrUserNotFound, got %v",
			err,
		)
	}
}

func TestBlockPrayerGroupService_PrayerGroupNotFound(t *testing.T) {
	ctx := context.Background()

	users, groups, roles, actor := setupBlockAdmin()

	user := &domainuser.User{
		ID: domainid.UserID("user-1"),
	}

	users.usersByID[user.ID] = user

	service := NewBlockPrayerGroupService(
		users,
		groups,
		roles,
	)

	err := service.Block(
		ctx,
		BlockCommand{
			ActorExternalID: actor.ExternalID,
			UserID:          user.ID,
			GroupID:         domainid.PrayerGroupID("missing-group"),
		},
	)

	if !errors.Is(
		err,
		domainprayergroup.ErrPrayerGroupNotFound,
	) {
		t.Fatalf(
			"expected ErrPrayerGroupNotFound, got %v",
			err,
		)
	}
}

func TestBlockPrayerGroupService_PrayerGroupNotAssigned(t *testing.T) {
	ctx := context.Background()

	users, groups, roles, actor := setupBlockAdmin()

	user := &domainuser.User{
		ID: domainid.UserID("user-1"),
	}

	group := &domainprayergroup.PrayerGroup{
		ID: domainid.PrayerGroupID("group-1"),
	}

	users.usersByID[user.ID] = user
	groups.groups[group.ID] = group

	service := NewBlockPrayerGroupService(
		users,
		groups,
		roles,
	)

	err := service.Block(
		ctx,
		BlockCommand{
			ActorExternalID: actor.ExternalID,
			UserID:          user.ID,
			GroupID:         group.ID,
		},
	)

	if !errors.Is(
		err,
		domainprayergroup.ErrPrayerGroupNotAssigned,
	) {
		t.Fatalf(
			"expected ErrPrayerGroupNotAssigned, got %v",
			err,
		)
	}
}

func TestBlockPrayerGroupService_SaveAccessError(t *testing.T) {
	ctx := context.Background()

	users, groups, roles, actor := setupBlockAdmin()

	user := &domainuser.User{
		ID: domainid.UserID("user-1"),
	}

	group := &domainprayergroup.PrayerGroup{
		ID: domainid.PrayerGroupID("group-1"),
	}

	users.usersByID[user.ID] = user
	groups.groups[group.ID] = group

	groups.access[stubAccessKey(
		user.ID,
		group.ID,
	)] = domainprayergroup.Access{
		UserID:  user.ID,
		GroupID: group.ID,
	}

	expectedErr := errors.New("save access failed")
	groups.saveAccessErr = expectedErr

	service := NewBlockPrayerGroupService(
		users,
		groups,
		roles,
	)

	err := service.Block(
		ctx,
		BlockCommand{
			ActorExternalID: actor.ExternalID,
			UserID:          user.ID,
			GroupID:         group.ID,
		},
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected save error, got %v",
			err,
		)
	}
}

// --------------------
// Test setup
// --------------------

func setupBlockAdmin() (
	*stubUserRepository,
	*stubPrayerGroupRepository,
	*stubRoleRepository,
	*domainuser.User,
) {
	users := newStubUserRepository()
	groups := newStubPrayerGroupRepository()
	roles := newStubRoleRepository()

	roleID := domainid.RoleID("admin-role")

	role := &domainrole.Role{
		ID: roleID,
		Permissions: []domainrole.Permission{
			domainrole.PermissionManagePrayerGroups,
		},
	}

	actor := &domainuser.User{
		ID:         domainid.UserID("actor-1"),
		ExternalID: "actor-external-id",
		RoleIDs: []domainid.RoleID{
			roleID,
		},
	}

	roles.roles[roleID] = role

	users.usersByID[actor.ID] = actor
	users.usersByExternal[actor.ExternalID] = actor

	return users, groups, roles, actor
}