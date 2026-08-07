package memory

import (
	"context"

	"prayer-api/internal/domain/role"
)

type RoleRepository struct {
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

	return &RoleRepository{
		memberRole: memberRole,
	}
}

func (r *RoleRepository) FindDefaultMemberRole(ctx context.Context) (*role.Role, error) {
	return r.memberRole, nil
}
