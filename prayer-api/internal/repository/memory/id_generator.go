package memory

import (
	"fmt"
	"sync/atomic"

	"prayer-api/internal/domain/identity"
)

type IDGenerator struct {
	next uint64
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

func (g *IDGenerator) NewUserID() identity.UserID {
	id := atomic.AddUint64(&g.next, 1)
	return identity.UserID(fmt.Sprintf("user_%d", id))
}

func (g *IDGenerator) NewPrayerGroupID() identity.PrayerGroupID {
	id := atomic.AddUint64(&g.next, 1)
	return identity.PrayerGroupID(fmt.Sprintf("group_%d", id))
}
