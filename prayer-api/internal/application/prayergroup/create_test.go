package prayergroup

import (
	"context"
	"testing"
	"errors"
	domainpg "prayer-api/internal/domain/prayergroup"
	domainid "prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
	"prayer-api/internal/repository/memory"
)

type createFixture struct {
	ctx context.Context

	users  *memory.UserRepository
	roles  *memory.RoleRepository
	groups *memory.PrayerGroupRepository
	ids    *memory.IDGenerator

	service *CreateService
}

func newCreateFixture(t *testing.T) *createFixture {
	t.Helper()

	users := memory.NewUserRepository()
	roles := memory.NewRoleRepository()
	groups := memory.NewPrayerGroupRepository()
	ids := memory.NewIDGenerator()

	return &createFixture{
		ctx:    context.Background(),
		users:  users,
		roles:  roles,
		groups: groups,
		ids:    ids,
		service: NewCreateService(
			users,
			roles,
			groups,
			ids,
		),
	}
}

func (f *createFixture) newUser(
	t *testing.T,
	id domainid.UserID,
	externalID string,
	displayName string,
	roleIDs ...domainid.RoleID,
) *user.User {
	t.Helper()

	if len(roleIDs) == 0 {
		t.Fatal("expected at least one role")
	}

	u, err := user.New(
		id,
		externalID,
		displayName,
		roleIDs[0],
	)
	if err != nil {
		t.Fatalf("creating user: %v", err)
	}

	for _, roleID := range roleIDs[1:] {
		if err := u.AssignRole(roleID); err != nil {
			t.Fatalf("assigning role: %v", err)
		}
	}

	if err := f.users.Save(
		f.ctx,
		u,
	); err != nil {
		t.Fatalf("saving user: %v", err)
	}

	return u
}

func (f *createFixture) newAdministrator(t *testing.T) *user.User {
	t.Helper()

	return f.newUser(
		t,
		domainid.UserID("user_admin"),
		"clerk_admin",
		"Administrator",
		domainid.RoleID("role_admin"),
	)
}

func (f *createFixture) newMember(t *testing.T) *user.User {
	t.Helper()

	return f.newUser(
		t,
		domainid.UserID("user_member"),
		"clerk_member",
		"Member",
		domainid.RoleID("role_members"),
	)
}

func (f *createFixture) create(
	t *testing.T,
	actor string,
	name string,
	description string,
) (*CreateResult, error) {
	t.Helper()

	return f.service.Create(
		f.ctx,
		CreateCommand{
			ActorExternalID: actor,
			Name:            name,
			Description:     description,
		},
	)
}

func findGroup(
	t *testing.T,
	repo *memory.PrayerGroupRepository,
	id domainid.PrayerGroupID,
) *domainpg.PrayerGroup {
	t.Helper()

	group, err := repo.FindByID(
		context.Background(),
		id,
	)

	if err != nil {
		t.Fatalf("finding prayer group: %v", err)
	}

	if group == nil {
		t.Fatalf("expected prayer group to exist")
	}

	return group
}

type fixedPrayerGroupIDGenerator struct {
	id domainid.PrayerGroupID
}

func (g fixedPrayerGroupIDGenerator) NewPrayerGroupID() domainid.PrayerGroupID {
	return g.id
}

type saveFailsPrayerGroupRepository struct {
	*memory.PrayerGroupRepository
	err error
}

func (r *saveFailsPrayerGroupRepository) Save(
	ctx context.Context,
	group *domainpg.PrayerGroup,
) error {
	return r.err
}

type roleLookupFailsRepository struct {
	*memory.RoleRepository
	err error
}

func (r *roleLookupFailsRepository) FindByID(
	ctx context.Context,
	id domainid.RoleID,
) (*role.Role, error) {
	return nil, r.err
}

type userLookupFailsRepository struct {
	*memory.UserRepository
	err error
}

func (r *userLookupFailsRepository) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*user.User, error) {
	return nil, r.err
}


