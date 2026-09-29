
package presence

import (
    "context"

    "prayer-api/internal/domain/identity"
	domainSession "prayer-api/internal/domain/prayersession"
	domainUser "prayer-api/internal/domain/user"
)

type PrayerSessionRepository interface {
	FindByID(
		ctx context.Context,
		id identity.PrayerSessionID,
	) (*domainSession.PrayerSession, error)

	ListByGroupID(
		ctx context.Context,
		groupID identity.PrayerGroupID,
	) ([]domainSession.PrayerSession, error)

	Save(
		ctx context.Context,
		session *domainSession.PrayerSession,
	) error

	Update(
		ctx context.Context,
		session *domainSession.PrayerSession,
	) error

	Delete(
		ctx context.Context,
		id identity.PrayerSessionID,
	) error
}

type UserRepository interface {
	FindByExternalID(ctx context.Context, externalID string) (*domainUser.User, error)
	Save(ctx context.Context, u *domainUser.User) error
	FindByEmail(ctx context.Context, email string) (*domainUser.User, error)
	Update(ctx context.Context, u *domainUser.User) error
}