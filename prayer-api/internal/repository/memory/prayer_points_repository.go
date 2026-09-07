package memory

import (
	"context"
	"sync"

	"prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/prayerpoint"
)

type PrayerPointRepository struct {
	mu     sync.RWMutex
	points map[identity.PrayerPointID]*prayerpoint.PrayerPoint
}

func NewPrayerPointRepository() *PrayerPointRepository {
	return &PrayerPointRepository{
		points: make(map[identity.PrayerPointID]*prayerpoint.PrayerPoint),
	}
}

func (r *PrayerPointRepository) FindByID(
	ctx context.Context,
	id identity.PrayerPointID,
) (*prayerpoint.PrayerPoint, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	found, exists := r.points[id]
	if !exists {
		return nil, nil
	}

	copy := *found
	return &copy, nil
}

func (r *PrayerPointRepository) ListByGroupID(
	ctx context.Context,
	groupID identity.PrayerGroupID,
) ([]prayerpoint.PrayerPoint, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	result := make([]prayerpoint.PrayerPoint, 0)

	for _, point := range r.points {
		if point.GroupID == groupID {
			result = append(result, *point)
		}
	}

	return result, nil
}

func (r *PrayerPointRepository) Save(
	ctx context.Context,
	point *prayerpoint.PrayerPoint,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	copy := *point
	r.points[point.ID] = &copy

	return nil
}

func (r *PrayerPointRepository) Update(
	ctx context.Context,
	point *prayerpoint.PrayerPoint,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.points[point.ID]; !exists {
		return prayerpoint.ErrNotFound
	}

	copy := *point
	r.points[point.ID] = &copy

	return nil
}

func (r *PrayerPointRepository) Delete(
	ctx context.Context,
	id identity.PrayerPointID,
) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.points[id]; !exists {
		return prayerpoint.ErrNotFound
	}

	delete(r.points, id)

	return nil
}

