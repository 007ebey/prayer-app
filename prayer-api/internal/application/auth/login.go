package auth

import (
	"context"
	"errors"

	"prayer-api/internal/domain/prayergroup"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

var (
	ErrIdentityRequired     = errors.New("authenticated identity is required")
	ErrDefaultRoleNotFound  = errors.New("default member role not found")
	ErrVisitorGroupNotFound = errors.New("visitor prayer group not found")
	ErrVisitorGroupInactive = errors.New("visitor prayer group is inactive")
)

type UserRepository interface {
	FindByExternalID(ctx context.Context, externalID string) (*user.User, error)
	Save(ctx context.Context, u *user.User) error
}

type RoleRepository interface {
	FindDefaultMemberRole(ctx context.Context) (*role.Role, error)
}

type PrayerGroupRepository interface {
	FindVisitorGroup(ctx context.Context) (*prayergroup.PrayerGroup, error)
	FindAccess(ctx context.Context, userID user.ID, groupID prayergroup.ID) (*prayergroup.Access, error)
	SaveAccess(ctx context.Context, access prayergroup.Access) error
}

type IDGenerator interface {
	NewUserID() user.ID
}

type LoginCommand struct {
	ExternalID string
	Name       string
}

type LoginResult struct {
	User             *user.User
	PrayerGroupAccess []prayergroup.Access
	Created           bool
}

type Service struct {
	users  UserRepository
	roles  RoleRepository
	groups PrayerGroupRepository
	ids    IDGenerator
}

func NewService(users UserRepository, roles RoleRepository, groups PrayerGroupRepository, ids IDGenerator) *Service {
	return &Service{
		users:  users,
		roles:  roles,
		groups: groups,
		ids:    ids,
	}
}

func (s *Service) Login(ctx context.Context, cmd LoginCommand) (*LoginResult, error) {
	if cmd.ExternalID == "" {
		return nil, ErrIdentityRequired
	}

	existing, err := s.users.FindByExternalID(ctx, cmd.ExternalID)
	if err != nil {
		return nil, err
	}

	if existing != nil {
		if !existing.IsActive() {
			return nil, user.ErrUserBlocked
		}

		return s.existingUserResult(ctx, existing)
	}

	return s.provisionUser(ctx, cmd)
}

func (s *Service) provisionUser(ctx context.Context, cmd LoginCommand) (*LoginResult, error) {
	memberRole, err := s.roles.FindDefaultMemberRole(ctx)
	if err != nil {
		return nil, err
	}

	if memberRole == nil {
		return nil, ErrDefaultRoleNotFound
	}

	visitorGroup, err := s.groups.FindVisitorGroup(ctx)
	if err != nil {
		return nil, err
	}

	if visitorGroup == nil {
		return nil, ErrVisitorGroupNotFound
	}

	if !visitorGroup.IsActive() {
		return nil, ErrVisitorGroupInactive
	}

	newUser, err := user.New(
		s.ids.NewUserID(),
		cmd.ExternalID,
		cmd.Name,
		memberRole.ID,
	)
	if err != nil {
		return nil, err
	}

	visitorAccess := prayergroup.NewAccess(
		newUser.ID,
		visitorGroup.ID,
	)

	if err := s.users.Save(ctx, newUser); err != nil {
		return nil, err
	}

	if err := s.groups.SaveAccess(ctx, visitorAccess); err != nil {
		return nil, err
	}

	return &LoginResult{
		User: newUser,
		PrayerGroupAccess: []prayergroup.Access{
			visitorAccess,
		},
		Created: true,
	}, nil
}

func (s *Service) existingUserResult(ctx context.Context, existing *user.User) (*LoginResult, error) {
	visitorGroup, err := s.groups.FindVisitorGroup(ctx)
	if err != nil {
		return nil, err
	}

	if visitorGroup == nil {
		return nil, ErrVisitorGroupNotFound
	}

	access, err := s.groups.FindAccess(ctx, existing.ID, visitorGroup.ID)
	if err != nil {
		return nil, err
	}

	if access == nil {
		newAccess := prayergroup.NewAccess(
			existing.ID,
			visitorGroup.ID,
		)

		if err := s.groups.SaveAccess(ctx, newAccess); err != nil {
			return nil, err
		}

		access = &newAccess
	}

	return &LoginResult{
		User: existing,
		PrayerGroupAccess: []prayergroup.Access{
			*access,
		},
		Created: false,
	}, nil
}
