package auth

import 
  (
	"context"
	domainid "prayer-api/internal/domain/identity"
  )

type Identity struct {
	ExternalID string
	Name       string
	Email      domainid.Email
}

type Provider interface {
	GetIdentity(ctx context.Context, externalID string) (Identity, error)
}
