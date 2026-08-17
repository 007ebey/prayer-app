package role

import (
  "strings"
  "prayer-api/internal/domain/identity"
)

type Role struct {
	ID          identity.RoleID
	Name        string
	Description string
	Permissions []Permission
	IsSystem    bool
}

func New(
	id identity.RoleID,
	name string,
	description string,
	permissions []Permission,
	isSystem bool,
) (*Role, error) {
	name = strings.TrimSpace(name)

	if name == "" {
		return nil, ErrNameRequired
	}

	if len(permissions) == 0 {
		return nil, ErrPermissionRequired
	}

	if err := validatePermissions(permissions); err != nil {
		return nil, err
	}

	return &Role{
		ID:          id,
		Name:        name,
		Description: strings.TrimSpace(description),
		Permissions: uniquePermissions(permissions),
		IsSystem:    isSystem,
	}, nil
}

func (r *Role) HasPermission(permission Permission) bool {
	for _, current := range r.Permissions {
		if current == permission {
			return true
		}
	}

	return false
}

func (r *Role) ReplacePermissions(permissions []Permission) error {
	if len(permissions) == 0 {
		return ErrPermissionRequired
	}

	if err := validatePermissions(permissions); err != nil {
		return err
	}

	r.Permissions = uniquePermissions(permissions)
	return nil
}

func (r *Role) Rename(name string) error {
	name = strings.TrimSpace(name)

	if name == "" {
		return ErrNameRequired
	}

	r.Name = name
	return nil
}

func (r *Role) ChangeDescription(description string) {
	r.Description = strings.TrimSpace(description)
}

func validatePermissions(permissions []Permission) error {
	for _, permission := range permissions {
		if !IsValidPermission(permission) {
			return ErrInvalidPermission
		}
	}

	return nil
}

func uniquePermissions(permissions []Permission) []Permission {
	seen := make(map[Permission]struct{})
	result := make([]Permission, 0, len(permissions))

	for _, permission := range permissions {
		if _, exists := seen[permission]; exists {
			continue
		}

		seen[permission] = struct{}{}
		result = append(result, permission)
	}

	return result
}
