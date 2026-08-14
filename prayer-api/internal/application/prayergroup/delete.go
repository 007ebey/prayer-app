package prayergroup

import (
	"context"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainrole "prayer-api/internal/domain/role"
)

type DeleteService struct {
    users  UserRepository
	roles  RoleRepository
	groups PrayerGroupRepository
}

func NewDeleteService(
	prayerGroups PrayerGroupRepository,
	users UserRepository,
	roles RoleRepository,
) *DeleteService {
	return &DeleteService{
		prayerGroups: prayerGroups,
		users:        users,
		roles:          roles,
	}
}

func (s *DeleteService) Delete(
	ctx context.Context,
	actorID user.ID,
	groupID domainprayergroup.ID,
) error {

	actor, err := s.users.FindByID(ctx, actorID)
	if err != nil {
		return err
	}

	group, err := s.prayerGroups.FindByID(ctx, groupID)
	if err != nil {
		return err
	}

	// Business rule
	actor, err := s.users.FindByID(ctx, actorID)
    if err != nil {
	  return err
    }

	allowed := false

	for _, roleID := range actor.RoleIDs {
	  r, err := s.roles.FindByID(ctx, roleID)
	  if err != nil {
		return err
	  }

	  if r.HasPermission(domainrole.PermissionManagePrayerGroups) {
		allowed = true
		break
	  }
    }

	if !allowed {
	  return ErrForbidden
    }

	_, err = s.prayerGroups.FindByID(ctx, groupID)
    if err != nil {
	  return err
    }

	return s.prayerGroups.Delete(ctx, groupID)
}
