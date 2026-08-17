package user

import (
	"strings"
	"prayer-api/internal/domain/identity"
)

type Status string

const (
	StatusActive  Status = "active"
	StatusBlocked Status = "blocked"
)

type User struct {
	ID              identity.UserID
	ExternalID      string
	Name            string
	Status          Status
	RoleIDs         []identity.roleID
	PrayerGroupIDs  []identity.PrayerGroupID
}

func New(
	id identity.UserID,
	externalID string,
	name string,
	defaultRoleID identity.roleID,
) (*User, error) {
	if id == "" {
		return nil, ErrIDRequired
	}

	externalID = strings.TrimSpace(externalID)
	if externalID == "" {
		return nil, ErrExternalIDRequired
	}

	name = strings.TrimSpace(name)
	if name == "" {
		return nil, ErrDisplayNameRequired
	}

	return &User{
		ID:             id,
		ExternalID:     externalID,
		Name:           name,
		Status:         StatusActive,
		RoleIDs:        []identity.roleID{defaultRoleID},
		PrayerGroupIDs: []identity.PrayerGroupID{},
	}, nil
}

func (u *User) IsActive() bool {
	return u.Status == StatusActive
}

func (u *User) Block() {
	u.Status = StatusBlocked
}

func (u *User) Unblock() {
	u.Status = StatusActive
}

func (u *User) AssignRole(roleID identity.roleID) error {
	for _, existing := range u.RoleIDs {
		if existing == roleID {
			return ErrRoleAlreadyAssigned
		}
	}

	u.RoleIDs = append(u.RoleIDs, roleID)
	return nil
}

func (u *User) RemoveRole(roleID identity.roleID) error {
	for index, existing := range u.RoleIDs {
		if existing != roleID {
			continue
		}

		u.RoleIDs = append(u.RoleIDs[:index], u.RoleIDs[index+1:]...)
		return nil
	}

	return ErrRoleNotAssigned
}

func (u *User) AssignPrayerGroup(
	groupID identity.PrayerGroupID,
) error {
	for _, existing := range u.PrayerGroupIDs {
		if existing == groupID {
			return ErrPrayerGroupAlreadyAssigned
		}
	}

	u.PrayerGroupIDs = append(
		u.PrayerGroupIDs,
		groupID,
	)

	return nil
}

func (u *User) RemovePrayerGroup(
	groupID identity.PrayerGroupID,
) error {
	for index, existing := range u.PrayerGroupIDs {
		if existing != groupID {
			continue
		}

		u.PrayerGroupIDs = append(
			u.PrayerGroupIDs[:index],
			u.PrayerGroupIDs[index+1:]...,
		)

		return nil
	}

	return ErrPrayerGroupNotAssigned
}

func (u *User) HasPrayerGroup(
	groupID identity.PrayerGroupID,
) bool {
	for _, existing := range u.PrayerGroupIDs {
		if existing == groupID {
			return true
		}
	}

	return false
}