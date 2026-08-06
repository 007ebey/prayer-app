package user

import (
	"strings"

	"prayer-api/internal/domain/role"
)

type ID string

type Status string

const (
	StatusActive  Status = "active"
	StatusBlocked Status = "blocked"
)

type User struct {
	ID         ID
	ExternalID string
	Name       string
	Status     Status
	RoleIDs    []role.ID
}

func New(
	id ID,
	externalID string,
	name string,
	defaultRoleID role.ID,
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
		ID:         id,
		ExternalID: externalID,
		Name:       name,
		Status:     StatusActive,
		RoleIDs:    []role.ID{defaultRoleID},
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

func (u *User) AssignRole(roleID role.ID) error {
	for _, existing := range u.RoleIDs {
		if existing == roleID {
			return ErrRoleAlreadyAssigned
		}
	}

	u.RoleIDs = append(u.RoleIDs, roleID)
	return nil
}

func (u *User) RemoveRole(roleID role.ID) error {
	for index, existing := range u.RoleIDs {
		if existing != roleID {
			continue
		}

		u.RoleIDs = append(u.RoleIDs[:index], u.RoleIDs[index+1:]...)
		return nil
	}

	return ErrRoleNotAssigned
}
