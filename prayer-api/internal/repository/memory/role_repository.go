package memory

import (
	"context"

	"prayer-api/internal/domain/role"
)

type RoleRepository struct {
	roles      map[role.ID]*role.Role
	memberRole *role.Role
}

func NewRoleRepository() *RoleRepository {
	memberRole, err := role.New(
		role.ID("role_members"),
		"Members",
		"Standard access to prayer sessions.",
		[]role.Permission{
			role.PermissionViewPrayerSessions,
			role.PermissionJoinPrayerSessions,
		},
		true,
	)

	if err != nil {
		panic(err)
	}

	adminRole, err := role.New(
		role.ID("role_admin"),
		"Administrator",
		"Administrative access to users, prayer sessions, and access groups.",
		[]role.Permission{
			role.PermissionViewPrayerSessions,
			role.PermissionJoinPrayerSessions,
			role.PermissionHostPrayerSessions,
			role.PermissionManagePrayerSessions,
			role.PermissionManageUsers,
			role.PermissionManageAccessGroups,
			role.PermissionManagePrayerGroups,
		},
		true,
	)

	if err != nil {
		panic(err)
	}

	return &RoleRepository{
		roles: map[role.ID]*role.Role{
			memberRole.ID: memberRole,
			adminRole.ID:  adminRole,
		},
		memberRole: memberRole,
	}
}

func (r *RoleRepository) FindDefaultMemberRole(ctx context.Context) (*role.Role, error) {
	return r.memberRole, nil
}

func (r *RoleRepository) FindByID(ctx context.Context, id role.ID) (*role.Role, error) {
	found, exists := r.roles[id]
	
	if !exists {
		return nil, role.ErrNotFound
	}

	return found, nil
}
