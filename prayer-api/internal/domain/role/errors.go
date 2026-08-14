package role

import "errors"

var (
	ErrNameRequired       = errors.New("role name is required")
	ErrInvalidPermission  = errors.New("invalid permission")
	ErrPermissionRequired = errors.New("at least one permission is required")
	ErrNotFound 		  = errors.New("role not found")
)
