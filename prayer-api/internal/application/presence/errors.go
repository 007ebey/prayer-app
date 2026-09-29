package presence

import (
	"errors"
)

var (
    ErrUserCannotJoinSession = errors.New("user cannot join the prayer session")
	ErrSessionNotFound = errors.New("prayer session not found")
)