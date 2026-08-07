package userrole

import (
	"context"
	"errors"

	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

var ErrProtectedRole = errors.New("protected role")

type RemoveCommand struct {
	ActorExternalID string
	TargetUserID    user.ID
	RoleID          role.ID
}

type RemoveResult struct {
	User    *user.User
	Role    *role.Role
	Removed bool
}

func (s *Service) Remove(
	ctx context.Context,
	cmd RemoveCommand,
) (*RemoveResult, error) {
	actor, err := s.users.FindByExternalID(
		ctx,
		cmd.ActorExternalID,
	)

	if err != nil {
		return nil, err
	}

	if actor == nil {
		return nil, ErrForbidden
	}

	allowed, err := s.hasPermission(
		ctx,
		actor,
		role.PermissionManageUsers,
	)

	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrForbidden
	}

	target, err := s.users.FindByID(
		ctx,
		cmd.TargetUserID,
	)

	if err != nil {
		return nil, err
	}

	if target == nil {
		return nil, ErrUserNotFound
	}

	resolvedRole, err := s.roles.FindByID(
		ctx,
		cmd.RoleID,
	)

	if err != nil {
		return nil, err
	}

	if resolvedRole == nil {
		return nil, ErrRoleNotFound
	}

	if resolvedRole.ID == role.ID("role_members") {
		return nil, ErrProtectedRole
	}

	removed := target.RemoveRole(
		resolvedRole.ID,
	)

	if !removed {
		return &RemoveResult{
			User:    target,
			Role:    resolvedRole,
			Removed: false,
		}, nil
	}

	if err := s.users.Save(
		ctx,
		target,
	); err != nil {
		return nil, err
	}

	return &RemoveResult{
		User:    target,
		Role:    resolvedRole,
		Removed: true,
	}, nil
}
