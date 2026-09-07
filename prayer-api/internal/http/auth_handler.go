package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
	identityauth "prayer-api/internal/auth"
	domainid "prayer-api/internal/domain/identity"
)

type AuthHandler struct {
	login    *appauth.Service
	identity identityauth.Provider
}

func NewAuthHandler(
	login *appauth.Service,
	identity identityauth.Provider,
) *AuthHandler {
	return &AuthHandler{
		login:    login,
		identity: identity,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Verify that the Clerk session claims exist in the request context.
	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok || claims == nil || claims.Subject == "" {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "unauthorized",
		})
		return
	}

	// Retrieve the identity associated with the Clerk subject.
	externalIdentity, err := h.identity.GetIdentity(
		r.Context(),
		claims.Subject,
	)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"error": "unable to retrieve identity",
		})
		return
	}

	// Provision or retrieve the application user.
	result, err := h.login.Login(
		r.Context(),
		appauth.LoginCommand{
			ExternalID: externalIdentity.ExternalID,
			Name:       externalIdentity.Name,
			Email:      domainid.Email(externalIdentity.Email),
		},
	)
	if err != nil {
		writeJSON(w, http.StatusInternalServerError, map[string]any{
			"error": err.Error(),
		})
		return
	}

	status := http.StatusOK
	if result.Created {
		status = http.StatusCreated
	}

	groups := make([]map[string]any, 0, len(result.PrayerGroupAccess))

	for _, access := range result.PrayerGroupAccess {
		groups = append(groups, map[string]any{
			"groupId": access.GroupID,
			"status":  access.Status,
		})
	}

	writeJSON(w, status, map[string]any{
		"user": map[string]any{
			"id":     result.User.ID,
			"name":   result.User.Name,
			"email":  result.User.Email,
			"status": result.User.Status,
			"roles":  result.User.RoleIDs,
		},
		"prayerGroups": groups,
		"created":      result.Created,
	})
}

func writeJSON(w http.ResponseWriter, status int, value any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(value)
}