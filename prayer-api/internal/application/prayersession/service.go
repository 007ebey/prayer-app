package prayersession

type Service struct {
	sessions PrayerSessionRepository
	users    UserRepository
}

func NewService(
	sessions PrayerSessionRepository,
	users UserRepository,
) *Service {
	return &Service{
		sessions: sessions,
		users:    users,
	}
}

type CreateCommand struct {
	ActorID        identity.UserID
	PrayerGroupID  identity.PrayerGroupID
	Title          string
	Date           time.Time
	Time           string
	Duration       int
	PrayerPointIDs []identity.PrayerPointID
}

type UpdateCommand struct {
	ActorID        identity.UserID
	SessionID      identity.PrayerSessionID
	Title          string
	Date           time.Time
	Time           string
	Duration       int
	PrayerPointIDs []identity.PrayerPointID
}
