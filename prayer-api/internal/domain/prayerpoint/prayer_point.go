package prayerpoint

import (
	"strings"
	"time"

	"prayer-api/internal/domain/identity"
)

type Status string

const (
	StatusActive  Status = "active"
	StatusBlocked Status = "blocked"
)

type PrayerPoint struct {
	ID        identity.PrayerPointID
	GroupID   identity.PrayerGroupID
	Title     string
	Content   string
	Status    Status
	CreatedAt time.Time
	UpdatedAt time.Time
}

func New(
	id identity.PrayerPointID,
	groupID identity.PrayerGroupID,
	title string,
	content string,
) (*PrayerPoint, error) {
	if id == "" {
		return nil, ErrIDRequired
	}

	if groupID == "" {
		return nil, ErrGroupIDRequired
	}

	title = strings.TrimSpace(title)
	if title == "" {
		return nil, ErrTitleRequired
	}

	content = strings.TrimSpace(content)
	if content == "" {
		return nil, ErrContentRequired
	}

	now := time.Now().UTC()

	return &PrayerPoint{
		ID:        id,
		GroupID:   groupID,
		Title:     title,
		Content:   content,
		Status:    StatusActive,
		CreatedAt: now,
		UpdatedAt: now,
	}, nil
}

func (p *PrayerPoint) IsActive() bool {
	return p.Status == StatusActive
}

func (p *PrayerPoint) Block() {
	p.Status = StatusBlocked
	p.UpdatedAt = time.Now().UTC()
}

func (p *PrayerPoint) Unblock() {
	p.Status = StatusActive
	p.UpdatedAt = time.Now().UTC()
}

func (p *PrayerPoint) UpdateTitle(title string) error {
	title = strings.TrimSpace(title)

	if title == "" {
		return ErrTitleRequired
	}

	p.Title = title
	p.UpdatedAt = time.Now().UTC()

	return nil
}

func (p *PrayerPoint) UpdateContent(content string) error {
	content = strings.TrimSpace(content)

	if content == "" {
		return ErrContentRequired
	}

	p.Content = content
	p.UpdatedAt = time.Now().UTC()

	return nil
}

func (p *PrayerPoint) BelongsToGroup(
	groupID identity.PrayerGroupID,
) bool {
	return p.GroupID == groupID
}