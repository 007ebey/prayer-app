
package memory

import (
    "context"
    "sync"

    "prayer-api/internal/domain/identity"
    domain "prayer-api/internal/domain/presence"
)

type presenceKey struct {
    sessionID identity.PrayerSessionID
    userID    identity.UserID
}

type PresenceRepository struct {
    mu        sync.RWMutex
    presences map[presenceKey]*domain.Presence
}

func NewPresenceRepository() *PresenceRepository {
    return &PresenceRepository{
        presences: make(map[presenceKey]*domain.Presence),
    }
}

func (r *PresenceRepository) Save(
    ctx context.Context,
    p *domain.Presence,
) error {
    if err := ctx.Err(); err != nil {
        return err
    }

    r.mu.Lock()
    defer r.mu.Unlock()

    key := presenceKey{
        sessionID: p.SessionID,
        userID:    p.UserID,
    }

    copy := *p
    r.presences[key] = &copy

    return nil
}

func (r *PresenceRepository) Find(
    ctx context.Context,
    sessionID identity.PrayerSessionID,
    userID identity.UserID,
) (*domain.Presence, error) {
    if err := ctx.Err(); err != nil {
        return nil, err
    }

    r.mu.RLock()
    defer r.mu.RUnlock()

    p, ok := r.presences[presenceKey{
        sessionID: sessionID,
        userID:    userID,
    }]

    if !ok {
        return nil, nil
    }

    copy := *p
    return &copy, nil
}

func (r *PresenceRepository) ListBySessionID(
    ctx context.Context,
    sessionID identity.PrayerSessionID,
) ([]*domain.Presence, error) {
    if err := ctx.Err(); err != nil {
        return nil, err
    }

    r.mu.RLock()
    defer r.mu.RUnlock()

    result := make([]*domain.Presence, 0)

    for key, p := range r.presences {
        if key.sessionID == sessionID {
            copy := *p
            result = append(result, &copy)
        }
    }

    return result, nil
}

func (r *PresenceRepository) Delete(
    ctx context.Context,
    sessionID identity.PrayerSessionID,
    userID identity.UserID,
) error {
    if err := ctx.Err(); err != nil {
        return err
    }

    r.mu.Lock()
    defer r.mu.Unlock()

    delete(r.presences, presenceKey{
        sessionID: sessionID,
        userID:    userID,
    })

    return nil
}