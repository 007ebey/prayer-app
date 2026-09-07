package prayerpoint

import (
	"errors"
	"testing"
	"time"

	"prayer-api/internal/domain/identity"
)

func TestNewPrayerPoint(t *testing.T) {
	t.Run("creates a prayer point successfully", func(t *testing.T) {
		id := identity.PrayerPointID("prayer-point-123")
		groupID := identity.PrayerGroupID("group-123")

		before := time.Now().UTC()

		point, err := New(
			id,
			groupID,
			"  Pray for healing  ",
			"  Please pray for complete healing.  ",
		)

		after := time.Now().UTC()

		if err != nil {
			t.Fatalf("expected no error, got %v", err)
		}

		if point == nil {
			t.Fatal("expected prayer point, got nil")
		}

		if point.ID != id {
			t.Errorf("expected ID %q, got %q", id, point.ID)
		}

		if point.GroupID != groupID {
			t.Errorf(
				"expected group ID %q, got %q",
				groupID,
				point.GroupID,
			)
		}

		if point.Title != "Pray for healing" {
			t.Errorf(
				"expected trimmed title %q, got %q",
				"Pray for healing",
				point.Title,
			)
		}

		if point.Content != "Please pray for complete healing." {
			t.Errorf(
				"expected trimmed content %q, got %q",
				"Please pray for complete healing.",
				point.Content,
			)
		}

		if point.Status != StatusActive {
			t.Errorf(
				"expected status %q, got %q",
				StatusActive,
				point.Status,
			)
		}

		if point.CreatedAt.Before(before) ||
			point.CreatedAt.After(after) {
			t.Errorf(
				"CreatedAt %v is outside expected range",
				point.CreatedAt,
			)
		}

		if point.UpdatedAt.Before(before) ||
			point.UpdatedAt.After(after) {
			t.Errorf(
				"UpdatedAt %v is outside expected range",
				point.UpdatedAt,
			)
		}
	})

	t.Run("requires an ID", func(t *testing.T) {
		_, err := New(
			"",
			identity.PrayerGroupID("group-123"),
			"Pray for healing",
			"Please pray for healing.",
		)

		if !errors.Is(err, ErrIDRequired) {
			t.Fatalf(
				"expected ErrIDRequired, got %v",
				err,
			)
		}
	})

	t.Run("requires a group ID", func(t *testing.T) {
		_, err := New(
			identity.PrayerPointID("prayer-point-123"),
			"",
			"Pray for healing",
			"Please pray for healing.",
		)

		if !errors.Is(err, ErrGroupIDRequired) {
			t.Fatalf(
				"expected ErrGroupIDRequired, got %v",
				err,
			)
		}
	})

	t.Run("requires a title", func(t *testing.T) {
		_, err := New(
			identity.PrayerPointID("prayer-point-123"),
			identity.PrayerGroupID("group-123"),
			"   ",
			"Please pray for healing.",
		)

		if !errors.Is(err, ErrTitleRequired) {
			t.Fatalf(
				"expected ErrTitleRequired, got %v",
				err,
			)
		}
	})

	t.Run("requires content", func(t *testing.T) {
		_, err := New(
			identity.PrayerPointID("prayer-point-123"),
			identity.PrayerGroupID("group-123"),
			"Pray for healing",
			"   ",
		)

		if !errors.Is(err, ErrContentRequired) {
			t.Fatalf(
				"expected ErrContentRequired, got %v",
				err,
			)
		}
	})
}

func TestPrayerPointIsActive(t *testing.T) {
	point := &PrayerPoint{
		Status: StatusActive,
	}

	if !point.IsActive() {
		t.Fatal("expected prayer point to be active")
	}

	point.Status = StatusBlocked

	if point.IsActive() {
		t.Fatal("expected blocked prayer point to be inactive")
	}
}