func TestAdministratorCanCreatePrayerGroup(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	result, err := f.create(
		t,
		"clerk_admin",
		"Youth Prayer",
		"Weekly youth prayer meeting",
	)

	if err != nil {
		t.Fatalf(
			"expected prayer group creation to succeed: %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected create result")
	}

	if result.PrayerGroup == nil {
		t.Fatal("expected created prayer group")
	}

	if result.PrayerGroup.ID == "" {
		t.Fatal("expected generated prayer group ID")
	}

	if result.PrayerGroup.Name != "Youth Prayer" {
		t.Fatalf(
			"expected name %q, got %q",
			"Youth Prayer",
			result.PrayerGroup.Name,
		)
	}

	if result.PrayerGroup.Description != "Weekly youth prayer meeting" {
		t.Fatalf(
			"expected description %q, got %q",
			"Weekly youth prayer meeting",
			result.PrayerGroup.Description,
		)
	}

	if result.PrayerGroup.Type != domainpg.TypeRegular {
		t.Fatalf(
			"expected type %q, got %q",
			domainpg.TypeRegular,
			result.PrayerGroup.Type,
		)
	}

	if result.PrayerGroup.Status != domainpg.StatusActive {
		t.Fatalf(
			"expected status %q, got %q",
			domainpg.StatusActive,
			result.PrayerGroup.Status,
		)
	}

	persisted := findGroup(
		t,
		f.groups,
		result.PrayerGroup.ID,
	)

	if persisted.Name != result.PrayerGroup.Name {
		t.Fatalf(
			"expected persisted name %q, got %q",
			result.PrayerGroup.Name,
			persisted.Name,
		)
	}

	if persisted.Description != result.PrayerGroup.Description {
		t.Fatalf(
			"expected persisted description %q, got %q",
			result.PrayerGroup.Description,
			persisted.Description,
		)
	}
}

