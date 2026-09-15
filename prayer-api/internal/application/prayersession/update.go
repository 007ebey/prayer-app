package prayersession

import (
	"context"
	"strings"

	domain "prayer-api/internal/domain/prayersession"
)

func (s *Service) Update(
	ctx context.Context,
	cmd UpdateCommand,
) (*domain.PrayerSession, error) {

	if cmd.ActorID == "" {
		return nil, ErrUnauthorized
	}

	if cmd.SessionID == "" {
		return nil, ErrInvalidID
	}

	session, err := s.sessions.FindByID(
		ctx,
		cmd.SessionID,
	)
	if err != nil {
		return nil, err
	}

	if session == nil {
		return nil, ErrNotFound
	}

	// A user can only update a session belonging
	// to a prayer group they have access to.
	if err := s.requireGroupAccess(
		ctx,
		cmd.ActorID,
		session.PrayerGroupID,
	); err != nil {
		return nil, err
	}

	// Validate title.
	title := strings.TrimSpace(cmd.Title)

	if title == "" {
		return nil, ErrTitleRequired
	}

	// Validate time.
	sessionTime := strings.TrimSpace(cmd.Time)

	if sessionTime == "" {
		return nil, ErrTimeRequired
	}

	// Validate duration.
	if cmd.Duration <= 0 {
		return nil, ErrInvalidDuration
	}

	// Update the existing session.
	session.Title = title
	session.Date = cmd.Date
	session.Time = sessionTime
	session.Duration = cmd.Duration

	// Copy the slice so the domain object does not share
	// the caller's underlying array.
	session.PrayerPointIDs = append(
		[]identity.PrayerPointID(nil),
		cmd.PrayerPointIDs...,
	)

	if err := s.sessions.Update(
		ctx,
		session,
	); err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return nil, ErrNotFound
		}

		return nil, err
	}

	return session, nil
}