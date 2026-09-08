package prayerpoint

import (
	"context"

	"prayer-api/internal/domain/prayerpoint"
)

func (s *Service) ListByGroupID(
	ctx context.Context,
	input ListInput,
) ([]prayerpoint.PrayerPoint, error) {
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

	return s.points.ListByGroupID(ctx, input.GroupID)
}
