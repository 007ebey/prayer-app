package prayersession

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	domainuser "prayer-api/internal/domain/user"
)

// ---------------------------------------------------------
// Session repository stub
// ---------------------------------------------------------

type listPrayerSessionRepositoryStub struct {
	sessionsByGroup map[identity.PrayerGroupID][]domain.PrayerSession
	listErr         error

	listCalled bool
	calledGroups []identity.PrayerGroupID
}

func (r *listPrayerSessionRepositoryStub) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*domain.PrayerSession, error) {
	return nil, nil
}

func (r *listPrayerSessionRepositoryStub) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]domain.PrayerSession, error) {

	r.listCalled = true
	r.calledGroups = append(r.calledGroups, groupID)

	if r.listErr != nil {
		return nil, r.listErr
	}

	return r.sessionsByGroup[groupID], nil
}

func (r *listPrayerSessionRepositoryStub) Save(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (r *listPrayerSessionRepositoryStub) Update(
	ctx context.Context,
	session *domain.PrayerSession,
) error {
	return nil
}

func (r *listPrayerSessionRepositoryStub) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	return nil
}

// ---------------------------------------------------------
// User repository stub
// ---------------------------------------------------------

type listUserRepositoryStub struct {
	user       *domainuser.User
	findErr    error
	findCalled bool
	externalID string
}

func (r *listUserRepositoryStub) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*domainuser.User, error) {

	r.findCalled = true
	r.externalID = externalID

	return r.user, r.findErr
}

func (r *listUserRepositoryStub) Save(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

func (r *listUserRepositoryStub) FindByEmail(
	ctx context.Context,
	email string,
) (*domainuser.User, error) {
	return nil, nil
}

func (r *listUserRepositoryStub) Update(
	ctx context.Context,
	u *domainuser.User,
) error {
	return nil
}

// ---------------------------------------------------------
// Helpers
// ---------------------------------------------------------

func newListTestSession(
	t *testing.T,
	id identity.PrayerSessionID,
	groupID identity.PrayerGroupID,
) domain.PrayerSession {
	t.Helper()

	session, err := domain.New(
		id,
		groupID,
		"Morning Prayer",
		testSessionDate(),
		"07:00",
		30,
		nil,
	)

	require.NoError(t, err)

	return *session
}

// Keep the test deterministic.
func testSessionDate() (t time.Time) {
	return time.Date(
		2026,
		9,
		15,
		0,
		0,
		0,
		0,
		time.UTC,
	)
}