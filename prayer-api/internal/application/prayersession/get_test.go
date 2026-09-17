package prayersession

import (
	"context"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	domainuser "prayer-api/internal/domain/user"
)

// ---------------------------------------------------------
// Session repository stub
// ---------------------------------------------------------

type getPrayerSessionRepositoryStub struct {
	session       *domain.PrayerSession
	findErr       error
	findCalled    bool
	requestedID   identity.PrayerSessionID
}

func (r *getPrayerSessionRepositoryStub) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	r.findCalled = true
	r.requestedID = id

	return r.session, r.findErr
}

func (r *getPrayerSessionRepositoryStub) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {
	return nil, nil
}

func (r *getPrayerSessionRepositoryStub) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (r *getPrayerSessionRepositoryStub) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (r *getPrayerSessionRepositoryStub) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	return nil
}

// ---------------------------------------------------------
// User repository stub
// ---------------------------------------------------------

type getUserRepositoryStub struct {
	user         *domainuser.User
	findErr      error
	findCalled   bool
	externalID   string
}

func (r *getUserRepositoryStub) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {
	r.findCalled = true
	r.externalID = externalID

	return r.user, r.findErr
}

func (r *getUserRepositoryStub) Save(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

func (r *getUserRepositoryStub) FindByEmail(
	ctx context.Context,
	email string,
) (*domainuser.User, error) {
	return nil, nil
}

func (r *getUserRepositoryStub) Update(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}