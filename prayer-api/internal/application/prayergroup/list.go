package prayergroup

import (
	"context"
    "errors"
	"strings"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
)

type ListQuery struct {
	ActorExternalID string
}

type ListService struct {
	groups PrayerGroupRepository
	users  UserRepository
}

type ListResult struct {
	PrayerGroups []domainprayergroup.PrayerGroup
}

func (s *ListService) List(
	ctx context.Context,
	query ListQuery,
) (*ListResult, error) {
	if strings.TrimSpace(query.ActorExternalID) == "" {
		return nil, ErrUnauthorized
	}

	actor, err := s.users.FindByExternalID(ctx, query.ActorExternalID)
	if err != nil {
		if errors.Is(err, ErrActorNotFound) {
			return nil, ErrActorNotFound
		}
		return nil, err
	}

	groups, err := s.groups.List(ctx, actor.ID)
	if err != nil {
		return nil, err
	}

	return &ListResult{
		PrayerGroups: groups,
	}, nil
}

func NewListService(
	users UserRepository,
	groups PrayerGroupRepository,
) *ListService {
	return &ListService{
		users:  users,
		groups: groups,
	}
}