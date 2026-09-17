package identity

type UserID string
type PrayerGroupID string
type RoleID string
type SessionID string
type Email string
type PrayerPointID string
type PrayerSessionID string

func (id PrayerSessionID) String() string {
	return string(id)
}

func (id PrayerGroupID) String() string {
	return string(id)
}

func (id UserID) String() string {
	return string(id)
}