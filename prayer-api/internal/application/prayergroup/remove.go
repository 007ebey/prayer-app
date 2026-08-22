package prayergroup

import (
	"context"
	"errors"

	domainuser "prayer-api/internal/domain/user"
	domainid "prayer-api/internal/domain/identity"
	domainrole "prayer-api/internal/domain/role"
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
    user, err := s.users.Get(ctx, cmd.UserID)
    if err != nil {
        return err
    }

    // Ensure the prayer group exists.
    if _, err := s.prayerGroups.Get(ctx, cmd.GroupID); err != nil {
        return err
    }

    if err := user.RemovePrayerGroup(cmd.GroupID); err != nil {
        return err
    }

    return s.users.Update(ctx, user)
}