package prayergroup

import (
	"context"
	"errors"
	"testing"

	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainuser "prayer-api/internal/domain/user"
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

func (r *userRepositoryStub) Find(
	ctx context.Context,
	id domainuser.ID,
) (*domainuser.User, error) {
	return nil, nil
}

func (r *userRepositoryStub) Save(
	ctx context.Context,
	user domainuser.User,
) error {
	return nil
}

type prayerGroupRepositoryStub struct {
	groups []domainprayergroup.PrayerGroup
	err    error
}

func (r *prayerGroupRepositoryStub) List(
	ctx context.Context,
	actorID domainuser.ID,
) ([]domainprayergroup.PrayerGroup, error) {
	return r.groups, r.err
}

// Implement the remaining methods required by PrayerGroupRepository.

func (r *prayerGroupRepositoryStub) FindVisitorGroup(context.Context) (*domainprayergroup.PrayerGroup, error) {
	return nil, nil
}

func (r *prayerGroupRepositoryStub) FindByID(context.Context, domainprayergroup.ID) (*domainprayergroup.PrayerGroup, error) {
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

func (r *prayerGroupRepositoryStub) FindAccess(
	context.Context,
	domainuser.ID,
	domainprayergroup.ID,
) (*domainprayergroup.Access, error) {
	return nil, nil
}

func (r *prayerGroupRepositoryStub) SaveAccess(
	context.Context,
	domainprayergroup.Access,
) error {
	return nil
}

func TestListReturnsPrayerGroups(t *testing.T) {
	userRepo := &userRepositoryStub{
		user: &domainuser.User{
			ID: domainuser.ID("user-1"),
		},
	}

	groupRepo := &prayerGroupRepositoryStub{
		groups: []domainprayergroup.PrayerGroup{
			{
				ID:          domainprayergroup.ID("group-1"),
				Name:        "Youth",
				Description: "Youth Prayer",
			},
			{
				ID:          domainprayergroup.ID("group-2"),
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
			ID: domainuser.ID("user-1"),
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
			ID: domainuser.ID("user-1"),
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