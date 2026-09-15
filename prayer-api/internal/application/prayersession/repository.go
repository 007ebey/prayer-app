package prayersession

type PrayerSessionRepository interface {
	FindByID(
		ctx context.Context,
		id identity.PrayerSessionID,
	) (*PrayerSession, error)

	ListByGroupID(
		ctx context.Context,
		groupID identity.PrayerGroupID,
	) ([]PrayerSession, error)

	Save(
		ctx context.Context,
		session *PrayerSession,
	) error

	Update(
		ctx context.Context,
		session *PrayerSession,
	) error

	Delete(
		ctx context.Context,
		id identity.PrayerSessionID,
	) error
}