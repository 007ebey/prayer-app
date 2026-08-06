package auth

import (
	"context"
	"strings"

	clerkuser "github.com/clerk/clerk-sdk-go/v2/user"
)

type ClerkProvider struct{}

func NewClerkProvider() *ClerkProvider {
	return &ClerkProvider{}
}

func (p *ClerkProvider) GetIdentity(ctx context.Context, externalID string) (Identity, error) {
	clerkUser, err := clerkuser.Get(ctx, externalID)
	if err != nil {
		return Identity{}, err
	}

	parts := make([]string, 0, 2)

	if clerkUser.FirstName != nil {
		parts = append(parts, *clerkUser.FirstName)
	}

	if clerkUser.LastName != nil {
		parts = append(parts, *clerkUser.LastName)
	}

	name := strings.TrimSpace(strings.Join(parts, " "))

	if name == "" {
		name = "Prayer User"
	}

	return Identity{
		ExternalID: externalID,
		Name:       name,
	}, nil
}
