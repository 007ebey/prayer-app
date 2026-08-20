package prayergroup

import "errors"

var (
	ErrIDRequired           = errors.New("prayer group id is required")
	ErrNameRequired         = errors.New("prayer group name is required")
	ErrAccessBlocked        = errors.New("prayer group access is blocked")
	ErrAccessAlreadyGranted = errors.New("prayer group access already granted")
	ErrSessionAlreadyAdded  = errors.New("prayer session already assigned")
	ErrSessionNotAssigned   = errors.New("prayer session is not assigned")
	ErrNotFound  		    = errors.New("prayer group not found")
	ErrPrayerGroupNotFound  = errors.New("prayer group not found")
	ErrPrayerGroupAlreadyAssigned = errors.New("prayer group already assigned")
)
