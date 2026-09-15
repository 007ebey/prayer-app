package memory

import (
	"context"
	"sync"

	"prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/prayersession"
)

type PrayerSessionRepository struct {
	mu       sync.RWMutex
	sessions map[identity.PrayerSessionID]*prayersession.PrayerSession
}

func NewPrayerSessionRepository() *PrayerSessionRepository {
	return &PrayerSessionRepository{
		sessions: make(map[identity.PrayerSessionID]*prayersession.PrayerSession),
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