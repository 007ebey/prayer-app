package prayergroup

import (
	"context"
	domainrole "prayer-api/internal/domain/role"
	domainid "prayer-api/internal/domain/identity"
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
		groups: prayerGroups,
		users:  users,
		roles:  roles,
	}
}

func (s *DeleteService) Delete(
	ctx context.Context,
	actorID domainid.UserID,
	groupID domainid.PrayerGroupID,
) error {

	actor, err := s.users.FindByID(ctx, actorID)
	if err != nil {
		return err
	}

	_, err = s.groups.FindByID(ctx, groupID)
	if err != nil {
		return err
	}

	// Business rule
	actor, err2 := s.users.FindByID(ctx, actorID)
    if err2 != nil {
	  return err2
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

	_, err = s.groups.FindByID(ctx, groupID)
    if err != nil {
	  return err
    }

	return s.groups.Delete(ctx, groupID)
}
