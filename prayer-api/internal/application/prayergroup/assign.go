package prayergroup

import (
	"context"
	"errors"

	domainuser "prayer-api/internal/domain/user"
	domainidentity "prayer-api/internal/domain/identity"
	domainrole "prayer-api/internal/domain/role"
)

type AssignCommand struct {
	ActorExternalID string
	UserID          domainidentity.UserID
	GroupID         domainidentity.PrayerGroupID
}

type AssignPrayerGroupService struct {
	users        UserRepository
	groups       PrayerGroupRepository
	roles  	     RoleRepository
}

func NewAssignPrayerGroupService(
	users UserRepository,
	groups PrayerGroupRepository,
	roles RoleRepository,
) *AssignPrayerGroupService {
	return &AssignPrayerGroupService{
		users:  users,
		groups: groups,
		roles:  roles,
	}
}

func (s *AssignPrayerGroupService) Assign(
	ctx context.Context,
	cmd AssignCommand,
) error {
	if cmd.ActorExternalID == "" {
		return ErrUnauthorized
	}

	actor, err := s.users.FindByExternalID(
		ctx,
		cmd.ActorExternalID,
	)

	if err != nil {
		return err
	}

	if actor == nil {
		return ErrActorNotFound
	}

	allowed := false

    for _, roleID := range actor.RoleIDs {
      r, err := s.roles.FindByID(ctx, roleID)
      if err != nil {
        return err
      }

      if r != nil &&
        r.HasPermission(domainrole.PermissionManagePrayerGroups) {
        allowed = true
        break
      }
    }

    if !allowed {
      return ErrForbidden
    }

	user, err := s.users.FindByID(
		ctx,
		cmd.UserID,
	)
	if err != nil {
		return err
	}

	if user == nil {
		return domainuser.ErrUserNotFound
	}

	group, err := s.groups.FindByID(
		ctx,
		cmd.GroupID,
	)
	if err != nil {
		return err
	}

	if group == nil {
		return ErrPrayerGroupNotFound
	}

	if err := user.AssignPrayerGroup(group.ID); err != nil {
		if errors.Is(
			err,
			domainuser.ErrPrayerGroupAlreadyAssigned,
		) {
			return err
		}

		return err
	}

	if err := s.users.Update(
		ctx,
		user,
	); err != nil {
		return err
	}

	return nil
}