package prayergroup

import (
	"context"
	"errors"

	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainuser "prayer-api/internal/domain/user"
)

type AssignCommand struct {
	ActorExternalID string
	UserID          domainuser.ID
	GroupID         domainprayergroup.ID
}

type AssignPrayerGroupService struct {
	users        UserRepository
	groups       PrayerGroupRepository
}

func NewAssignPrayerGroupService(
	users UserRepository,
	groups PrayerGroupRepository,
) *AssignPrayerGroupService {
	return &AssignPrayerGroupService{
		users:  users,
		groups: groups,
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

	if actor.Role != domainuser.RoleAdmin {
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

	if err := s.userRepository.Update(
		ctx,
		user,
	); err != nil {
		return err
	}

	return nil
}