package prayergroup

import "prayer-api/internal/domain/user"

type AccessStatus string

const (
	AccessActive  AccessStatus = "active"
	AccessBlocked AccessStatus = "blocked"
)

type Access struct {
	UserID  user.ID
	GroupID ID
	Status  AccessStatus
}

func NewAccess(userID user.ID, groupID ID) Access {
	return Access{
		UserID:  userID,
		GroupID: groupID,
		Status:  AccessActive,
	}
}

func (a *Access) IsActive() bool {
	return a.Status == AccessActive
}

func (a *Access) Block() {
	a.Status = AccessBlocked
}

func (a *Access) Unblock() {
	a.Status = AccessActive
}
