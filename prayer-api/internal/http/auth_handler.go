package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
	identityauth "prayer-api/internal/auth"
)

type AuthHandler struct {
	login    *appauth.Service
	identity identityauth.Provider
}

func NewAuthHandler(login *appauth.Service, identity identityauth.Provider) *AuthHandler {
	return &AuthHandler{
		login:    login,
		identity: identity,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	// Is the Clerk token valid?
	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "unauthorized",
		})
		return
	}

	// Does the token belong to someone?
	identity, err := h.identity.GetIdentity(r.Context(), claims.Subject)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"error": "unable to retrieve identity",
		})
		return
	}

	// Is the token expired?
	result, err := h.login.Login(
		r.Context(),
		appauth.LoginCommand{
			ExternalID: identity.ExternalID,
			Name:       identity.Name,
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
