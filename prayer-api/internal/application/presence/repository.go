
package presence

import (
    "context"

    "prayer-api/internal/domain/identity"
    domain "prayer-api/internal/domain/presence"
)

type Repository interface {
    Save(
        ctx context.Context,
        presence *domain.Presence,
    ) error

    Find(
        ctx context.Context,
        sessionID identity.PrayerSessionID,
        userID identity.UserID,
    ) (*domain.Presence, error)

    ListBySessionID(
        ctx context.Context,
        sessionID identity.PrayerSessionID,
    ) ([]*domain.Presence, error)

    Delete(
        ctx context.Context,
        sessionID identity.PrayerSessionID,
        userID identity.UserID,
    ) error
}