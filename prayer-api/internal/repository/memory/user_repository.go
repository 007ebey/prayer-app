package memory

import (
	"context"
	"strings"
	"sync"

	"prayer-api/internal/config"
	"prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/user"
)

type UserRepository struct {
	mu           sync.RWMutex
	byID         map[identity.UserID]*user.User
	byExternalID map[string]*user.User
	byEmail      map[string]*user.User
	config       config.Config
}

func NewUserRepository(cfg config.Config) *UserRepository {
	repo := &UserRepository{
		byID:         make(map[identity.UserID]*user.User),
		byExternalID: make(map[string]*user.User),
		byEmail:      make(map[string]*user.User),
		config:       cfg,
	}

	repo.seedDefaultAdmin()

	return repo
}

func normalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func (r *UserRepository) seedDefaultAdmin() {
	adminEmail := normalizeEmail(r.config.DefaultAdminEmail)

	admin := &user.User{
		ID:         identity.UserID("user-007ebey"),
		ExternalID: "clerk-user-007ebey",
		Name:       "Default Admin",
		Email:      identity.Email(adminEmail),
		Status:     user.StatusActive,
		RoleIDs: []identity.RoleID{
			identity.RoleID("role_members"),
			identity.RoleID("role_admin"),
		},
		PrayerGroupIDs: []identity.PrayerGroupID{
			identity.PrayerGroupID("group-public"),
		},
	}

	r.byID[admin.ID] = admin
	r.byExternalID[admin.ExternalID] = admin
	r.byEmail[adminEmail] = admin
}

func (r *UserRepository) FindByID(
	ctx context.Context,
	id identity.UserID,
) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.byID[id]
	if !exists {
		return nil, nil
	}

	return found, nil
}

func (r *UserRepository) FindByExternalID(
	ctx context.Context,
	externalID string,
) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.byExternalID[externalID]
	if !exists {
		return nil, nil
	}

	return found, nil
}

func (r *UserRepository) FindByEmail(
	ctx context.Context,
	email string,
) (*user.User, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.byEmail[normalizeEmail(email)]
	if !exists {
		return nil, nil
	}

	return found, nil
}

func (r *UserRepository) Save(
	ctx context.Context,
	u *user.User,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.byID[u.ID] = u
	r.byExternalID[u.ExternalID] = u
	r.byEmail[normalizeEmail(string(u.Email))] = u

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
	r.byEmail[normalizeEmail(string(u.Email))] = u

	return nil
}