func TestMemberCannotCreatePrayerGroup(t *testing.T) {
	f := newCreateFixture(t)

	f.newMember(t)

	result, err := f.create(
		t,
		"clerk_member",
		"Youth Prayer",
		"Weekly youth prayer meeting",
	)

	if !errors.Is(err, ErrForbidden) {
		t.Fatalf(
			"expected ErrForbidden, got %v",
			err,
		)
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	group, err := f.groups.FindByID(
		f.ctx,
		domainid.PrayerGroupID("group_1"),
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if group != nil {
		t.Fatal(
			"expected prayer group not to be persisted",
		)
	}
}

func TestUnknownActorCannotCreatePrayerGroup(t *testing.T) {
	f := newCreateFixture(t)

	result, err := f.create(
		t,
		"unknown-user",
		"Youth Prayer",
		"Weekly youth prayer meeting",
	)

	if !errors.Is(err, ErrActorNotFound) {
		t.Fatalf(
			"expected ErrActorNotFound, got %v",
			err,
		)
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	// Verify no prayer group was created.
	visitor, err := f.groups.FindByID(
		f.ctx,
		domainid.PrayerGroupID("visitor"),
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if visitor == nil {
		t.Fatal("expected built-in visitor prayer group to exist")
	}

	group, err := f.groups.FindByID(
		f.ctx,
		domainid.PrayerGroupID("group_1"),
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if group != nil {
		t.Fatal("expected no additional prayer group to be created")
	}
}

func TestCreatePrayerGroupNameRequired(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	result, err := f.create(
		t,
		"clerk_admin",
		"",
		"Weekly youth prayer meeting",
	)

	if !errors.Is(err, domainpg.ErrNameRequired) {
		t.Fatalf(
			"expected ErrNameRequired, got %v",
			err,
		)
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	group, err := f.groups.FindByID(
		f.ctx,
		domainid.PrayerGroupID("group_1"),
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if group != nil {
		t.Fatal("expected prayer group not to be persisted")
	}
}

func TestCreatePrayerGroupDescriptionIsPersisted(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	const description = "A weekly prayer gathering for university students."

	result, err := f.create(
		t,
		"clerk_admin",
		"Campus Prayer",
		description,
	)

	if err != nil {
		t.Fatalf(
			"expected prayer group creation to succeed: %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected create result")
	}

	if result.PrayerGroup.Description != description {
		t.Fatalf(
			"expected description %q, got %q",
			description,
			result.PrayerGroup.Description,
		)
	}

	persisted := findGroup(
		t,
		f.groups,
		result.PrayerGroup.ID,
	)

	if persisted.Description != description {
		t.Fatalf(
			"expected persisted description %q, got %q",
			description,
			persisted.Description,
		)
	}
}

func TestCreatePrayerGroupDefaultsToRegularType(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	result, err := f.create(
		t,
		"clerk_admin",
		"Morning Prayer",
		"Daily morning prayer",
	)

	if err != nil {
		t.Fatalf(
			"expected prayer group creation to succeed: %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected create result")
	}

	if result.PrayerGroup.Type != domainpg.TypeRegular {
		t.Fatalf(
			"expected prayer group type %q, got %q",
			domainpg.TypeRegular,
			result.PrayerGroup.Type,
		)
	}

	persisted := findGroup(
		t,
		f.groups,
		result.PrayerGroup.ID,
	)

	if persisted.Type != domainpg.TypeRegular {
		t.Fatalf(
			"expected persisted prayer group type %q, got %q",
			domainpg.TypeRegular,
			persisted.Type,
		)
	}
}

func TestCreatePrayerGroupReceivesGeneratedID(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	result, err := f.create(
		t,
		"clerk_admin",
		"Evening Prayer",
		"Daily evening prayer",
	)

	if err != nil {
		t.Fatalf(
			"expected prayer group creation to succeed: %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected create result")
	}

	if result.PrayerGroup == nil {
		t.Fatal("expected created prayer group")
	}

	if result.PrayerGroup.ID == "" {
		t.Fatal("expected generated prayer group ID")
	}

	expectedID := f.ids.NewPrayerGroupID()

	if result.PrayerGroup.ID == expectedID {
		t.Fatalf(
			"expected created prayer group to use the first generated ID, got subsequent generated ID %q",
			expectedID,
		)
	}

	persisted := findGroup(
		t,
		f.groups,
		result.PrayerGroup.ID,
	)

	if persisted.ID != result.PrayerGroup.ID {
		t.Fatalf(
			"expected persisted ID %q, got %q",
			result.PrayerGroup.ID,
			persisted.ID,
		)
	}
}

func TestCreatePrayerGroupPersistsGroup(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	result, err := f.create(
		t,
		"clerk_admin",
		"Men's Prayer",
		"Friday morning prayer meeting",
	)

	if err != nil {
		t.Fatalf(
			"expected prayer group creation to succeed: %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected create result")
	}

	persisted, err := f.groups.FindByID(
		f.ctx,
		result.PrayerGroup.ID,
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if persisted == nil {
		t.Fatal("expected prayer group to be persisted")
	}

	if persisted.ID != result.PrayerGroup.ID {
		t.Fatalf(
			"expected persisted ID %q, got %q",
			result.PrayerGroup.ID,
			persisted.ID,
		)
	}

	if persisted.Name != result.PrayerGroup.Name {
		t.Fatalf(
			"expected persisted name %q, got %q",
			result.PrayerGroup.Name,
			persisted.Name,
		)
	}

	if persisted.Description != result.PrayerGroup.Description {
		t.Fatalf(
			"expected persisted description %q, got %q",
			result.PrayerGroup.Description,
			persisted.Description,
		)
	}

	if persisted.Type != result.PrayerGroup.Type {
		t.Fatalf(
			"expected persisted type %q, got %q",
			result.PrayerGroup.Type,
			persisted.Type,
		)
	}

	if persisted.Status != result.PrayerGroup.Status {
		t.Fatalf(
			"expected persisted status %q, got %q",
			result.PrayerGroup.Status,
			persisted.Status,
		)
	}
}

func TestCreatePrayerGroupDuplicateIDReturnsConflict(t *testing.T) {
	ctx := context.Background()

	users := memory.NewUserRepository()
	roles := memory.NewRoleRepository()
	groups := memory.NewPrayerGroupRepository()

	admin, err := user.New(
		domainid.UserID("user_admin"),
		"clerk_admin",
		"Administrator",
		domainid.RoleID("role_members"),
	)
	if err != nil {
		t.Fatalf("creating administrator: %v", err)
	}

	if err := admin.AssignRole(domainid.RoleID("role_admin")); err != nil {
		t.Fatalf("assigning administrator role: %v", err)
	}

	if err := users.Save(ctx, admin); err != nil {
		t.Fatalf("saving administrator: %v", err)
	}

	existing, err := domainpg.New(
		domainid.PrayerGroupID("group_duplicate"),
		"Existing Group",
		"Already exists",
		domainpg.TypeRegular,
	)
	if err != nil {
		t.Fatalf("creating existing prayer group: %v", err)
	}

	if err := groups.Save(ctx, existing); err != nil {
		t.Fatalf("saving existing prayer group: %v", err)
	}

	service := NewCreateService(
		users,
		roles,
		groups,
		fixedPrayerGroupIDGenerator{
			id: domainid.PrayerGroupID("group_duplicate"),
		},
	)

	result, err := service.Create(
		ctx,
		CreateCommand{
			ActorExternalID: "clerk_admin",
			Name:            "Youth Prayer",
			Description:     "Weekly youth meeting",
		},
	)

	if !errors.Is(err, ErrPrayerGroupExists) {
		t.Fatalf(
			"expected ErrPrayerGroupExists, got %v",
			err,
		)
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	persisted, err := groups.FindByID(
		ctx,
		domainid.PrayerGroupID("group_duplicate"),
	)
	if err != nil {
		t.Fatalf("finding prayer group: %v", err)
	}

	if persisted == nil {
		t.Fatal("expected original prayer group to remain")
	}

	if persisted.Name != "Existing Group" {
		t.Fatalf(
			"expected existing prayer group to remain unchanged, got %q",
			persisted.Name,
		)
	}
}

func TestCreatePrayerGroupTrimsWhitespaceFromName(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	result, err := f.create(
		t,
		"clerk_admin",
		"   Youth Prayer Group   ",
		"Weekly youth prayer meeting",
	)

	if err != nil {
		t.Fatalf(
			"expected prayer group creation to succeed: %v",
			err,
		)
	}

	if result == nil || result.PrayerGroup == nil {
		t.Fatal("expected created prayer group")
	}

	const expectedName = "Youth Prayer Group"

	if result.PrayerGroup.Name != expectedName {
		t.Fatalf(
			"expected trimmed name %q, got %q",
			expectedName,
			result.PrayerGroup.Name,
		)
	}

	persisted := findGroup(
		t,
		f.groups,
		result.PrayerGroup.ID,
	)

	if persisted.Name != expectedName {
		t.Fatalf(
			"expected persisted trimmed name %q, got %q",
			expectedName,
			persisted.Name,
		)
	}
}

func TestCreatePrayerGroupRejectsWhitespaceOnlyName(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	result, err := f.create(
		t,
		"clerk_admin",
		"      ",
		"Weekly prayer meeting",
	)

	if err == nil {
		t.Fatal("expected validation error")
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	if !errors.Is(err, domainpg.ErrNameRequired) {
		t.Fatalf(
			"expected ErrNameRequired, got %v",
			err,
		)
	}

	groups, err := f.groups.FindByID(
		f.ctx,
		domainid.PrayerGroupID("group_1"),
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if groups != nil {
		t.Fatal("expected no prayer group to be persisted")
	}
}

func TestCreatePrayerGroupDefaultsToActiveStatus(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	result, err := f.create(
		t,
		"clerk_admin",
		"Morning Prayer",
		"Daily morning prayer meeting",
	)

	if err != nil {
		t.Fatalf(
			"expected prayer group creation to succeed: %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected create result")
	}

	if result.PrayerGroup == nil {
		t.Fatal("expected created prayer group")
	}

	if result.PrayerGroup.Status != domainpg.StatusActive {
		t.Fatalf(
			"expected status %q, got %q",
			domainpg.StatusActive,
			result.PrayerGroup.Status,
		)
	}

	persisted := findGroup(
		t,
		f.groups,
		result.PrayerGroup.ID,
	)

	if persisted.Status != domainpg.StatusActive {
		t.Fatalf(
			"expected persisted status %q, got %q",
			domainpg.StatusActive,
			persisted.Status,
		)
	}
}

func TestCreatePrayerGroupCannotCreateWhenRepositoryFails(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	expectedErr := errors.New("repository unavailable")

	service := NewCreateService(
		f.users,
		f.roles,
		&saveFailsPrayerGroupRepository{
			PrayerGroupRepository: f.groups,
			err:                  expectedErr,
		},
		f.ids,
	)

	result, err := service.Create(
		f.ctx,
		CreateCommand{
			ActorExternalID: "clerk_admin",
			Name:            "Morning Prayer",
			Description:     "Daily prayer",
		},
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected repository error %v, got %v",
			expectedErr,
			err,
		)
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	persisted, err := f.groups.FindByID(
		f.ctx,
		domainid.PrayerGroupID("group_1"),
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if persisted != nil {
		t.Fatal("expected prayer group not to be persisted")
	}
}

func TestCreatePrayerGroupCannotCreateWhenRoleLookupFails(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	expectedErr := errors.New("role repository unavailable")

	service := NewCreateService(
		f.users,
		&roleLookupFailsRepository{
			RoleRepository: f.roles,
			err:            expectedErr,
		},
		f.groups,
		f.ids,
	)

	result, err := service.Create(
		f.ctx,
		CreateCommand{
			ActorExternalID: "clerk_admin",
			Name:            "Youth Prayer",
			Description:     "Weekly youth prayer meeting",
		},
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected %v, got %v",
			expectedErr,
			err,
		)
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	group, err := f.groups.FindByID(
		f.ctx,
		domainid.PrayerGroupID("group_1"),
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if group != nil {
		t.Fatal("expected prayer group not to be persisted")
	}
}

func TestCreatePrayerGroupCannotCreateWhenUserLookupFails(t *testing.T) {
	f := newCreateFixture(t)

	expectedErr := errors.New("user repository unavailable")

	service := NewCreateService(
		&userLookupFailsRepository{
			UserRepository: f.users,
			err:            expectedErr,
		},
		f.roles,
		f.groups,
		f.ids,
	)

	result, err := service.Create(
		f.ctx,
		CreateCommand{
			ActorExternalID: "clerk_admin",
			Name:            "Youth Prayer",
			Description:     "Weekly youth prayer meeting",
		},
	)

	if !errors.Is(err, expectedErr) {
		t.Fatalf(
			"expected %v, got %v",
			expectedErr,
			err,
		)
	}

	if result != nil {
		t.Fatal("expected nil result")
	}

	group, err := f.groups.FindByID(
		f.ctx,
		domainid.PrayerGroupID("group_1"),
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if group != nil {
		t.Fatal("expected prayer group not to be persisted")
	}
}

func TestCreatePrayerGroupUsesGeneratedID(t *testing.T) {
	f := newCreateFixture(t)

	f.newAdministrator(t)

	const generatedID = domainid.PrayerGroupID("group_test_123")

	service := NewCreateService(
		f.users,
		f.roles,
		f.groups,
		fixedPrayerGroupIDGenerator{
			id: generatedID,
		},
	)

	result, err := service.Create(
		f.ctx,
		CreateCommand{
			ActorExternalID: "clerk_admin",
			Name:            "Youth Prayer",
			Description:     "Weekly youth prayer meeting",
		},
	)

	if err != nil {
		t.Fatalf(
			"expected prayer group creation to succeed: %v",
			err,
		)
	}

	if result == nil {
		t.Fatal("expected create result")
	}

	if result.PrayerGroup == nil {
		t.Fatal("expected created prayer group")
	}

	if result.PrayerGroup.ID != generatedID {
		t.Fatalf(
			"expected generated ID %q, got %q",
			generatedID,
			result.PrayerGroup.ID,
		)
	}

	persisted, err := f.groups.FindByID(
		f.ctx,
		generatedID,
	)

	if err != nil {
		t.Fatalf(
			"unexpected repository error: %v",
			err,
		)
	}

	if persisted == nil {
		t.Fatal("expected prayer group to be persisted")
	}

	if persisted.ID != generatedID {
		t.Fatalf(
			"expected persisted ID %q, got %q",
			generatedID,
			persisted.ID,
		)
	}
}