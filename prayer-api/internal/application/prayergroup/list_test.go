package prayergroup

import (
	"context"
	"errors"
	"testing"

	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainuser "prayer-api/internal/domain/user"
	domainid "prayer-api/internal/domain/identity"
)

type userRepositoryStub struct {
	user *domainuser.User
	err  error
}

func (r *userRepositoryStub) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {
	return r.user, r.err
}

func (r *userRepositoryStub) FindByID(
	ctx context.Context,
	id domainid.UserID,
) (*domainuser.User, error) {
	return nil, nil
}

func (r *userRepositoryStub) Save(
	ctx context.Context,
	user domainuser.User,
) error {
	return nil
}

func (r *userRepositoryStub) Update(
	ctx context.Context,
	user *domainuser.User,
) error {
	return nil
}

type prayerGroupRepositoryStub struct {
	groups []domainprayergroup.PrayerGroup
	err    error
}

func (r *prayerGroupRepositoryStub) List(
	ctx context.Context,
	actorID domainid.UserID,
) ([]domainprayergroup.PrayerGroup, error) {
	return r.groups, r.err
}

// Implement the remaining methods required by PrayerGroupRepository.

func (r *prayerGroupRepositoryStub) FindVisitorGroup(context.Context) (*domainprayergroup.PrayerGroup, error) {
	return nil, nil
}

func (r *prayerGroupRepositoryStub) FindByID(context.Context, domainid.PrayerGroupID) (*domainprayergroup.PrayerGroup, error) {
	return nil, nil
}

func (r *prayerGroupRepositoryStub) FindByName(context.Context, string) (*domainprayergroup.PrayerGroup, error) {
	return nil, nil
}

func (r *prayerGroupRepositoryStub) Save(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	return nil
}

func (r *prayerGroupRepositoryStub) Delete(
	ctx context.Context,
	id domainid.PrayerGroupID,
) error {
	return nil
}

func (r *prayerGroupRepositoryStub) FindAccess(
	context.Context,
	domainid.UserID,
	domainid.PrayerGroupID,
) (*domainprayergroup.Access, error) {
	return nil, nil
}

func (r *prayerGroupRepositoryStub) SaveAccess(
	context.Context,
	domainprayergroup.Access,
) error {
	return nil
}

func (r *prayerGroupRepositoryStub) Update(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	return nil
}

func TestListReturnsPrayerGroups(t *testing.T) {
	userRepo := &userRepositoryStub{
		user: &domainuser.User{
			ID: domainid.UserID("user-1"),
		},
	}

	groupRepo := &prayerGroupRepositoryStub{
		groups: []domainprayergroup.PrayerGroup{
			{
				ID:          domainid.PrayerGroupID("group-1"),
				Name:        "Youth",
				Description: "Youth Prayer",
			},
			{
				ID:          domainid.PrayerGroupID("group-2"),
				Name:        "Family",
				Description: "Family Prayer",
			},
		},
	}

	service := NewListService(userRepo, groupRepo)

	result, err := service.List(
		context.Background(),
		ListQuery{
			ActorExternalID: "clerk-user",
		},
	)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(result.PrayerGroups) != 2 {
		t.Fatalf("expected 2 prayer groups, got %d", len(result.PrayerGroups))
	}
}

func TestListReturnsEmptyList(t *testing.T) {
	userRepo := &userRepositoryStub{
		user: &domainuser.User{
			ID: domainid.UserID("user-1"),
		},
	}

	groupRepo := &prayerGroupRepositoryStub{}

	service := NewListService(userRepo, groupRepo)

	result, err := service.List(
		context.Background(),
		ListQuery{
			ActorExternalID: "clerk-user",
		},
	)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(result.PrayerGroups) != 0 {
		t.Fatalf("expected empty list")
	}
}

func TestListReturnsRepositoryError(t *testing.T) {
	expected := errors.New("database unavailable")

	userRepo := &userRepositoryStub{
		user: &domainuser.User{
			ID: domainid.UserID("user-1"),
		},
	}

	groupRepo := &prayerGroupRepositoryStub{
		err: expected,
	}

	service := NewListService(userRepo, groupRepo)

	_, err := service.List(
		context.Background(),
		ListQuery{
			ActorExternalID: "clerk-user",
		},
	)

	if !errors.Is(err, expected) {
		t.Fatalf("expected %v, got %v", expected, err)
	}
}