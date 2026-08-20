package userprofile

import (
	"context"
	"errors"

	"prayer-api/internal/domain/prayergroup"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
	"prayer-api/internal/domain/identity"
)

var (
	ErrUserNotFound = errors.New("user not found")
	ErrForbidden    = errors.New("forbidden")
)

type UserRepository interface {
	FindByID(ctx context.Context, id identity.UserID) (*user.User, error)
	FindByExternalID(ctx context.Context, externalID string) (*user.User, error)
}

type RoleRepository interface {
	FindByID(ctx context.Context, id identity.RoleID) (*role.Role, error)
}

type PrayerGroupRepository interface {
	FindByID(ctx context.Context, id identity.PrayerGroupID) (*prayergroup.PrayerGroup, error)
	FindAccessByUserID(ctx context.Context, userID identity.UserID) ([]prayergroup.Access, error)
}

type RoleView struct {
	ID          identity.RoleID
	Name        string
	Permissions []role.Permission
}

type PrayerGroupView struct {
	ID           identity.PrayerGroupID
	Name         string
	AccessStatus prayergroup.AccessStatus
}

type Result struct {
	User         *user.User
	Roles        []RoleView
	PrayerGroups []PrayerGroupView
}

type Service struct {
	users  UserRepository
	roles  RoleRepository
	groups PrayerGroupRepository
}

func NewService(users UserRepository, roles RoleRepository, groups PrayerGroupRepository) *Service {
	return &Service{
		users:  users,
		roles:  roles,
		groups: groups,
	}
}

func (s *Service) Get(ctx context.Context, actorExternalID string, requestedUserID identity.UserID) (*Result, error) {
	actor, err := s.users.FindByExternalID(ctx, actorExternalID)
	if err != nil {
		return nil, err
	}

	if actor == nil {
		return nil, ErrUserNotFound
	}

	if actor.ID != requestedUserID {
		return nil, ErrForbidden
	}

	target, err := s.users.FindByID(ctx, requestedUserID)
	if err != nil {
		return nil, err
	}

	if target == nil {
		return nil, ErrUserNotFound
	}

	roles := make([]RoleView, 0, len(target.RoleIDs))

	for _, roleID := range target.RoleIDs {
		resolvedRole, err := s.roles.FindByID(ctx, roleID)
		if err != nil {
			return nil, err
		}

		if resolvedRole == nil {
			continue
		}

		roles = append(roles, RoleView{
			ID:          resolvedRole.ID,
			Name:        resolvedRole.Name,
			Permissions: resolvedRole.Permissions,
		})
	}

	accesses, err := s.groups.FindAccessByUserID(ctx, target.ID)
	if err != nil {
		return nil, err
	}

	groups := make([]PrayerGroupView, 0, len(accesses))

	for _, access := range accesses {
		group, err := s.groups.FindByID(ctx, access.GroupID)
		if err != nil {
			return nil, err
		}

		if group == nil {
			continue
		}

		groups = append(groups, PrayerGroupView{
			ID:           group.ID,
			Name:         group.Name,
			AccessStatus: access.Status,
		})
	}

	return &Result{
		User:         target,
		Roles:        roles,
		PrayerGroups: groups,
	}, nil
}
