
package presence

import (
    "time"

    "prayer-api/internal/domain/identity"
)

type Presence struct {
    SessionID identity.PrayerSessionID
    UserID    identity.UserID
    LastSeen  time.Time
}

func New(
    sessionID identity.PrayerSessionID,
    userID identity.UserID,
    now time.Time,
) *Presence {
    return &Presence{
        SessionID: sessionID,
        UserID:    userID,
        LastSeen:  now,
    }
}

func (p *Presence) Heartbeat(now time.Time) {
    p.LastSeen = now
}

func (p *Presence) IsActive(now time.Time, timeout time.Duration) bool {
    return !p.LastSeen.Before(now.Add(-timeout))
}