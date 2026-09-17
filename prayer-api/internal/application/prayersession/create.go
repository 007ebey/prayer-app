package prayersession

import (
	"context"
	"strings"
	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
)

func (s *Service) Create(
	ctx context.Context,
	cmd CreateCommand,
) (*domain.PrayerSession, error) {
	if cmd.PrayerGroupID.String() == "" {
		return nil, ErrInvalidGroupID
	}

	if strings.TrimSpace(cmd.Title) == "" {
		return nil, ErrTitleRequired
	}

	if cmd.Date.IsZero() {
		return nil, ErrInvalidDate
	}

	if strings.TrimSpace(cmd.Time) == "" {
		return nil, ErrTimeRequired
	}

	if cmd.Duration <= 0 {
		return nil, ErrInvalidDuration
	}

	u, err := s.users.FindByExternalID(ctx, strings.TrimSpace(cmd.ActorID.String()))
	if err != nil {
		return nil, ErrUnknown
	}
	if u == nil {
		return nil, ErrUserNotFound
	}

	if !u.HasPrayerGroup(
		cmd.PrayerGroupID,
	) {
		return nil, ErrForbidden
	}

	session, err := domain.New(
		identity.PrayerSessionID("session-generated"),
		cmd.PrayerGroupID,
		strings.TrimSpace(cmd.Title),
		cmd.Date,
		strings.TrimSpace(cmd.Time),
		cmd.Duration,
		cmd.PrayerPointIDs,
	)
	
	if err != nil {
		return nil, ErrUnknown
	}

	if err := s.sessions.Save(ctx, session); err != nil {
		return nil, err
	}

	return session, nil
}


