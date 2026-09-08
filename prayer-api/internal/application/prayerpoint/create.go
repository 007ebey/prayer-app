
package prayerpoint

import (
	"context"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayerpoint"
)

func (s *Service) Create(
	ctx context.Context,
	input CreateInput,
) (*domain.PrayerPoint, error) {
	if _, err := s.groups.FindByID(ctx, input.GroupID); err != nil {
		return nil, err
	}

	hasAccess, err := s.access.HasPrayerGroup(
		ctx,
		input.UserID,
		input.GroupID,
	)
	if err != nil {
		return nil, err
	}

	if !hasAccess {
		return nil, ErrAccessDenied
	}

	pointID := identity.PrayerPointID(
		"prayer-point-" + string(input.GroupID),
	)

	point, err := domain.New(
		pointID,
		input.GroupID,
		input.Title,
		input.Content,
	)
	if err != nil {
		return nil, err
	}

	if err := s.points.Save(ctx, point); err != nil {
		return nil, err
	}

	return point, nil
}