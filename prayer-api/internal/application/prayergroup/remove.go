package prayergroup

import (
	"context"

	domainid "prayer-api/internal/domain/identity"
    domainuser "prayer-api/internal/domain/user"
    domainprayergroup "prayer-api/internal/domain/prayergroup"
)

type RemoveCommand struct {
	UserID  domainid.UserID
	GroupID domainid.PrayerGroupID
}

type RemovePrayerGroupService struct {
	users        UserRepository
	prayerGroups PrayerGroupRepository
}

func NewRemovePrayerGroupService(
	users UserRepository,
	prayerGroups PrayerGroupRepository,
) *RemovePrayerGroupService {
	return &RemovePrayerGroupService{
		users:        users,
		prayerGroups: prayerGroups,
	}
}

func (s *RemovePrayerGroupService) Remove(
    ctx context.Context,
    cmd RemoveCommand,
) error {
    user, err := s.users.FindByID(ctx, cmd.UserID)
    if err != nil {
        return err
    }

    if user == nil {
        return domainuser.ErrUserNotFound
    }

    // Ensure the prayer group exists.
    if group, err := s.prayerGroups.FindByID(ctx, cmd.GroupID); err != nil {
	    return err
    } else if group == nil {
	    return domainprayergroup.ErrPrayerGroupNotFound
    }

    if err := user.RemovePrayerGroup(cmd.GroupID); err != nil {
        return err
    }

    return s.users.Update(ctx, user)
}