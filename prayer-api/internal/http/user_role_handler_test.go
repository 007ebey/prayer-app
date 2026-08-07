package httpapi

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
	"prayer-api/internal/application/userrole"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
	"prayer-api/internal/repository/memory"
)

type userRoleHandlerTestEnvironment struct {
	handler *UserRoleHandler
	users   *memory.UserRepository
	login   *appauth.Service
}

func newUserRoleHandlerTestEnvironment() *userRoleHandlerTestEnvironment {
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

	service := userrole.NewService(
		users,
		roles,
	)

	return &userRoleHandlerTestEnvironment{
		handler: NewUserRoleHandler(service),
		users:   users,
		login:   loginService,
	}
}

func createHTTPTestUser(
	t *testing.T,
	env *userRoleHandlerTestEnvironment,
	externalID string,
	name string,
) *user.User {
	t.Helper()

	result, err := env.login.Login(
		t.Context(),
		appauth.LoginCommand{
			ExternalID: externalID,
			Name:       name,
		},
	)

	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	return result.User
}

func createHTTPTestAdministrator(
	t *testing.T,
	env *userRoleHandlerTestEnvironment,
	externalID string,
) *user.User {
	t.Helper()

	admin := createHTTPTestUser(
		t,
		env,
		externalID,
		"Administrator",
	)

	if err := admin.AssignRole(role.ID("role_admin")); err != nil {
		t.Fatalf("failed to assign Administrator role: %v", err)
	}

	if err := env.users.Save(t.Context(), admin); err != nil {
		t.Fatalf("failed to save Administrator: %v", err)
	}

	return admin
}

func authenticatedRoleRequest(
	externalID string,
	targetUserID string,
	roleID string,
) *http.Request {
	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users/"+targetUserID+"/roles/"+roleID,
		nil,
	)

	req.SetPathValue("userID", targetUserID)
	req.SetPathValue("roleID", roleID)

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

type assignRoleResponse struct {
	User struct {
		ID     string   `json:"id"`
		Name   string   `json:"name"`
		Status string   `json:"status"`
		Roles  []string `json:"roles"`
	} `json:"user"`

	Role struct {
		ID          string   `json:"id"`
		Name        string   `json:"name"`
		Permissions []string `json:"permissions"`
	} `json:"role"`

	Assigned bool `json:"assigned"`
}

func TestAssignRoleAPIReturns200(t *testing.T) {
	env := newUserRoleHandlerTestEnvironment()

	admin := createHTTPTestAdministrator(
		t,
		env,
		"clerk_http_admin",
	)

	target := createHTTPTestUser(
		t,
		env,
		"clerk_http_target",
		"Target User",
	)

	response := httptest.NewRecorder()

	env.handler.Assign(
		response,
		authenticatedRoleRequest(
			admin.ExternalID,
			string(target.ID),
			"role_admin",
		),
	)

	if response.Code != http.StatusOK {
		t.Fatalf(
			"expected HTTP 200, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}

	var body assignRoleResponse

	if err := json.NewDecoder(response.Body).Decode(&body); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	if !body.Assigned {
		t.Fatal("expected assigned=true")
	}

	if body.User.ID != string(target.ID) {
		t.Fatalf("expected user %s, got %s", target.ID, body.User.ID)
	}

	if body.Role.ID != "role_admin" {
		t.Fatalf("expected role_admin, got %s", body.Role.ID)
	}

	if body.Role.Name != "Administrator" {
		t.Fatalf("expected Administrator, got %s", body.Role.Name)
	}

	if len(body.Role.Permissions) == 0 {
		t.Fatal("expected Administrator permissions")
	}
}

func TestAssignRoleAPIMemberReturns403(t *testing.T) {
	env := newUserRoleHandlerTestEnvironment()

	member := createHTTPTestUser(
		t,
		env,
		"clerk_http_member",
		"Member",
	)

	target := createHTTPTestUser(
		t,
		env,
		"clerk_http_member_target",
		"Target",
	)

	response := httptest.NewRecorder()

	env.handler.Assign(
		response,
		authenticatedRoleRequest(
			member.ExternalID,
			string(target.ID),
			"role_admin",
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

func TestAssignRoleAPIRequiresAuthentication(t *testing.T) {
	env := newUserRoleHandlerTestEnvironment()

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/users/user_1/roles/role_admin",
		nil,
	)

	req.SetPathValue("userID", "user_1")
	req.SetPathValue("roleID", "role_admin")

	response := httptest.NewRecorder()

	env.handler.Assign(response, req)

	if response.Code != http.StatusUnauthorized {
		t.Fatalf(
			"expected HTTP 401, got %d: %s",
			response.Code,
			response.Body.String(),
		)
	}
}

func TestAssignRoleAPIUnknownTargetReturns404(t *testing.T) {
	env := newUserRoleHandlerTestEnvironment()

	admin := createHTTPTestAdministrator(
		t,
		env,
		"clerk_http_missing_target_admin",
	)

	response := httptest.NewRecorder()

	env.handler.Assign(
		response,
		authenticatedRoleRequest(
			admin.ExternalID,
			"user_missing",
			"role_admin",
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

func TestAssignRoleAPIUnknownRoleReturns404(t *testing.T) {
	env := newUserRoleHandlerTestEnvironment()

	admin := createHTTPTestAdministrator(
		t,
		env,
		"clerk_http_missing_role_admin",
	)

	target := createHTTPTestUser(
		t,
		env,
		"clerk_http_missing_role_target",
		"Target",
	)

	response := httptest.NewRecorder()

	env.handler.Assign(
		response,
		authenticatedRoleRequest(
			admin.ExternalID,
			string(target.ID),
			"role_missing",
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

func TestAssignRoleAPIIsIdempotent(t *testing.T) {
	env := newUserRoleHandlerTestEnvironment()

	admin := createHTTPTestAdministrator(
		t,
		env,
		"clerk_http_idempotent_admin",
	)

	target := createHTTPTestUser(
		t,
		env,
		"clerk_http_idempotent_target",
		"Target",
	)

	first := httptest.NewRecorder()

	env.handler.Assign(
		first,
		authenticatedRoleRequest(
			admin.ExternalID,
			string(target.ID),
			"role_admin",
		),
	)

	if first.Code != http.StatusOK {
		t.Fatalf("expected first HTTP 200, got %d", first.Code)
	}

	second := httptest.NewRecorder()

	env.handler.Assign(
		second,
		authenticatedRoleRequest(
			admin.ExternalID,
			string(target.ID),
			"role_admin",
		),
	)

	if second.Code != http.StatusOK {
		t.Fatalf("expected second HTTP 200, got %d", second.Code)
	}

	var body assignRoleResponse

	if err := json.NewDecoder(second.Body).Decode(&body); err != nil {
		t.Fatal(err)
	}

	if body.Assigned {
		t.Fatal("expected second request assigned=false")
	}

	count := 0

	for _, roleID := range body.User.Roles {
		if roleID == "role_admin" {
			count++
		}
	}

	if count != 1 {
		t.Fatalf("expected role_admin exactly once, got %d", count)
	}
}
