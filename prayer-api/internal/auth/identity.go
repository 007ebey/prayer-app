package auth

import "context"

type Identity struct {
	ExternalID string
	Name       string
}

type Provider interface {
	GetIdentity(ctx context.Context, externalID string) (Identity, error)
}
