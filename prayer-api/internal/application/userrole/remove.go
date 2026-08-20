package userrole

import (
	"context"
	"errors"

	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
	"prayer-api/internal/domain/identity"
)

var ErrProtectedRole = errors.New("protected role")

type RemoveCommand struct {
	ActorExternalID string
	TargetUserID    identity.UserID
	RoleID          identity.RoleID
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
		if errors.Is(err, role.ErrNotFound) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}

	if resolvedRole == nil {
		return nil, ErrRoleNotFound
	}

	if resolvedRole.ID == identity.RoleID("role_members") {
		return nil, ErrProtectedRole
	}

	wasAssigned := false

	for _, current := range target.RoleIDs {
		if current == resolvedRole.ID {
			wasAssigned = true
			break
		}
	}

	if !wasAssigned {
		return &RemoveResult{
			User:    target,
			Role:    resolvedRole,
			Removed: false,
		}, nil
	}

	if err := target.RemoveRole(
		resolvedRole.ID,
	); err != nil {
		return nil, err
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
