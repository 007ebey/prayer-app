package memory

import (
	"fmt"
	"sync/atomic"

	"prayer-api/internal/domain/user"
)

type IDGenerator struct {
	next uint64
}

func NewIDGenerator() *IDGenerator {
	return &IDGenerator{}
}

func (g *IDGenerator) NewUserID() user.ID {
	id := atomic.AddUint64(&g.next, 1)
	return user.ID(fmt.Sprintf("user_%d", id))
}
