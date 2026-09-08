package prayerpoint

import ( 
  "context" 
  "prayer-api/internal/domain/identity" 
   domain "prayer-api/internal/domain/prayerpoint"
   "prayer-api/internal/domain/prayergroup"
)

type PrayerPointRepository interface {

  FindByID(
	ctx context.Context,
	id identity.PrayerPointID,
  ) (*domain.PrayerPoint, error)

  ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
  ) ([]domain.PrayerPoint, error)

  Save(
	ctx context.Context,
	point *domain.PrayerPoint,
  ) error

  Update(
	ctx context.Context,
	point *domain.PrayerPoint,
  ) error

  Delete(
	ctx context.Context,
	id identity.PrayerPointID,
  ) error
}

type PrayerGroupReader interface {
  FindByID(
	ctx context.Context,
	id identity.PrayerGroupID,
  ) (*prayergroup.PrayerGroup, error)
}

type UserAccessReader interface { 
  HasPrayerGroup( 
	ctx context.Context, 
	userID identity.UserID, 
	groupID identity.PrayerGroupID, 
  ) (bool, error)
}