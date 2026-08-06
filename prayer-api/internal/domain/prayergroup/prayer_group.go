package prayergroup

import "strings"

type ID string

type SessionID string

type Type string

type Status string

const (
	TypeVisitor Type = "visitor"
	TypeRegular Type = "regular"
)

const (
	StatusActive   Status = "active"
	StatusInactive Status = "inactive"
)

type PrayerGroup struct {
	ID          ID
	Name        string
	Description string
	Type        Type
	Status      Status
	SessionIDs  []SessionID
}

func New(
	id ID,
	name string,
	description string,
	groupType Type,
) (*PrayerGroup, error) {
	if id == "" {
		return nil, ErrIDRequired
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrNameRequired
	}

	return &PrayerGroup{
		ID:          id,
		Name:        name,
		Description: strings.TrimSpace(description),
		Type:        groupType,
		Status:      StatusActive,
		SessionIDs:  []SessionID{},
	}, nil
}

func (g *PrayerGroup) IsActive() bool {
	return g.Status == StatusActive
}

func (g *PrayerGroup) IsVisitor() bool {
	return g.Type == TypeVisitor
}

func (g *PrayerGroup) Rename(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return ErrNameRequired
	}

	g.Name = name
	return nil
}

func (g *PrayerGroup) ChangeDescription(description string) {
	g.Description = strings.TrimSpace(description)
}

func (g *PrayerGroup) AssignSession(sessionID SessionID) error {
	for _, existing := range g.SessionIDs {
		if existing == sessionID {
			return ErrSessionAlreadyAdded
		}
	}

	g.SessionIDs = append(g.SessionIDs, sessionID)
	return nil
}

func (g *PrayerGroup) RemoveSession(sessionID SessionID) error {
	for index, existing := range g.SessionIDs {
		if existing != sessionID {
			continue
		}

		g.SessionIDs = append(g.SessionIDs[:index], g.SessionIDs[index+1:]...)
		return nil
	}

	return ErrSessionNotAssigned
}

func (g *PrayerGroup) ReplaceSessions(sessionIDs []SessionID) {
	seen := make(map[SessionID]struct{})
	result := make([]SessionID, 0, len(sessionIDs))

	for _, sessionID := range sessionIDs {
		if _, exists := seen[sessionID]; exists {
			continue
		}

		seen[sessionID] = struct{}{}
		result = append(result, sessionID)
	}

	g.SessionIDs = result
}

func (g *PrayerGroup) Activate() {
	g.Status = StatusActive
}

func (g *PrayerGroup) Deactivate() {
	g.Status = StatusInactive
}