func TestPrayerPointBlock(t *testing.T) {
	point := &PrayerPoint{
		Status:    StatusActive,
		UpdatedAt: time.Now().UTC().Add(-time.Minute),
	}

	previousUpdatedAt := point.UpdatedAt

	point.Block()

	if point.Status != StatusBlocked {
		t.Fatalf(
			"expected status %q, got %q",
			StatusBlocked,
			point.Status,
		)
	}

	if !point.UpdatedAt.After(previousUpdatedAt) {
		t.Fatal("expected UpdatedAt to change after blocking")
	}

	if point.IsActive() {
		t.Fatal("expected blocked prayer point to be inactive")
	}
}

func TestPrayerPointUnblock(t *testing.T) {
	point := &PrayerPoint{
		Status:    StatusBlocked,
		UpdatedAt: time.Now().UTC().Add(-time.Minute),
	}

	previousUpdatedAt := point.UpdatedAt

	point.Unblock()

	if point.Status != StatusActive {
		t.Fatalf(
			"expected status %q, got %q",
			StatusActive,
			point.Status,
		)
	}

	if !point.UpdatedAt.After(previousUpdatedAt) {
		t.Fatal("expected UpdatedAt to change after unblocking")
	}

	if !point.IsActive() {
		t.Fatal("expected unblocked prayer point to be active")
	}
}

func TestPrayerPointUpdateTitle(t *testing.T) {
	point := &PrayerPoint{
		Title:     "Old title",
		UpdatedAt: time.Now().UTC().Add(-time.Minute),
	}

	previousUpdatedAt := point.UpdatedAt

	err := point.UpdateTitle("  New title  ")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if point.Title != "New title" {
		t.Errorf(
			"expected title %q, got %q",
			"New title",
			point.Title,
		)
	}

	if !point.UpdatedAt.After(previousUpdatedAt) {
		t.Fatal("expected UpdatedAt to change after title update")
	}
}

func TestPrayerPointUpdateTitleRequiresTitle(t *testing.T) {
	point := &PrayerPoint{
		Title: "Existing title",
	}

	err := point.UpdateTitle("   ")

	if !errors.Is(err, ErrTitleRequired) {
		t.Fatalf(
			"expected ErrTitleRequired, got %v",
			err,
		)
	}

	if point.Title != "Existing title" {
		t.Errorf(
			"expected title to remain unchanged, got %q",
			point.Title,
		)
	}
}

func TestPrayerPointUpdateContent(t *testing.T) {
	point := &PrayerPoint{
		Content:   "Old content",
		UpdatedAt: time.Now().UTC().Add(-time.Minute),
	}

	previousUpdatedAt := point.UpdatedAt

	err := point.UpdateContent("  New content  ")

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if point.Content != "New content" {
		t.Errorf(
			"expected content %q, got %q",
			"New content",
			point.Content,
		)
	}

	if !point.UpdatedAt.After(previousUpdatedAt) {
		t.Fatal("expected UpdatedAt to change after content update")
	}
}

func TestPrayerPointUpdateContentRequiresContent(t *testing.T) {
	point := &PrayerPoint{
		Content: "Existing content",
	}

	err := point.UpdateContent("   ")

	if !errors.Is(err, ErrContentRequired) {
		t.Fatalf(
			"expected ErrContentRequired, got %v",
			err,
		)
	}

	if point.Content != "Existing content" {
		t.Errorf(
			"expected content to remain unchanged, got %q",
			point.Content,
		)
	}
}

func TestPrayerPointBelongsToGroup(t *testing.T) {
	point := &PrayerPoint{
		GroupID: identity.PrayerGroupID("group-123"),
	}

	if !point.BelongsToGroup(
		identity.PrayerGroupID("group-123"),
	) {
		t.Fatal("expected prayer point to belong to group-123")
	}

	if point.BelongsToGroup(
		identity.PrayerGroupID("group-456"),
	) {
		t.Fatal("expected prayer point not to belong to group-456")
	}
}
