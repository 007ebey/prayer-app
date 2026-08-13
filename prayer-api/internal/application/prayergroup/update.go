package prayergroup

import (
	"context"
	"errors"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
)

type UpdateRequest struct {
    GroupID     domainprayergroup.ID
    Name        *string
    Description *string
}

type UpdateService struct {
    groups PrayerGroupRepository
}

func NewUpdateService(
	groups PrayerGroupRepository,
) *UpdateService {
	return &UpdateService{
		groups: groups,
	}
}


func (s *UpdateService) Update(
	ctx context.Context,
	req UpdateRequest,
) (*domainprayergroup.PrayerGroup, error) {

	group, err := s.groups.FindByID(ctx, req.GroupID)
	if err != nil {
		return nil, err
	}

	if group == nil {
		return nil, errors.New("prayer group not found")
	}

	if req.Name != nil {
		if err := group.Rename(*req.Name); err != nil {
			return nil, err
		}
	}

	if req.Description != nil {
		group.ChangeDescription(*req.Description)
	}

	if err := s.groups.Update(ctx, group); err != nil {
		return nil, err
	}

	return group, nil
}