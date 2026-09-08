package prayerpoint

import (
	"context"

	"prayer-api/internal/domain/prayerpoint"
)

func (s *Service) Update(
	ctx context.Context,
	input UpdateInput,
) (*prayerpoint.PrayerPoint, error) {
	point, err := s.points.FindByID(ctx, input.PointID)
	if err != nil {
		return nil, err
	}

	if point == nil {
		return nil, ErrPrayerPointNotFound
	}

	hasAccess, err := s.access.HasPrayerGroup(
		ctx,
		input.UserID,
		point.GroupID,
	)
	if err != nil {
		return nil, err
	}

	if !hasAccess {
		return nil, ErrAccessDenied
	}

	if err := point.UpdateTitle(input.Title); err != nil {
		return nil, err
	}

	if err := point.UpdateContent(input.Content); err != nil {
		return nil, err
	}

	if err := s.points.Update(ctx, point); err != nil {
		return nil, err
	}

	return point, nil
}
