package httpapi

import (
	"errors"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	"prayer-api/internal/application/userrole"
	"prayer-api/internal/domain/role"
	"prayer-api/internal/domain/user"
)

type UserRoleHandler struct {
	service *userrole.Service
}

func NewUserRoleHandler(service *userrole.Service) *UserRoleHandler {
	return &UserRoleHandler{
		service: service,
	}
}

func (h *UserRoleHandler) Assign(w http.ResponseWriter, r *http.Request) {
	claims, ok := clerk.SessionClaimsFromContext(r.Context())
	if !ok {
		writeJSON(w, http.StatusUnauthorized, map[string]any{
			"error": "unauthorized",
		})
		return
	}

	targetUserID := user.ID(r.PathValue("userID"))
	roleID := role.ID(r.PathValue("roleID"))

	result, err := h.service.Assign(
		r.Context(),
		userrole.AssignCommand{
			ActorExternalID: claims.Subject,
			TargetUserID:    targetUserID,
			RoleID:          roleID,
		},
	)

	if err != nil {
		switch {
		case errors.Is(err, userrole.ErrForbidden):
			writeJSON(w, http.StatusForbidden, map[string]any{
				"error": "forbidden",
			})

		case errors.Is(err, userrole.ErrUserNotFound):
			writeJSON(w, http.StatusNotFound, map[string]any{
				"error": "user not found",
			})

		case errors.Is(err, userrole.ErrRoleNotFound):
			writeJSON(w, http.StatusNotFound, map[string]any{
				"error": "role not found",
			})

		default:
			writeJSON(w, http.StatusInternalServerError, map[string]any{
				"error": "internal server error",
			})
		}

		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"user": map[string]any{
			"id":     result.User.ID,
			"name":   result.User.Name,
			"status": result.User.Status,
			"roles":  result.User.RoleIDs,
		},
		"role": map[string]any{
			"id":          result.Role.ID,
			"name":        result.Role.Name,
			"permissions": result.Role.Permissions,
		},
		"assigned": result.Assigned,
	})
}
