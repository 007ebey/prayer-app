package prayergroup

import (
	"context"
	"errors"
	"strings"

	domainprayergroup "prayer-api/internal/domain/prayergroup"
	domainid "prayer-api/internal/domain/identity"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

var (
	ErrUnauthorized      = errors.New("unauthorized")
	ErrForbidden         = errors.New("forbidden")
	ErrActorNotFound     = errors.New("actor not found")
	ErrPrayerGroupExists = errors.New("prayer group already exists")
)

type UserRepository interface {
	FindByExternalID(
		ctx context.Context,
		externalID string,
	) (*user.User, error)

	FindByID(
		ctx context.Context,
		id domainid.UserID,
	) (*user.User, error)

	Update(
		ctx context.Context,
		u *user.User,
	) error
}

type RoleRepository interface {
	FindByID(
		ctx context.Context,
		id domainid.RoleID,
	) (*role.Role, error)
}

type PrayerGroupRepository interface {
	FindByID(
		ctx context.Context,
		id domainid.PrayerGroupID,
	) (*domainprayergroup.PrayerGroup, error)

	Save(
		ctx context.Context,
		group *domainprayergroup.PrayerGroup,
	) error

	List(
		ctx context.Context,
		actorID domainid.UserID,
	) ([]domainprayergroup.PrayerGroup, error)

	Update(
		ctx context.Context,
		group *domainprayergroup.PrayerGroup,
	) error

	Delete(
		ctx context.Context,
		id domainid.PrayerGroupID,
	) error

	FindAccess(
		ctx context.Context, 
		userID domainid.UserID, 
		groupID domainid.PrayerGroupID,
	) (*domainprayergroup.Access, error)

	SaveAccess(
		ctx context.Context, 
		access domainprayergroup.Access,
	) error
}

type IDGenerator interface {
	NewPrayerGroupID() domainid.PrayerGroupID
}

type CreateCommand struct {
	ActorExternalID string
	Name            string
	Description     string
}

type CreateResult struct {
	PrayerGroup *domainprayergroup.PrayerGroup
}

type CreateService struct {
	users  UserRepository
	roles  RoleRepository
	groups PrayerGroupRepository
	ids    IDGenerator
}

func NewCreateService(
	users UserRepository,
	roles RoleRepository,
	groups PrayerGroupRepository,
	ids IDGenerator,
) *CreateService {
	return &CreateService{
		users:  users,
		roles:  roles,
		groups: groups,
		ids:    ids,
	}
}

func (s *CreateService) Create(
	ctx context.Context,
	command CreateCommand,
) (*CreateResult, error) {
	if strings.TrimSpace(command.ActorExternalID) == "" {
		return nil, ErrUnauthorized
	}

	actor, err := s.users.FindByExternalID(
		ctx,
		command.ActorExternalID,
	)

	if err != nil {
		return nil, err
	}

	if actor == nil {
		return nil, ErrActorNotFound
	}

	allowed := false

	for _, roleID := range actor.RoleIDs {
		resolvedRole, err := s.roles.FindByID(
			ctx,
			roleID,
		)

		if err != nil {
			return nil, err
		}

		if resolvedRole == nil {
			continue
		}

		if resolvedRole.HasPermission(
			role.PermissionManagePrayerGroups,
		) {
			allowed = true
			break
		}
	}

	if !allowed {
		return nil, ErrForbidden
	}

	groupID := s.ids.NewPrayerGroupID()

	existing, err := s.groups.FindByID(
		ctx,
		groupID,
	)

	if err != nil {
		return nil, err
	}

	if existing != nil {
		return nil, ErrPrayerGroupExists
	}

	group, err := domainprayergroup.New(
		groupID,
		command.Name,
		command.Description,
		domainprayergroup.TypeRegular,
	)

	if err != nil {
		return nil, err
	}

	if err := s.groups.Save(
		ctx,
		group,
	); err != nil {
		return nil, err
	}

	return &CreateResult{
		PrayerGroup: group,
	}, nil
}
