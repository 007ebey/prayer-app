package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clerk/clerk-sdk-go/v2"
	appauth "prayer-api/internal/application/auth"
	identityauth "prayer-api/internal/auth"
	"prayer-api/internal/repository/memory"
)

type fakeIdentityProvider struct {
	identity identityauth.Identity
	err      error
}

func (f *fakeIdentityProvider) GetIdentity(ctx context.Context, externalID string) (identityauth.Identity, error) {
	if f.err != nil {
		return identityauth.Identity{}, f.err
	}

	return f.identity, nil
}

func newTestAuthHandler(identity identityauth.Provider) *AuthHandler {
	users := memory.NewUserRepository()
	roles := memory.NewRoleRepository()
	groups := memory.NewPrayerGroupRepository()
	ids := memory.NewIDGenerator()

	loginService := appauth.NewService(
		users,
		roles,
		groups,
		ids,
	)

	return NewAuthHandler(loginService, identity)
}

func authenticatedRequest(externalID string) *http.Request {
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/auth/login",
		nil,
	)

	claims := &clerk.SessionClaims{
		RegisteredClaims: clerk.RegisteredClaims{
			Subject: externalID,
		},
	}

	ctx := clerk.ContextWithSessionClaims(
		req.Context(),
		claims,
	)

	return req.WithContext(ctx)
}

type loginResponse struct {
	Created bool `json:"created"`

	User struct {
		ID     string   `json:"id"`
		Name   string   `json:"name"`
		Status string   `json:"status"`
		Roles  []string `json:"roles"`
	} `json:"user"`

	PrayerGroups []struct {
		GroupID string `json:"groupId"`
		Status  string `json:"status"`
	} `json:"prayerGroups"`
}

func TestLoginAPICreatesUser(t *testing.T) {
	identity := &fakeIdentityProvider{
		identity: identityauth.Identity{
			ExternalID: "clerk_test_123",
			Name:       "Anna Mary",
		},
	}

	handler := newTestAuthHandler(identity)

	req := authenticatedRequest("clerk_test_123")
	response := httptest.NewRecorder()

	handler.Login(response, req)

	if response.Code != http.StatusCreated {
		t.Fatalf(
			"expected HTTP 201, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}

	var body loginResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !body.Created {
		t.Fatal("expected created=true")
	}

	if body.User.Name != "Anna Mary" {
		t.Fatalf("expected Anna Mary, got %s", body.User.Name)
	}

	if body.User.Status != "active" {
		t.Fatalf("expected active status, got %s", body.User.Status)
	}

	if len(body.User.Roles) != 1 {
		t.Fatalf("expected one role, got %d", len(body.User.Roles))
	}

	if body.User.Roles[0] != "role_members" {
		t.Fatalf("expected role_members, got %s", body.User.Roles[0])
	}

	if len(body.PrayerGroups) != 1 {
		t.Fatalf("expected one prayer group, got %d", len(body.PrayerGroups))
	}

	if body.PrayerGroups[0].GroupID != "visitor" {
		t.Fatalf("expected visitor group, got %s", body.PrayerGroups[0].GroupID)
	}

	if body.PrayerGroups[0].Status != "active" {
		t.Fatalf("expected active visitor access, got %s", body.PrayerGroups[0].Status)
	}
}

func TestLoginAPIExistingUserReturns200(t *testing.T) {
	identity := &fakeIdentityProvider{
		identity: identityauth.Identity{
			ExternalID: "clerk_existing_123",
			Name:       "Existing User",
		},
	}

	handler := newTestAuthHandler(identity)

	first := httptest.NewRecorder()
	handler.Login(first, authenticatedRequest("clerk_existing_123"))

	if first.Code != http.StatusCreated {
		t.Fatalf("expected first request HTTP 201, got %d", first.Code)
	}

	second := httptest.NewRecorder()
	handler.Login(second, authenticatedRequest("clerk_existing_123"))

	if second.Code != http.StatusOK {
		t.Fatalf(
			"expected second request HTTP 200, got %d: %s",
			second.Code,
			second.Body.String(),
		)
	}

	var body loginResponse

	if err := json.NewDecoder(second.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body.Created {
		t.Fatal("expected created=false on second login")
	}
}

func TestLoginAPIWithoutAuthenticationReturns401(t *testing.T) {
	identity := &fakeIdentityProvider{
		identity: identityauth.Identity{
			ExternalID: "unused",
			Name:       "Unused",
		},
	}

	handler := newTestAuthHandler(identity)

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/auth/login",
		nil,
	)

	response := httptest.NewRecorder()

	handler.Login(response, req)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf(
			"expected HTTP 401, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}
}

func TestLoginAPIIdentityProviderFailureReturns502(t *testing.T) {
	identity := &fakeIdentityProvider{
		err: errors.New("identity provider unavailable"),
	}

	handler := newTestAuthHandler(identity)

	response := httptest.NewRecorder()
	handler.Login(response, authenticatedRequest("clerk_failure_test"))

	if response.Code != http.StatusBadGateway {
		t.Fatalf(
			"expected HTTP 502, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}
}
