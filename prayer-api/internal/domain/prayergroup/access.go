package prayergroup

import 

(
	"prayer-api/internal/domain/identity"

)

type AccessStatus string

const (
	AccessActive  AccessStatus = "active"
	AccessBlocked AccessStatus = "blocked"
)

type Access struct {
	UserID  identity.UserID
	GroupID identity.PrayerGroupID
	Status  AccessStatus
}

func NewAccess(userID identity.UserID, groupID identity.PrayerGroupID) Access {
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
