package prayerpoint

import "errors"

var (
	ErrIDRequired           = errors.New("prayer group id is required")
	ErrGroupIDRequired      = errors.New("prayer group id is required")
	ErrTitleRequired        = errors.New("prayer group title is required")
	ErrContentRequired      = errors.New("prayer group content is required")
	ErrNotFound             = errors.New("prayer point not found")	
)