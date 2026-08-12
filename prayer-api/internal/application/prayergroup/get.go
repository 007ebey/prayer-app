package prayergroup

import (
	"context"
	"errors"
	"strings"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
)

var (
	ErrPrayerGroupNotFound = errors.New("prayer group not found")
)


type GetQuery struct {
	ActorExternalID string
	GroupID         domainprayergroup.ID
}

type GetResult struct {
	PrayerGroup domainprayergroup.PrayerGroup
}

func (s *GetService) Get(
	ctx context.Context,
	query GetQuery,
) (*GetResult, error) {
	if strings.TrimSpace(query.ActorExternalID) == "" {
		return nil, ErrUnauthorized
	}

	actor, err := s.users.FindByExternalID(
		ctx,
		query.ActorExternalID,
	)
	if err != nil {
		if errors.Is(err, ErrActorNotFound) {
			return nil, ErrActorNotFound
		}

		return nil, err
	}

	group, err := s.groups.FindByID(
		ctx,
		query.GroupID,
	)
	if err != nil {
		if errors.Is(err, ErrPrayerGroupNotFound) {
			return nil, ErrPrayerGroupNotFound
		}

		return nil, err
	}

	// Future authorization check.
	_ = actor

	return &GetResult{
		PrayerGroup: *group,
	}, nil
}