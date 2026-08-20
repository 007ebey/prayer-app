package memory

import (
	"context"
	"sync"
    "prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/user"
)

type UserRepository struct {
	mu           sync.RWMutex
	byID         map[identity.UserID]*user.User
	byExternalID map[string]*user.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		byID:         make(map[identity.UserID]*user.User),
		byExternalID: make(map[string]*user.User),
	}
}

func (r *UserRepository) FindByID(ctx context.Context, id identity.UserID) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.byID[id]
	if !exists {
		return nil, nil
	}

	return found, nil
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

func (r *UserRepository) Save(
	ctx context.Context,
	u *user.User) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.byID[u.ID] = u
	r.byExternalID[u.ExternalID] = u

	return nil
}

func (r *UserRepository) Update(
	ctx context.Context,
	u *user.User,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.byID[u.ID] = u
	r.byExternalID[u.ExternalID] = u

	return nil
}