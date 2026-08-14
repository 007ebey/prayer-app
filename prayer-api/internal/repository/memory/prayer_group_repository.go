package memory

import (
	"context"
	"sync"

	"prayer-api/internal/domain/prayergroup"
	"prayer-api/internal/domain/user"
)

type PrayerGroupRepository struct {
	mu      sync.RWMutex
	visitor *prayergroup.PrayerGroup
	groups  map[prayergroup.ID]*prayergroup.PrayerGroup
	access  map[string]prayergroup.Access
}

func NewPrayerGroupRepository() *PrayerGroupRepository {
	visitor, err := prayergroup.New(
		prayergroup.ID("visitor"),
		"Visitor",
		"Default prayer group access.",
		prayergroup.TypeVisitor,
	)

	if err != nil {
		panic(err)
	}

	return &PrayerGroupRepository{
		visitor: visitor,
		groups: map[prayergroup.ID]*prayergroup.PrayerGroup{
			visitor.ID: visitor,
		},
		access: make(map[string]prayergroup.Access),
	}
}

func accessKey(userID user.ID, groupID prayergroup.ID) string {
	return string(userID) + ":" + string(groupID)
}

func (r *PrayerGroupRepository) FindVisitorGroup(ctx context.Context) (*prayergroup.PrayerGroup, error) {
	return r.visitor, nil
}

func (r *PrayerGroupRepository) FindByID(ctx context.Context, id prayergroup.ID) (*prayergroup.PrayerGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.groups[id]
	if !exists {
		return nil, nil
	}

	return found, nil
}

func (r *PrayerGroupRepository) List(
	ctx context.Context,
	actorID user.ID,
	) ([]prayergroup.PrayerGroup, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]prayergroup.PrayerGroup, 0, len(r.groups))
	for _, group := range r.groups {
		result = append(result, *group)
	}

	return result, nil
}

func (r *PrayerGroupRepository) FindAccess(ctx context.Context, userID user.ID, groupID prayergroup.ID) (*prayergroup.Access, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.access[accessKey(userID, groupID)]
	if !exists {
		return nil, nil
	}

	copy := found
	return &copy, nil
}

func (r *PrayerGroupRepository) FindAccessByUserID(ctx context.Context, userID user.ID) ([]prayergroup.Access, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]prayergroup.Access, 0)

	for _, access := range r.access {
		if access.UserID == userID {
			result = append(result, access)
		}
	}

	return result, nil
}

func (r *PrayerGroupRepository) SaveAccess(ctx context.Context, access prayergroup.Access) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.access[accessKey(access.UserID, access.GroupID)] = access
	return nil
}

func (r *PrayerGroupRepository) Save(
	ctx context.Context,
	group *prayergroup.PrayerGroup,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.groups[group.ID] = group
	return nil
}

func (r *PrayerGroupRepository) Update(
	ctx context.Context,
	group *prayergroup.PrayerGroup,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.groups[group.ID] = group
	return nil
}

func (r *PrayerGroupRepository) Delete(
	ctx context.Context,
	id prayergroup.ID,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.groups[id]; !exists {
		return prayergroup.ErrNotFound
	}

	delete(r.groups, id)

	return nil
}
