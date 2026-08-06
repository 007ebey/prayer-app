package memory

import (
	"context"
	"sync"

	"prayer-api/internal/domain/user"
)

type UserRepository struct {
	mu           sync.RWMutex
	byID         map[user.ID]*user.User
	byExternalID map[string]*user.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		byID:         make(map[user.ID]*user.User),
		byExternalID: make(map[string]*user.User),
	}
}

func (r *UserRepository) FindByExternalID(ctx context.Context, externalID string) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.byExternalID[externalID]
	if !exists {
		return nil, nil
	}

	return found, nil
}

func (r *UserRepository) Save(ctx context.Context, u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.byID[u.ID] = u
	r.byExternalID[u.ExternalID] = u

	return nil
}
