package prayersession

import (
	"errors"
)

var (
	ErrNotFound = errors.New("prayer session not found")
	ErrUserNotFound = errors.New("user not found")
	ErrForbidden = errors.New("access forbidden")
	ErrInvalidID = errors.New("invalid prayer session ID")
	ErrTitleRequired = errors.New("title is required")
	ErrInvalidGroupID = errors.New("invalid prayer group ID")
	ErrInvalidDate = errors.New("invalid date")
	ErrTimeRequired = errors.New("time is required")
	ErrInvalidDuration = errors.New("invalid duration")
	ErrUnknown = errors.New("unknown error occurred")
	ErrUnauthorized = errors.New("unauthorized")
)