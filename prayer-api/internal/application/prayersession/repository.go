package prayersession

import (
	"context"
	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	domainuser "prayer-api/internal/domain/user"
)

type PrayerSessionRepository interface {
	FindByID(
		ctx context.Context,
		id identity.PrayerSessionID,
	) (*domain.PrayerSession, error)

	ListByGroupID(
		ctx context.Context,
		groupID identity.PrayerGroupID,
	) ([]domain.PrayerSession, error)

	Save(
		ctx context.Context,
		session *domain.PrayerSession,
	) error

	Update(
		ctx context.Context,
		session *domain.PrayerSession,
	) error

	Delete(
		ctx context.Context,
		id identity.PrayerSessionID,
	) error
}

type UserRepository interface {
	FindByExternalID(ctx context.Context, externalID string) (*domainuser.User, error)
	Save(ctx context.Context, u *domainuser.User) error
	FindByEmail(ctx context.Context, email string) (*domainuser.User, error)
	Update(ctx context.Context, u *domainuser.User) error
}