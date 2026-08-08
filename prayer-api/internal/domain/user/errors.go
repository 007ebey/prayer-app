package user

import "errors"

var (
	ErrIDRequired          = errors.New("user id is required")
	ErrExternalIDRequired  = errors.New("external identity id is required")
	ErrDisplayNameRequired = errors.New("display name is required")
	ErrUserBlocked         = errors.New("user is blocked")
	ErrRoleAlreadyAssigned = errors.New("role already assigned to user")
	ErrRoleNotAssigned     = errors.New("role is not assigned to user")
)
