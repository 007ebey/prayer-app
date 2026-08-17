package prayergroup

import (
	"context"
	"errors"
	"testing"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainuser "prayer-api/internal/domain/user"
)

type updatePrayerGroupRepositoryStub struct {
	group       *domainprayergroup.PrayerGroup
	findErr     error
	updateErr   error
	updated     *domainprayergroup.PrayerGroup
	updateCalls int
}

func (r *updatePrayerGroupRepositoryStub) FindByID(
    ctx context.Context,
    id domainprayergroup.ID,
) (*domainprayergroup.PrayerGroup, error) {
    return r.group, r.findErr
}

func (r *updatePrayerGroupRepositoryStub) Update(
    ctx context.Context,
    group *domainprayergroup.PrayerGroup,
) error {
    r.updateCalls++
    r.updated = group
    return r.updateErr
}

func (r *updatePrayerGroupRepositoryStub) List(
    ctx context.Context,
    actorID domainuser.ID,
) ([]domainprayergroup.PrayerGroup, error) {
    return nil, nil
}

func (r *updatePrayerGroupRepositoryStub) Save(
    ctx context.Context,
    group *domainprayergroup.PrayerGroup,
) error {
    return nil
}

func (r *updatePrayerGroupRepositoryStub) Delete(
    ctx context.Context,
    id domainprayergroup.ID,
) error {
    return nil
}

func TestUpdateReturnsFindError(t *testing.T) {
	expected := errors.New("database error")

	repo := &updatePrayerGroupRepositoryStub{
		findErr: expected,
	}

	service := NewUpdateService(repo)

	_, err := service.Update(context.Background(), UpdateRequest{
		GroupID: "group-1",
	})

	if !errors.Is(err, expected) {
		t.Fatalf("expected %v, got %v", expected, err)
	}
}

func TestUpdateReturnsNotFoundWhenGroupMissing(t *testing.T) {
	repo := &updatePrayerGroupRepositoryStub{}

	service := NewUpdateService(repo)

	_, err := service.Update(context.Background(), UpdateRequest{
		GroupID: "group-1",
	})

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "prayer group not found" {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestUpdateReturnsRenameValidationError(t *testing.T) {
	group := &domainprayergroup.PrayerGroup{}

	repo := &updatePrayerGroupRepositoryStub{
		group: group,
	}

	service := NewUpdateService(repo)

	name := ""

	_, err := service.Update(context.Background(), UpdateRequest{
		GroupID: "group-1",
		Name:    &name,
	})

	if err == nil {
		t.Fatal("expected validation error")
	}

	if repo.updateCalls != 0 {
		t.Fatal("repository should not have been updated")
	}
}

func TestUpdateReturnsRepositoryUpdateError(t *testing.T) {
    group, err := domainprayergroup.New(
        domainprayergroup.ID("group-1"),
        "Men",
        "Prayer group",
        domainprayergroup.TypeRegular,
    )
	if err != nil {
		t.Fatal(err)
	}

	expected := errors.New("update failed")

	repo := &updatePrayerGroupRepositoryStub{
		group:     group,
		updateErr: expected,
	}

	service := NewUpdateService(repo)

	description := "Updated description"

	_, err = service.Update(context.Background(), UpdateRequest{
		GroupID:     group.ID,
		Description: &description,
	})

	if !errors.Is(err, expected) {
		t.Fatalf("expected %v, got %v", expected, err)
	}
}

func TestUpdatePrayerGroup(t *testing.T) {
	group, err := domainprayergroup.New(
        domainprayergroup.ID("group-1"),
        "Men",
        "Prayer group",
        domainprayergroup.TypeRegular,
    )
	if err != nil {
		t.Fatal(err)
	}

	repo := &updatePrayerGroupRepositoryStub{
		group: group,
	}

	service := NewUpdateService(repo)

	name := "Young Adults"
	description := "Friday evenings"

	updated, err := service.Update(context.Background(), UpdateRequest{
		GroupID:     group.ID,
		Name:        &name,
		Description: &description,
	})

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if updated.Name != name {
		t.Fatalf("expected name %q, got %q", name, updated.Name)
	}

	if updated.Description != description {
		t.Fatalf("expected description %q, got %q", description, updated.Description)
	}

	if repo.updateCalls != 1 {
		t.Fatalf("expected Update to be called once")
	}

	if repo.updated != updated {
		t.Fatal("updated group was not passed to repository")
	}
}