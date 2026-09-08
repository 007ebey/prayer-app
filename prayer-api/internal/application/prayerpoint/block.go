package prayerpoint

import (
	"context"

	"prayer-api/internal/domain/identity"
)

func (s *Service) Block(
	ctx context.Context,
	userID identity.UserID,
	pointID identity.PrayerPointID,
) error {
	point, err := s.points.FindByID(ctx, pointID)
	if err != nil {
		return err
	}

	if point == nil {
		return ErrPrayerPointNotFound
	}

	hasAccess, err := s.access.HasPrayerGroup(
		ctx,
		userID,
		point.GroupID,
	)
	if err != nil {
		return err
	}

	if !hasAccess {
		return ErrAccessDenied
	}

	point.Block()

	return s.points.Update(ctx, point)
}

