package prayergroup

import (
	"context"
	"errors"
	"testing"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
)

type prayerGroupRepositoryStub struct {
	groups []domainprayergroup.PrayerGroup
	err    error
}

func (r *prayerGroupRepositoryStub) List(
	ctx context.Context,
) ([]domainprayergroup.PrayerGroup, error) {
	return r.groups, r.err
}

func TestListReturnsPrayerGroups(t *testing.T) {
	repo := &prayerGroupRepositoryStub{
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

	service := appprayergroup.NewListService(repo)

	result, err := service.List(context.Background())

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(result.PrayerGroups) != 2 {
		t.Fatalf("expected 2 prayer groups, got %d",
			len(result.PrayerGroups))
	}
}

func TestListReturnsEmptyList(t *testing.T) {
	repo := &prayerGroupRepositoryStub{}

	service := appprayergroup.NewListService(repo)

	result, err := service.List(context.Background())

	if err != nil {
		t.Fatalf("expected nil error, got %v", err)
	}

	if len(result.PrayerGroups) != 0 {
		t.Fatalf("expected empty list")
	}
}

func TestListReturnsRepositoryError(t *testing.T) {
	expected := errors.New("database unavailable")

	repo := &prayerGroupRepositoryStub{
		err: expected,
	}

	service := appprayergroup.NewListService(repo)

	_, err := service.List(context.Background())

	if !errors.Is(err, expected) {
		t.Fatalf(
			"expected %v, got %v",
			expected,
			err,
		)
	}
}