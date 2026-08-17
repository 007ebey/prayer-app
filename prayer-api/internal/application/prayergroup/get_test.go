package prayergroup

import (
	"context"
	"errors"
	"testing"

	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainuser "prayer-api/internal/domain/user"
)

type getUserRepositoryStub struct {
	user *domainuser.User
	err  error
}

func (r *getUserRepositoryStub) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {
	return r.user, r.err
}

func (r *getUserRepositoryStub) FindByID(
	ctx context.Context,
	id domainuser.ID,
) (*domainuser.User, error) {
	return nil, nil
}

func (r *getUserRepositoryStub) Save(
	ctx context.Context,
	user domainuser.User,
) error {
	return nil
}

type getPrayerGroupRepositoryStub struct {
	group *domainprayergroup.PrayerGroup
	err   error
}

func (r *getPrayerGroupRepositoryStub) FindByID(
	ctx context.Context,
	id domainprayergroup.ID,
) (*domainprayergroup.PrayerGroup, error) {
	return r.group, r.err
}

func (r *getPrayerGroupRepositoryStub) FindVisitorGroup(
	context.Context,
) (*domainprayergroup.PrayerGroup, error) {
	return nil, nil
}

func (r *getPrayerGroupRepositoryStub) FindByName(
	context.Context,
	string,
) (*domainprayergroup.PrayerGroup, error) {
	return nil, nil
}

func (r *getPrayerGroupRepositoryStub) Save(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	return nil
}

func (r *getPrayerGroupRepositoryStub) Delete(
	ctx context.Context,
	id domainprayergroup.ID,
) error {
	return nil
}

func (r *getPrayerGroupRepositoryStub) List(
	ctx context.Context,
	actorID domainuser.ID,
) ([]domainprayergroup.PrayerGroup, error) {
	return nil, nil
}

func (r *getPrayerGroupRepositoryStub) FindAccess(
	context.Context,
	domainuser.ID,
	domainprayergroup.ID,
) (*domainprayergroup.Access, error) {
	return nil, nil
}

func (r *getPrayerGroupRepositoryStub) SaveAccess(
	context.Context,
	domainprayergroup.Access,
) error {
	return nil
}

func (r *getPrayerGroupRepositoryStub) Update(
	ctx context.Context,
	group *domainprayergroup.PrayerGroup,
) error {
	return nil
}

func TestGetReturnsPrayerGroup(t *testing.T) {
	userRepo := &getUserRepositoryStub{
		user: &domainuser.User{
			ID: domainuser.ID("user-1"),
		},
	}

	groupRepo := &getPrayerGroupRepositoryStub{
		group: &domainprayergroup.PrayerGroup{
			ID:          domainprayergroup.ID("group-1"),
			Name:        "Youth",
			Description: "Youth Prayer",
		},
	}

	service := NewGetService(
		userRepo,
		groupRepo,
	)

	result, err := service.Get(
		context.Background(),
		GetQuery{
			ActorExternalID: "clerk-user",
			GroupID:         domainprayergroup.ID("group-1"),
		},
	)

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if result.PrayerGroup.ID != domainprayergroup.ID("group-1") {
		t.Fatalf("unexpected prayer group id")
	}
}

func TestGetReturnsPrayerGroupNotFound(t *testing.T) {
	userRepo := &getUserRepositoryStub{
		user: &domainuser.User{
			ID: domainuser.ID("user-1"),
		},
	}

	groupRepo := &getPrayerGroupRepositoryStub{
		err: ErrPrayerGroupNotFound,
	}

	service := NewGetService(
		userRepo,
		groupRepo,
	)

	_, err := service.Get(
		context.Background(),
		GetQuery{
			ActorExternalID: "clerk-user",
			GroupID:         domainprayergroup.ID("missing"),
		},
	)

	if !errors.Is(err, ErrPrayerGroupNotFound) {
		t.Fatalf("expected ErrPrayerGroupNotFound")
	}
}

func TestGetReturnsRepositoryError(t *testing.T) {
	expected := errors.New("database unavailable")

	userRepo := &getUserRepositoryStub{
		user: &domainuser.User{
			ID: domainuser.ID("user-1"),
		},
	}

	groupRepo := &getPrayerGroupRepositoryStub{
		err: expected,
	}

	service := NewGetService(
		userRepo,
		groupRepo,
	)

	_, err := service.Get(
		context.Background(),
		GetQuery{
			ActorExternalID: "clerk-user",
			GroupID:         domainprayergroup.ID("group-1"),
		},
	)

	if !errors.Is(err, expected) {
		t.Fatalf("expected %v, got %v", expected, err)
	}
}