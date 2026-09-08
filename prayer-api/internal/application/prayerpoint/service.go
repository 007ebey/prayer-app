package prayerpoint

import ( 
  "context" 
  "errors" 
  "prayer-api/internal/domain/identity" 
  domain "prayer-api/internal/domain/prayerpoint" 
)

var ( 
	ErrPrayerGroupNotFound = errors.New("prayer group not found")
	ErrPrayerPointNotFound = errors.New("prayer point not found")
	ErrAccessDenied = errors.New("access denied") 
)

type Service struct { 
	points PrayerPointRepository 
	groups PrayerGroupReader 
	access UserAccessReader 
}

func NewService( 
	points PrayerPointRepository,
	groups PrayerGroupReader,
	access UserAccessReader,
 ) *Service { 
	return &Service{ 
		points: points, 
		groups: groups, 
		access: access,
    }
}

type CreateInput struct { 
	UserID identity.UserID 
	GroupID identity.PrayerGroupID 
	Title string 
	Content string 
}

type UpdateInput struct { 
	UserID identity.UserID 
	PointID identity.PrayerPointID 
	Title string 
	Content string 
}

type ListInput struct { 
	UserID identity.UserID 
	GroupID identity.PrayerGroupID 
}