package prayersession

import (
	"context"
	"strings"
	"prayer-api/internal/domain/identity"
)

func (s *Service) Delete(
	ctx context.Context,
	userID identity.UserID,
	sessionID identity.PrayerSessionID,
) error {

	// 1. Find the session.
	session, err := s.sessions.FindByID(ctx, sessionID)
	if err != nil {
		return err
	}

	if session == nil {
		return ErrNotFound
	}

	// 2. Find the user.
	u, err := s.users.FindByExternalID(
		ctx, 
		strings.TrimSpace(userID.String()),
	)
	
	if err != nil {
		return err
	}

	if u == nil {
		return ErrUserNotFound
	}

	// 3. Check whether the user has access
	//    to the prayer group containing this session.
	if !u.HasPrayerGroup(session.PrayerGroupID) {
		return ErrForbidden
	}

	// 4. Delete the session.
	return s.sessions.Delete(ctx, sessionID)
}