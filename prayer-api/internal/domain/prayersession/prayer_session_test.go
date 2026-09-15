package prayersession

import (
	"testing"
	"time"

	"prayer-api/internal/domain/identity"
)

func TestNewPrayerSession(t *testing.T) {
	date := time.Date(2026, 8, 3, 0, 0, 0, 0, time.UTC)

	session, err := New(
		identity.PrayerSessionID("session-1"),
		identity.PrayerGroupID("group-1"),
		"Morning Prayer",
		date,
		"06:00",
		[]identity.PrayerPointID{identity.PrayerPointID("prayer-1"), identity.PrayerPointID("prayer-2")},
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if session.ID != identity.PrayerSessionID("session-1") {
		t.Errorf("unexpected ID: %s", session.ID)
	}

	if session.PrayerGroupID != identity.PrayerGroupID("group-1") {
		t.Errorf("unexpected group ID: %s", session.PrayerGroupID)
	}

	if session.Title != "Morning Prayer" {
		t.Errorf("unexpected title: %s", session.Title)
	}

	if session.Time != "06:00" {
		t.Errorf("unexpected time: %s", session.Time)
	}

	if len(session.PrayerPointIDs) != 2 {
		t.Errorf("expected 2 prayer points, got %d", len(session.PrayerPointIDs))
	}
}