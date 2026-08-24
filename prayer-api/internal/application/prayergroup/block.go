package prayergroup

import (
	"context"


	domainid   "prayer-api/internal/domain/identity"
	domainuser "prayer-api/internal/domain/user"
	
	domainrole "prayer-api/internal/domain/role"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
)

type BlockCommand struct {
	ActorExternalID string
    UserID  domainid.UserID
    GroupID domainid.PrayerGroupID
}

type BlockPrayerGroupService struct {
    users        UserRepository
    groups       PrayerGroupRepository
	roles        RoleRepository
}


func NewBlockPrayerGroupService (
    users        UserRepository,
    groups       PrayerGroupRepository,
	roles        RoleRepository,
) *BlockPrayerGroupService {
    return &BlockPrayerGroupService{
       users:  users,
	   groups: groups,
	   roles:  roles,
	}
}

func (s *BlockPrayerGroupService) Block(
	ctx context.Context,
	cmd BlockCommand,
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
		r, err := s.roles.FindByID(
			ctx,
			roleID,
		)
		if err != nil {
			return err
		}

		if r != nil &&
			r.HasPermission(
				domainrole.PermissionManagePrayerGroups,
			) {
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
		return domainprayergroup.ErrPrayerGroupNotFound
	}

	access, err := s.groups.FindAccess(
		ctx,
		cmd.UserID,
		cmd.GroupID,
	)
	if err != nil {
		return err
	}

	if access == nil {
		return domainprayergroup.ErrPrayerGroupNotAssigned
	}

	access.Block()

	if err := s.groups.SaveAccess(
		ctx,
		*access,
	); err != nil {
		return err
	}

	return nil
}