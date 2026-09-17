package prayersession

import "errors"

var (
	ErrNotFound           = errors.New("prayer session not found")
	ErrInvalidDuration    = errors.New("prayer session duration must be greater than zero")
)