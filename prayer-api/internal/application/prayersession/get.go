package prayersession

import (
	"context"
	"strings"
)

func (s *Service) GetByID(
	ctx context.Context,
	actorID identity.UserID,
	sessionID identity.PrayerSessionID,
) (*PrayerSession, error) {
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

	if err := s.requireGroupAccess(
		ctx,
		actorID,
		session.PrayerGroupID,
	); err != nil {
		return nil, err
	}

	return session, nil
}