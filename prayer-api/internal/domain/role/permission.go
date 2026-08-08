package role

type Permission string

const (
	PermissionViewPrayerSessions   Permission = "prayer_sessions:view"
	PermissionJoinPrayerSessions   Permission = "prayer_sessions:join"
	PermissionHostPrayerSessions   Permission = "prayer_sessions:host"
	PermissionManagePrayerSessions Permission = "prayer_sessions:manage"
	PermissionManageUsers          Permission = "users:manage"
	PermissionManageAccessGroups   Permission = "access_groups:manage"
	PermissionManagePrayerGroups   Permission = "prayer_groups:manage"
)

var AllPermissions = []Permission{
	PermissionViewPrayerSessions,
	PermissionJoinPrayerSessions,
	PermissionHostPrayerSessions,
	PermissionManagePrayerSessions,
	PermissionManageUsers,
	PermissionManageAccessGroups,
	PermissionManagePrayerGroups,
}

func IsValidPermission(permission Permission) bool {
	for _, allowed := range AllPermissions {
		if allowed == permission {
			return true
		}
	}

	return false
}
