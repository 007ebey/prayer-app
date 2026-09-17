package prayersession

import (
	"context"
	"strings"
	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
)

func (s *Service) GetByID(
	ctx context.Context,
	actorID identity.UserID,
	sessionID identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	if strings.TrimSpace(sessionID.String()) == "" {
		return nil, ErrInvalidID
	}

	session, err := s.sessions.FindByID(ctx, sessionID)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, ErrNotFound
	}
    
	u, err := s.users.FindByExternalID(ctx, strings.TrimSpace(actorID.String())); if err != nil {
		return nil, ErrUserNotFound
	}

	if !u.HasPrayerGroup(session.PrayerGroupID) {
		return nil, ErrForbidden
	}

	return session, nil
}