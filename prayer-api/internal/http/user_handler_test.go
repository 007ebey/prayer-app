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
	identity "prayer-api/internal/domain/identity"
	"prayer-api/internal/config"
)

type userHandlerTestEnvironment struct {
	handler *UserHandler
	login   *appauth.Service
}

func newUserHandlerTestEnvironment() *userHandlerTestEnvironment {
	users := memory.NewUserRepository(config.Config{})
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

	return &userHandlerTestEnvironment{
		handler: NewUserHandler(profileService),
		login:   loginService,
	}
}

func authenticatedUserRequest(
	externalID string,
	userID string,
) *http.Request {
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

type userProfileResponse struct {
	User struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Email  string `json:"email"`
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

func TestGetUserAPIReturnsOwnProfile(t *testing.T) {
	env := newUserHandlerTestEnvironment()

	loginResult, err := env.login.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: "clerk_api_profile",
			Name:       "Anna Mary",
			Email:      identity.Email("anna.mary@example.com"),
		},
	)

	if err != nil {
		t.Fatalf("login failed: %v", err)
	}

	response := httptest.NewRecorder()

	env.handler.Get(
		response,
		authenticatedUserRequest(
			"clerk_api_profile",
			string(loginResult.User.ID),
		),
	)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"expected HTTP 200, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}

	var body userProfileResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if body.User.ID != string(loginResult.User.ID) {
		t.Fatalf(
			"expected user %s, got %s",
			loginResult.User.ID,
			body.User.ID,
		)
	}

	if body.User.Name != "Anna Mary" {
		t.Fatalf(
			"expected Anna Mary, got %s",
			body.User.Name,
		)
	}

	if body.User.Email != "anna.mary@example.com" {
		t.Fatalf(
			"expected anna.mary@example.com, got %s",
			body.User.Email,
		)
	}

	if body.User.Status != "active" {
		t.Fatalf(
			"expected active user, got %s",
			body.User.Status,
		)
	}
}

func TestGetUserAPIReturnsRoleDetails(t *testing.T) {
	env := newUserHandlerTestEnvironment()

	loginResult, err := env.login.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: "clerk_api_roles",
			Name:       "Role User",
			Email:      identity.Email("role.user@example.com"),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	env.handler.Get(
		response,
		authenticatedUserRequest(
			"clerk_api_roles",
			string(loginResult.User.ID),
		),
	)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"expected HTTP 200, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}

	var body userProfileResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}

	if len(body.Roles) != 1 {
		t.Fatalf(
			"expected one role, got %d",
			len(body.Roles),
		)
	}

	if body.Roles[0].ID != "role_members" {
		t.Fatalf(
			"expected role_members, got %s",
			body.Roles[0].ID,
		)
	}

	if body.Roles[0].Name != "Members" {
		t.Fatalf(
			"expected Members, got %s",
			body.Roles[0].Name,
		)
	}

	if len(body.Roles[0].Permissions) == 0 {
		t.Fatal("expected Members permissions")
	}
}

func TestGetUserAPIReturnsVisitorGroup(t *testing.T) {
	env := newUserHandlerTestEnvironment()

	loginResult, err := env.login.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: "clerk_api_groups",
			Name:       "Group User",
			Email:      identity.Email("group.user@example.com"),
		},
	)

	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()

	env.handler.Get(
		response,
		authenticatedUserRequest(
			"clerk_api_groups",
			string(loginResult.User.ID),
		),
	)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"expected HTTP 200, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}

	var body userProfileResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}

	if len(body.PrayerGroups) != 1 {
		t.Fatalf(
			"expected one prayer group, got %d",
			len(body.PrayerGroups),
		)
	}

	group := body.PrayerGroups[0]

	if group.ID != "visitor" {
		t.Fatalf(
			"expected visitor, got %s",
			group.ID,
		)
	}

	if group.Name != "Visitor" {
		t.Fatalf(
			"expected Visitor, got %s",
			group.Name,
		)
	}

	if group.AccessStatus != "active" {
		t.Fatalf(
			"expected active access, got %s",
			group.AccessStatus,
		)
	}
}

func TestGetUserAPIRejectsOtherUser(t *testing.T) {
	env := newUserHandlerTestEnvironment()

	actor, err := env.login.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: "clerk_actor_api",
			Name:       "Actor",
			Email:      identity.Email("actor@example.com"),
		},
	)

	if err != nil {
		t.Fatalf("actor login failed: %v", err)
	}

	target, err := env.login.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: "clerk_target_api",
			Name:       "Target",
			Email:      identity.Email("target@example.com"),
		},
	)

	if err != nil {
		t.Fatalf("target login failed: %v", err)
	}

	response := httptest.NewRecorder()

	env.handler.Get(
		response,
		authenticatedUserRequest(
			actor.User.ExternalID,
			string(target.User.ID),
		),
	)

	if response.Code != http.StatusForbidden {
		t.Fatalf(
			"expected HTTP 403, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}
}

func TestGetUserAPIRequiresAuthentication(t *testing.T) {
	env := newUserHandlerTestEnvironment()

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/users/user_1",
		nil,
	)

	req.SetPathValue("id", "user_1")

	response := httptest.NewRecorder()

	env.handler.Get(response, req)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf(
			"expected HTTP 401, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}
}

func TestGetUserAPIUnknownActorReturns404(t *testing.T) {
	env := newUserHandlerTestEnvironment()

	response := httptest.NewRecorder()

	env.handler.Get(
		response,
		authenticatedUserRequest(
			"clerk_unknown_api",
			"user_999",
		),
	)

	if response.Code != http.StatusNotFound {
		t.Fatalf(
			"expected HTTP 404, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}
}