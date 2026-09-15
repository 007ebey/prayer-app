package prayersession

import (
	"errors"
	"strings"
	"time"
	"prayer-api/internal/domain/identity"
)

var (
	ErrInvalidID            = errors.New("prayer session id is required")
	ErrInvalidPrayerGroupID = errors.New("prayer group id is required")
	ErrInvalidTitle         = errors.New("prayer session title is required")
	ErrInvalidDate          = errors.New("prayer session date is required")
)

type PrayerSession struct {
	ID             identity.PrayerSessionID
	PrayerGroupID  identity.PrayerGroupID
	Title          string
	Date           time.Time
	Time           string
	Duration       int // duration in minutes
	PrayerPointIDs []identity.PrayerPointID
}

func New(
	id identity.PrayerSessionID,
	prayerGroupID identity.PrayerGroupID,
	title string,
	date time.Time,
	sessionTime string,
	prayerPointIDs []identity.PrayerPointID,
) (*PrayerSession, error) {
	if strings.TrimSpace(id.String()) == "" {
		return nil, ErrInvalidID
	}

	if strings.TrimSpace(prayerGroupID.String()) == "" {
		return nil, ErrInvalidPrayerGroupID
	}

	if strings.TrimSpace(title) == "" {
		return nil, ErrInvalidTitle
	}

	if date.IsZero() {
		return nil, ErrInvalidDate
	}

	return &PrayerSession{
		ID:             id,
		PrayerGroupID:  prayerGroupID,
		Title:          strings.TrimSpace(title),
		Date:           date,
		Time:           strings.TrimSpace(sessionTime),
		Duration:       0, // Initialize duration to 0
		PrayerPointIDs: append([]identity.PrayerPointID(nil), prayerPointIDs...),
	}, nil
}