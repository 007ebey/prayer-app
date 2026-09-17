package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
    "time"
	"github.com/clerk/clerk-sdk-go/v2"

	domainid "prayer-api/internal/domain/identity"
	domainuser "prayer-api/internal/domain/user"
	domain "prayer-api/internal/domain/prayersession"
	appprayersession "prayer-api/internal/application/prayersession"
)

// ------------------------------------------------------------
// Service interfaces
// ------------------------------------------------------------

type CreatePrayerSessionService interface {
	Create(
		ctx context.Context,
		cmd appprayersession.CreateCommand,
	) (*domain.PrayerSession, error)
}

type UpdatePrayerSessionService interface {
	Update(
		ctx context.Context,
		cmd appprayersession.UpdateCommand,
	) (*domain.PrayerSession, error)
}

type DeletePrayerSessionService interface {
	Delete(
		ctx context.Context,
		userID domainid.UserID,
		sessionID domainid.PrayerSessionID,
	) error
}

// Add these when your application services exist.
type ListPrayerSessionService interface {
	ListForUser(
		ctx context.Context,
		userID domainid.UserID,
	) ([]domain.PrayerSession, error)
}


// ------------------------------------------------------------
// Handler
// ------------------------------------------------------------

type PrayerSessionHandler struct {
	create CreatePrayerSessionService
	update UpdatePrayerSessionService
	delete DeletePrayerSessionService
	list   ListPrayerSessionService
}

func NewPrayerSessionHandler(
	create CreatePrayerSessionService,
	update UpdatePrayerSessionService,
	delete DeletePrayerSessionService,
	list ListPrayerSessionService,
) *PrayerSessionHandler {
	return &PrayerSessionHandler{
		create: create,
		update: update,
		delete: delete,
		list:   list,
	}
}

// ------------------------------------------------------------
// Request DTOs
// ------------------------------------------------------------

type createPrayerSessionRequest struct {
	PrayerGroupID  string                    `json:"prayerGroupID"`
	Title          string                    `json:"title"`
	Date           string                    `json:"date"`
	Time           string                    `json:"time"`
	Duration       int                       `json:"duration"`
	PrayerPointIDs []domainid.PrayerPointID `json:"prayerPointIDs"`
}

type updatePrayerSessionRequest struct {
	Title          *string                   `json:"title"`
	Date           *string                   `json:"date"`
	Time           *string                   `json:"time"`
	Duration       *int                      `json:"duration"`
	PrayerPointIDs *[]domainid.PrayerPointID `json:"prayerPointIDs"`
}

func (h *PrayerSessionHandler) Create(
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

	var request createPrayerSessionRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]any{
				"error": "invalid request body",
			},
		)
		return
	}

	sessionDate, err := time.Parse("2006-01-02", request.Date)
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{
			"error": "invalid date",
		})
		return
	}

	result, err := h.create.Create(
		r.Context(),
		appprayersession.CreateCommand{
			ActorID:        domainid.UserID(claims.Subject),
			PrayerGroupID:  domainid.PrayerGroupID(request.PrayerGroupID),
			Title:          request.Title,
			Date:           sessionDate,
			Time:           request.Time,
			Duration:       request.Duration,
			PrayerPointIDs: request.PrayerPointIDs,
		},
	)

	if err != nil {
		switch {
		case errors.Is(
			err,
			appprayersession.ErrInvalidGroupID,
		):
			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]any{
					"error": err.Error(),
				},
			)

		case errors.Is(
			err,
			appprayersession.ErrTitleRequired,
		):
			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]any{
					"error": err.Error(),
				},
			)

		case errors.Is(
			err,
			appprayersession.ErrInvalidDate,
		):
			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]any{
					"error": err.Error(),
				},
			)

		case errors.Is(
			err,
			appprayersession.ErrTimeRequired,
		):
			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]any{
					"error": err.Error(),
				},
			)

		case errors.Is(
			err,
			appprayersession.ErrInvalidDuration,
		):
			writeJSON(
				w,
				http.StatusBadRequest,
				map[string]any{
					"error": err.Error(),
				},
			)

		case errors.Is(
			err,
			appprayersession.ErrUserNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "user not found",
				},
			)

		case errors.Is(
			err,
			appprayersession.ErrForbidden,
		):
			writeJSON(
				w,
				http.StatusForbidden,
				map[string]any{
					"error": "forbidden",
				},
			)

		default:
			writeJSON(
				w,
				http.StatusInternalServerError,
				map[string]any{
					"error": "internal server error",
				},
			)
		}

		return
	}

	writeJSON(
		w,
		http.StatusCreated,
		map[string]any{
			"prayerSession": result,
		},
	)
}

func (h *PrayerSessionHandler) Delete(
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

	sessionID := domainid.PrayerSessionID(
		r.PathValue("sessionID"),
	)

	if sessionID.String() == "" {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]any{
				"error": "sessionID is required",
			},
		)
		return
	}

	err := h.delete.Delete(
		r.Context(),
		domainid.UserID(claims.Subject),
		sessionID,
	)

	if err != nil {
		switch {
		case errors.Is(
			err,
			appprayersession.ErrNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "prayer session not found",
				},
			)

		case errors.Is(
			err,
			appprayersession.ErrUserNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "user not found",
				},
			)

		case errors.Is(
			err,
			appprayersession.ErrForbidden,
		):
			writeJSON(
				w,
				http.StatusForbidden,
				map[string]any{
					"error": "forbidden",
				},
			)

		default:
			writeJSON(
				w,
				http.StatusInternalServerError,
				map[string]any{
					"error": "internal server error",
				},
			)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PrayerSessionHandler) List(
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

	sessions, err := h.list.ListForUser(
		r.Context(),
		domainid.UserID(claims.Subject),
	)

	if err != nil {
		switch {
		case errors.Is(
			err,
			domainuser.ErrUserNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "user not found",
				},
			)

		default:
			writeJSON(
				w,
				http.StatusInternalServerError,
				map[string]any{
					"error": "internal server error",
				},
			)
		}

		return
	}

	writeJSON(
		w,
		http.StatusOK,
		map[string]any{
			"prayerSessions": sessions,
		},
	)
}