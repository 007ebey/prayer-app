package prayergroup_test

import (
	"testing"

	"prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/prayergroup"
)

func TestVisitorGroupIsActive(t *testing.T) {
	group, err := prayergroup.New(
		identity.PrayerGroupID("visitor"),
		"Visitor",
		"Default prayer group access.",
		prayergroup.TypeVisitor,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !group.IsActive() {
		t.Fatal("expected visitor group to be active")
	}

	if !group.IsVisitor() {
		t.Fatal("expected group to be visitor group")
	}
}

func TestPrayerGroupAccessCanBeBlocked(t *testing.T) {
	access := prayergroup.NewAccess(
		identity.UserID("user_1"),
		identity.PrayerGroupID("youth"),
	)

	if !access.IsActive() {
		t.Fatal("expected new access to be active")
	}

	access.Block()

	if access.IsActive() {
		t.Fatal("expected blocked access to be inactive")
	}

	if access.Status != prayergroup.AccessBlocked {
		t.Fatalf("expected blocked status, got %s", access.Status)
	}

	access.Unblock()

	if !access.IsActive() {
		t.Fatal("expected access to be active after unblock")
	}
}

func TestPrayerGroupReplaceSessionsRemovesDuplicates(t *testing.T) {
	group, err := prayergroup.New(
		identity.PrayerGroupID("youth"),
		"Youth Prayer",
		"Prayer sessions for youth.",
		prayergroup.TypeRegular,
	)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	group.ReplaceSessions([]identity.SessionID{
		"morning",
		"evening",
		"morning",
	})

	if len(group.SessionIDs) != 2 {
		t.Fatalf("expected 2 unique sessions, got %d", len(group.SessionIDs))
	}
}
