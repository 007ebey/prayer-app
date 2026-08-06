package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
	clerkuser "github.com/clerk/clerk-sdk-go/v2/user"

	appauth "prayer-api/internal/application/auth"
)

type AuthHandler struct {
	login *appauth.Service
}

func NewAuthHandler(login *appauth.Service) *AuthHandler {
	return &AuthHandler{
		login: login,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "unauthorized",
		})
		return
	}

	clerkID := claims.Subject

	clerkUser, err := clerkuser.Get(r.Context(), clerkID)
	if err != nil {
		writeJSON(w, http.StatusBadGateway, map[string]any{
			"error": "unable to retrieve clerk user",
		})
		return
	}

	name := ""

    if clerkUser.FirstName != nil {
		name = *clerkUser.FirstName
    }

	if clerkUser.LastName != nil {
		if name != "" {
			name += " "
		}

		name += *clerkUser.LastName
  	}

	if name == "" {
		name = "Prayer User"
	}


	result, err := h.login.Login(
		r.Context(),
		appauth.LoginCommand{
			ExternalID: clerkID,
			Name:       name,
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
