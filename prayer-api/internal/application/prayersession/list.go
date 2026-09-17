package prayersession

import
(
  domainuser "prayer-api/internal/domain/user"
  domain "prayer-api/internal/domain/prayersession"
  "context"
  "prayer-api/internal/domain/identity"
)

func (s *Service) ListForUser(
	ctx context.Context,
	userID identity.UserID,
) ([]domain.PrayerSession, error) {

	// 1. Find the user.
	u, err := s.users.FindByExternalID(ctx, userID.String())
	if err != nil {
		return nil, err
	}

	if u == nil {
		return nil, domainuser.ErrUserNotFound
	}

	// 2. Get the groups the user has access to.
	var allSessions []domain.PrayerSession

	// 3. Get sessions for each group.
	for _, groupID := range u.PrayerGroupIDs {
		sessions, err := s.sessions.ListByGroupID(ctx, groupID)
		if err != nil {
			return nil, err
		}

		// 4. Add those sessions to the user's complete session list.
		allSessions = append(allSessions, sessions...)
	}

	// 5. Return all sessions across all accessible groups.
	return allSessions, nil
}