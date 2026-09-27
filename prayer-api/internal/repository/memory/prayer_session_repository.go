package memory

import (
	"context"
	"sync"
    "time"
	"prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/prayersession"
)

type PrayerSessionRepository struct {
	mu       sync.RWMutex
	sessions map[identity.PrayerSessionID]*prayersession.PrayerSession
}


func NewPrayerSessionRepository() *PrayerSessionRepository {
	sessions := make(map[identity.PrayerSessionID]*prayersession.PrayerSession)

	now := time.Now()

	// Active session:
	// Started 15 minutes ago and lasts 60 minutes.
	activeStart := now.Add(-15 * time.Minute)

	// Upcoming session:
	// Starts tomorrow at 19:00.
	upcomingDate := now.AddDate(0, 0, 1)

	defaultSessions := []*prayersession.PrayerSession{
		{
			ID:            identity.PrayerSessionID("session-active"),
			PrayerGroupID: defaultPrayerGroupID,
			Title:         "Cool Evening Prayer",
			Description:   "A cool evening prayer session",
			Date:          activeStart,
			Time:          activeStart.Format("15:04"),
			Duration:      60,
			PrayerPointIDs: []identity.PrayerPointID{
				identity.PrayerPointID("prayer-point-001"),
				identity.PrayerPointID("prayer-point-002"),
			},
		},
		{
			ID:            identity.PrayerSessionID("session-upcoming"),
			PrayerGroupID: defaultPrayerGroupID,
			Title:         "Prayer for Students",
			Description:   "A prayer session for students",
			Date:          upcomingDate,
			Time:          "19:00",
			Duration:      45,
			PrayerPointIDs: []identity.PrayerPointID{
				identity.PrayerPointID("prayer-point-003"),
			},
		},
		{
			ID:            identity.PrayerSessionID("session-community"),
			PrayerGroupID: defaultPrayerGroupID,
			Title:         "Community Prayer",
			Date:          now.AddDate(0, 0, 2),
			Time:          "07:00",
			Description:   "A community prayer session",
			Duration:      30,
			PrayerPointIDs: []identity.PrayerPointID{
				identity.PrayerPointID("prayer-point-001"),
				identity.PrayerPointID("prayer-point-002"),
				identity.PrayerPointID("prayer-point-003"),
			},
		},
	}

	for _, session := range defaultSessions {
		copy := *session

		copy.PrayerPointIDs = append(
			[]identity.PrayerPointID(nil),
			session.PrayerPointIDs...,
		)

		sessions[session.ID] = &copy
	}

	return &PrayerSessionRepository{
		sessions: sessions,
	}
}

func (r *PrayerSessionRepository) FindByID(
	ctx context.Context,
	id identity.PrayerSessionID,
) (*prayersession.PrayerSession, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.sessions[id]
	if !exists {
		return nil, nil
	}

	copy := *found

	// Copy the slice as well so callers cannot mutate repository state.
	copy.PrayerPointIDs = append(
		[]identity.PrayerPointID(nil),
		found.PrayerPointIDs...,
	)

	return &copy, nil
}

func (r *PrayerSessionRepository) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]prayersession.PrayerSession, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]prayersession.PrayerSession, 0)

	for _, session := range r.sessions {
		if session.PrayerGroupID == groupID {
			copy := *session

			copy.PrayerPointIDs = append(
				[]identity.PrayerPointID(nil),
				session.PrayerPointIDs...,
			)

			result = append(result, copy)
		}
	}

	return result, nil
}

func (r *PrayerSessionRepository) Save(
	ctx context.Context,
	session *prayersession.PrayerSession,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	copy := *session

	copy.PrayerPointIDs = append(
		[]identity.PrayerPointID(nil),
		session.PrayerPointIDs...,
	)

	r.sessions[session.ID] = &copy

	return nil
}

func (r *PrayerSessionRepository) Update(
	ctx context.Context,
	session *prayersession.PrayerSession,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.sessions[session.ID]; !exists {
		return prayersession.ErrNotFound
	}

	copy := *session

	copy.PrayerPointIDs = append(
		[]identity.PrayerPointID(nil),
		session.PrayerPointIDs...,
	)

	r.sessions[session.ID] = &copy

	return nil
}

func (r *PrayerSessionRepository) Delete(
	ctx context.Context,
	id identity.PrayerSessionID,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.sessions[id]; !exists {
		return prayersession.ErrNotFound
	}

	delete(r.sessions, id)

	return nil
}