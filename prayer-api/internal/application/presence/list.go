
package presence

import (
    "context"
    "time"

    "prayer-api/internal/domain/identity"
    domainuser "prayer-api/internal/domain/user"
	domainpresence "prayer-api/internal/domain/presence"
)

type Repository interface {
	Find(
		ctx context.Context,
		sessionID identity.PrayerSessionID,
		userID identity.UserID,
	) (*domainpresence.Presence, error)

	Save(
		ctx context.Context,
		presence *domainpresence.Presence,
	) error

    ListBySessionID(
        ctx context.Context,
        sessionID identity.PrayerSessionID,
    ) ([]*domainpresence.Presence, error)
}

type UserFinder interface {
    FindByID(
        ctx context.Context,
        id identity.UserID,
    ) (*domainuser.User, error)
}

type Participant struct {
    UserID   identity.UserID
    Name     string
    LastSeen time.Time
}

type ListParticipantsService struct {
    repo     Repository
    users    UserFinder
    timeout time.Duration
    now      func() time.Time
}

func NewListParticipantsService(
    repo Repository,
    users UserFinder,
    timeout time.Duration,
) *ListParticipantsService {
    return &ListParticipantsService{
        repo:    repo,
        users:   users,
        timeout: timeout,
        now:     time.Now,
    }
}

func (s *ListParticipantsService) Execute(
    ctx context.Context,
    sessionID identity.PrayerSessionID,
) ([]Participant, error) {
    records, err := s.repo.ListBySessionID(ctx, sessionID)
    if err != nil {
        return nil, err
    }

    now := s.now()
    participants := make([]Participant, 0, len(records))

    for _, record := range records {
        if !record.IsActive(now, s.timeout) {
            continue
        }

        user, err := s.users.FindByID(ctx, record.UserID)
        if err != nil {
            return nil, err
        }

        participants = append(participants, Participant{
            UserID:   record.UserID,
            Name:     user.Name,
            LastSeen: record.LastSeen,
        })
    }

    return participants, nil
}