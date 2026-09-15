import (
	"errors"
)

var (
	ErrNotFound = errors.New("prayer session not found")
	ErrUserNotFound = errors.New("user not found")
	ErrForbidden = errors.New("access forbidden")
)