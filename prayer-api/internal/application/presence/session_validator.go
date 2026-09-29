package presence

import (
	"context"

	"prayer-api/internal/domain/identity"
	domainuser "prayer-api/internal/domain/user"
)

type SessionAccessValidator interface {
    CanJoin(
        ctx context.Context,
        sessionID identity.PrayerSessionID,
        userID identity.UserID,
    ) error
}

type SessionValidator struct {
    sessions PrayerSessionRepository
    users    UserRepository
}

func NewSessionValidator(
    sessions PrayerSessionRepository,
    users UserRepository,
) *SessionValidator {
    return &SessionValidator{
        sessions: sessions,
        users:    users,
    }
}

func (v *SessionValidator) CanJoin(
	ctx context.Context,
	sessionID identity.PrayerSessionID,
	userID identity.UserID,
) error {
	session, err := v.sessions.FindByID(ctx, sessionID)
	if err != nil {
		return err
	}

	if session == nil {
		return ErrSessionNotFound
	}

	user, err := v.users.FindByExternalID(ctx, userID.String())
	if err != nil {
		return err
	}

	if user == nil {
		return domainuser.ErrUserNotFound
	}

	if !user.IsActive() {
		return domainuser.ErrUserBlocked
	}

	if !user.HasPrayerGroup(session.PrayerGroupID) {
		return ErrUserCannotJoinSession
	}

	return nil
}