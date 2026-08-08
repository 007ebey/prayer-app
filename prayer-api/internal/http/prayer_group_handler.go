package httpapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	appprayergroup "prayer-api/internal/application/prayergroup"
)

type PrayerGroupHandler struct {
	create *appprayergroup.CreateService
}

func NewPrayerGroupHandler(
	create *appprayergroup.CreateService,
) *PrayerGroupHandler {
	return &PrayerGroupHandler{
		create: create,
	}
}

type createPrayerGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (h *PrayerGroupHandler) Create(
	w http.ResponseWriter,
	r *http.Request,
) {
	claims, ok := clerk.SessionClaimsFromContext(
		r.Context(),
	)

	if !ok {
		writeJSON(
			w,
			http.StatusUnauthorized,
			map[string]any{
				"error": "unauthorized",
			},
		)

		return
	}

	var request createPrayerGroupRequest

	if err := json.NewDecoder(
		r.Body,
	).Decode(&request); err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]any{
				"error": "invalid request body",
			},
		)

		return
	}

	result, err := h.create.Create(
		r.Context(),
		appprayergroup.CreateCommand{
			ActorExternalID: claims.Subject,
			Name:            request.Name,
			Description:     request.Description,
		},
	)

	if err != nil {
		switch {
		case errors.Is(
			err,
			appprayergroup.ErrUnauthorized,
		):
			writeJSON(
				w,
				http.StatusUnauthorized,
				map[string]any{
					"error": "unauthorized",
				},
			)

		case errors.Is(
			err,
			appprayergroup.ErrActorNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "actor not found",
				},
			)

		case errors.Is(
			err,
			appprayergroup.ErrForbidden,
		):
			writeJSON(
				w,
				http.StatusForbidden,
				map[string]any{
					"error": "forbidden",
				},
			)

		case errors.Is(
			err,
			appprayergroup.ErrPrayerGroupExists,
		):
			writeJSON(
				w,
				http.StatusConflict,
				map[string]any{
					"error": "prayer group already exists",
				},
			)

		default:
			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]any{
					"error": err.Error(),
				},
			)
		}

		return
	}

	group := result.PrayerGroup

	writeJSON(
		w,
		http.StatusCreated,
		map[string]any{
			"prayerGroup": map[string]any{
				"id":          group.ID,
				"name":        group.Name,
				"description": group.Description,
				"status":      group.Status,
			},
		},
	)
}
