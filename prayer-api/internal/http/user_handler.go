package httpapi

import (
	"errors"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	"prayer-api/internal/application/userprofile"
	domainid "prayer-api/internal/domain/identity"
)

type UserHandler struct {
	profiles *userprofile.Service
}

func NewUserHandler(profiles *userprofile.Service) *UserHandler {
	return &UserHandler{
		profiles: profiles,
	}
}

func (h *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "unauthorized",
		})
		return
	}

	requestedUserID := domainid.UserID(r.PathValue("id"))

	result, err := h.profiles.Get(
		r.Context(),
		claims.Subject,
		requestedUserID,
	)

	if err != nil {
		switch {
		case errors.Is(err, userprofile.ErrForbidden):
			writeJSON(w, http.StatusForbidden, map[string]any{
				"error": "forbidden",
			})

		case errors.Is(err, userprofile.ErrUserNotFound):
			writeJSON(w, http.StatusNotFound, map[string]any{
				"error": "user not found",
			})

		default:
			writeJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
		}

		return
	}

	roles := make([]map[string]any, 0, len(result.Roles))

	for _, role := range result.Roles {
		roles = append(roles, map[string]any{
			"id":          role.ID,
			"name":        role.Name,
			"permissions": role.Permissions,
		})
	}

	groups := make([]map[string]any, 0, len(result.PrayerGroups))

	for _, group := range result.PrayerGroups {
		groups = append(groups, map[string]any{
			"id":           group.ID,
			"name":         group.Name,
			"accessStatus": group.AccessStatus,
		})
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"user": map[string]any{
			"id":     result.User.ID,
			"name":   result.User.Name,
			"status": result.User.Status,
		},
		"roles":        roles,
		"prayerGroups": groups,
	})
}
