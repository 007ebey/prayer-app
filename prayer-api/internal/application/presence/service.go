
package presence

import (
    "context"
    "time"

    "prayer-api/internal/domain/identity"
    domain "prayer-api/internal/domain/presence"
)

type SessionValidator interface {
    CanJoin(
        ctx context.Context,
        sessionID identity.PrayerSessionID,
        userID identity.UserID,
    ) error
}

type HeartbeatService struct {
    repo      Repository
    validator SessionValidator
    now       func() time.Time
}

func NewHeartbeatService(
    repo Repository,
    validator SessionValidator,
) *HeartbeatService {
    return &HeartbeatService{
        repo:      repo,
        validator: validator,
        now:       time.Now,
    }
}

func (s *HeartbeatService) Execute(
    ctx context.Context,
    sessionID identity.PrayerSessionID,
    userID identity.UserID,
) (*domain.Presence, error) {
    if err := s.validator.CanJoin(ctx, sessionID, userID); err != nil {
        return nil, err
    }

    current, err := s.repo.Find(ctx, sessionID, userID)
    if err != nil {
        return nil, err
    }

    now := s.now()

    if current == nil {
        current = domain.New(sessionID, userID, now)
    } else {
        current.Heartbeat(now)
    }

    if err := s.repo.Save(ctx, current); err != nil {
        return nil, err
    }

    return current, nil
}