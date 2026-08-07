package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
	"prayer-api/internal/application/userprofile"
	"prayer-api/internal/repository/memory"
)

func newUserHandlerTestEnvironment() (*UserHandler, *appauth.Service) {
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

	profileService := userprofile.NewService(
		users,
		roles,
		groups,
	)

	return NewUserHandler(profileService), loginService
}

func userRequest(externalID string, userID string) *http.Request {
	req := httptest.NewRequest(
		http.MethodGet,
		"/api/users/"+userID,
		nil,
	)

	req.SetPathValue("id", userID)

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

func TestGetUserAPIReturnsOwnProfile(t *testing.T) {
	handler, loginService := newUserHandlerTestEnvironment()

	loginResult, err := loginService.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: "clerk_api_profile",
			Name:       "Anna Mary",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	handler.Get(
		response,
		userRequest(
			"clerk_api_profile",
			string(loginResult.User.ID),
		),
	)

	if response.Code != http.StatusOK {
		t.Fatalf("expected HTTP 200, got %d: %s", response.Code, response.Body.String())
	}

	var body struct {
		User struct {
			ID     string `json:"id"`
			Name   string `json:"name"`
			Status string `json:"status"`
		} `json:"user"`

		Roles []struct {
			ID          string   `json:"id"`
			Name        string   `json:"name"`
			Permissions []string `json:"permissions"`
		} `json:"roles"`

		PrayerGroups []struct {
			ID           string `json:"id"`
			Name         string `json:"name"`
			AccessStatus string `json:"accessStatus"`
		} `json:"prayerGroups"`
	}

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body.User.ID != string(loginResult.User.ID) {
		t.Fatalf("expected user %s, got %s", loginResult.User.ID, body.User.ID)
	}

	if len(body.Roles) != 1 || body.Roles[0].ID != "role_members" {
		t.Fatalf("expected Members role, got %+v", body.Roles)
	}

	if len(body.PrayerGroups) != 1 || body.PrayerGroups[0].ID != "visitor" {
		t.Fatalf("expected Visitor prayer group, got %+v", body.PrayerGroups)
	}
}

func TestGetUserAPIRejectsOtherUser(t *testing.T) {
	handler, loginService := newUserHandlerTestEnvironment()

	actor, err := loginService.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: "clerk_actor_api",
			Name:       "Actor",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	target, err := loginService.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: "clerk_target_api",
			Name:       "Target",
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	handler.Get(
		response,
		userRequest(
			actor.User.ExternalID,
			string(target.User.ID),
		),
	)

	if response.Code != http.StatusForbidden {
		t.Fatalf("expected HTTP 403, got %d: %s", response.Code, response.Body.String())
	}
}

func TestGetUserAPIRequiresAuthentication(t *testing.T) {
	handler, _ := newUserHandlerTestEnvironment()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/users/user_1",
		nil,
	)

	req.SetPathValue("id", "user_1")

	response := httptest.NewRecorder()

	handler.Get(response, req)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf("expected HTTP 401, got %d: %s", response.Code, response.Body.String())
	}
}
