package userrole

import (
	"context"
	"errors"
     "prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

var (
	ErrForbidden    = errors.New("forbidden")
	ErrUserNotFound = errors.New("user not found")
	ErrRoleNotFound = errors.New("role not found")
)

type UserRepository interface {
	FindByID(ctx context.Context, id identity.UserID) (*user.User, error)
	FindByExternalID(ctx context.Context, externalID string) (*user.User, error)
	Save(ctx context.Context, u *user.User) error
}

type RoleRepository interface {
	FindByID(ctx context.Context, id identity.RoleID) (*role.Role, error)
}

type Service struct {
	users UserRepository
	roles RoleRepository
}

func NewService(users UserRepository, roles RoleRepository) *Service {
	return &Service{
		users: users,
		roles: roles,
	}
}

type AssignCommand struct {
	ActorExternalID string
	TargetUserID    identity.UserID
	RoleID          identity.RoleID
}

type AssignResult struct {
	User     *user.User
	Role     *role.Role
	Assigned bool
}

func (s *Service) Assign(ctx context.Context, cmd AssignCommand) (*AssignResult, error) {
	actor, err := s.users.FindByExternalID(ctx, cmd.ActorExternalID)
	if err != nil {
		return nil, err
	}

	if actor == nil {
		return nil, ErrForbidden
	}

	allowed, err := s.hasPermission(ctx, actor, role.PermissionManageUsers)
	if err != nil {
		return nil, err
	}

	if !allowed {
		return nil, ErrForbidden
	}

	target, err := s.users.FindByID(ctx, cmd.TargetUserID)
	if err != nil {
		return nil, err
	}

	if target == nil {
		return nil, ErrUserNotFound
	}

	resolvedRole, err := s.roles.FindByID(ctx, cmd.RoleID)
	if err != nil {
		if errors.Is(err, role.ErrNotFound) {
			return nil, ErrRoleNotFound
		}
		return nil, err
	}

	if resolvedRole == nil {
		return nil, ErrRoleNotFound
	}

	alreadyAssigned := false

	for _, existingRoleID := range target.RoleIDs {
		if existingRoleID == resolvedRole.ID {
			alreadyAssigned = true
			break
		}
	}

	if alreadyAssigned {
		return &AssignResult{
			User:     target,
			Role:     resolvedRole,
			Assigned: false,
		}, nil
	}

	if err := target.AssignRole(resolvedRole.ID); err != nil {
		return nil, err
	}

	if err := s.users.Save(ctx, target); err != nil {
		return nil, err
	}

	return &AssignResult{
		User:     target,
		Role:     resolvedRole,
		Assigned: true,
	}, nil
}

func (s *Service) hasPermission(ctx context.Context, actor *user.User, permission role.Permission) (bool, error) {
	for _, roleID := range actor.RoleIDs {
		resolvedRole, err := s.roles.FindByID(ctx, roleID)
		if err != nil {
			return false, err
		}

		if resolvedRole == nil {
			continue
		}

		for _, existingPermission := range resolvedRole.Permissions {
			if existingPermission == permission {
				return true, nil
			}
		}
	}

	return false, nil
}
