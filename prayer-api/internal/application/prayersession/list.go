package prayersession

import
(
  domainuser "prayer-api/internal/domain/user"
  domain "prayer-api/internal/domain/prayersession"
  "prayer-api/internal/domain/prayerpoint"
  "context"
  "prayer-api/internal/domain/identity"
)

type PrayerSessionWithPoints struct {
	Session domain.PrayerSession
	Points  []prayerpoint.PrayerPoint
}

func (s *Service) ListForUser(
	ctx context.Context,
	userID identity.UserID,
) ([]PrayerSessionWithPoints, error) {

	// 1. Find the user.
	u, err := s.users.FindByExternalID(ctx, userID.String())
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, domainuser.ErrUserNotFound
	}

	var result []PrayerSessionWithPoints

	// 2. Get sessions for each group the user has access to.
	for _, groupID := range u.PrayerGroupIDs {
		sessions, err := s.sessions.ListByGroupID(ctx, groupID)
		if err != nil {
			return nil, err
		}

		for _, session := range sessions {
			var points []prayerpoint.PrayerPoint

			// 3. Resolve the prayer point IDs.
			for _, pointID := range session.PrayerPointIDs {
				point, err := s.prayerPoints.FindByID(ctx, pointID)
				if err != nil {
					return nil, err
				}

				if point != nil {
					points = append(points, *point)
				}
			}

			// 4. Combine the session with its resolved points.
			result = append(result, PrayerSessionWithPoints{
				Session: session,
				Points:  points,
			})
		}
	}

	return result, nil
}