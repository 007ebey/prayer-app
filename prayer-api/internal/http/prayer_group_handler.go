package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	domainuser "prayer-api/internal/domain/user"
	domainprayergroup "prayer-api/internal/domain/prayergroup"
    domainid "prayer-api/internal/domain/identity"
	appprayergroup "prayer-api/internal/application/prayergroup"
)

type CreatePrayerGroupService interface {
    Create(
        ctx context.Context,
        cmd appprayergroup.CreateCommand,
    ) (*appprayergroup.CreateResult, error)
}

type ListPrayerGroupService interface {
	List(
		ctx context.Context,
		cmd appprayergroup.ListQuery,
	) (*appprayergroup.ListResult, error)
}

type GetPrayerGroupService interface {
	Get(
		ctx context.Context,
		query appprayergroup.GetQuery,
	) (*appprayergroup.GetResult, error)
}	

type UpdatePrayerGroupService interface {
	Update(
		ctx context.Context,
		req appprayergroup.UpdateRequest,
	) (*domainprayergroup.PrayerGroup, error)
}

type DeletePrayerGroupService interface {
	Delete(
		ctx context.Context,
		actorID domainid.UserID,
		groupID domainid.PrayerGroupID,
	) error
}

type AssignPrayerGroupService interface {
	Assign(
		ctx context.Context,
		cmd appprayergroup.AssignCommand,
	) error
}

type RemovePrayerGroupService interface {
	Remove(
		ctx context.Context,
		cmd appprayergroup.RemoveCommand,
	) error
}

type BlockPrayerGroupService interface {
	Block(
	    ctx context.Context,
	    cmd appprayergroup.BlockCommand,
    ) error
}

type PrayerGroupHandler struct {
	create CreatePrayerGroupService
	list   ListPrayerGroupService
	get    GetPrayerGroupService
	update UpdatePrayerGroupService
	delete DeletePrayerGroupService
	assign AssignPrayerGroupService
	remove RemovePrayerGroupService
	block  BlockPrayerGroupService
}

func NewPrayerGroupHandler(
	create CreatePrayerGroupService,
	list ListPrayerGroupService,
	get GetPrayerGroupService,
	update UpdatePrayerGroupService,
	delete DeletePrayerGroupService,
	assign AssignPrayerGroupService,
	remove RemovePrayerGroupService,
	block  BlockPrayerGroupService,
) *PrayerGroupHandler {
	return &PrayerGroupHandler{
		create: create,
		list:   list,
		get:    get,
		update: update,
		delete: delete,
		assign: assign,
		remove: remove,
		block:  block,
	}
}

type createPrayerGroupRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type updatePrayerGroupRequest struct {
	Name        *string `json:"name"`
	Description *string `json:"description"`
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

func (h *PrayerGroupHandler) Get(
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

	result, err := h.list.List(
		r.Context(),
		appprayergroup.ListQuery{
			ActorExternalID: claims.Subject,
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

		default:
			writeJSON(
				w,
				http.StatusInternalServerError,
				map[string]any{
					"error": err.Error(),
				},
			)
		}

		return
	}

	prayerGroups := make(
		[]map[string]any,
		0,
		len(result.PrayerGroups),
	)

	for _, group := range result.PrayerGroups {
		prayerGroups = append(
			prayerGroups,
			map[string]any{
				"id":          group.ID,
				"name":        group.Name,
				"description": group.Description,
				"status":      group.Status,
			},
		)
	}

	writeJSON(
		w,
		http.StatusOK,
		map[string]any{
			"prayerGroups": prayerGroups,
		},
	)
}

func (h *PrayerGroupHandler) GetByID(
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

	groupID := r.PathValue("groupID")

	result, err := h.get.Get(
		r.Context(),
		appprayergroup.GetQuery{
			ActorExternalID: claims.Subject,
			GroupID:         domainid.PrayerGroupID(groupID),
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
			appprayergroup.ErrPrayerGroupNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "prayer group not found",
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

		default:
			writeJSON(
				w,
				http.StatusInternalServerError,
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
		http.StatusOK,
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

func (h *PrayerGroupHandler) Update(
	w http.ResponseWriter,
	r *http.Request,
) {
	_, ok := clerk.SessionClaimsFromContext(
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

	groupID := r.PathValue("groupID")

	var request updatePrayerGroupRequest

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

	group, err := h.update.Update(
		r.Context(),
		appprayergroup.UpdateRequest{
			GroupID:     domainid.PrayerGroupID(groupID),
			Name:        request.Name,
			Description: request.Description,
		},
	)

	if err != nil {
		switch {
		case errors.Is(
			err,
			appprayergroup.ErrPrayerGroupNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "prayer group not found",
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

	writeJSON(
		w,
		http.StatusOK,
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

func (h *PrayerGroupHandler) Delete(
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

	groupID := domainid.PrayerGroupID(
		r.PathValue("groupID"),
	)

	err := h.delete.Delete(
		r.Context(),
		domainid.UserID(claims.Subject),
		groupID,
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
			appprayergroup.ErrPrayerGroupNotFound,
		):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "prayer group not found",
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

		default:
			writeJSON(
				w,
				http.StatusInternalServerError,
				map[string]any{
					"error": err.Error(),
				},
			)
		}

		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PrayerGroupHandler) AssignPrayerGroup(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := r.PathValue("userID")
	groupID := r.PathValue("groupID")

	if userID == "" || groupID == "" {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]any{
				"error": "userID and groupID are required",
			},
		)
		return
	}
 
	if err := h.assign.Assign(
		r.Context(),
		appprayergroup.AssignCommand{
			UserID: domainid.UserID(userID),
			GroupID: domainid.PrayerGroupID(groupID),
		},
	); err != nil {
		switch {
		case errors.Is(err, domainuser.ErrUserNotFound):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "User not found.",
				},
			)

		case errors.Is(err, domainprayergroup.ErrPrayerGroupNotFound):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "Prayer group not found.",
				},
			)

		case errors.Is(err, domainprayergroup.ErrPrayerGroupAlreadyAssigned):
			writeJSON(
				w,
				http.StatusConflict,
				map[string]any{
					"error": "User already belongs to the prayer group.",
				},
			)
		default:
			writeJSON(
				w,
				http.StatusInternalServerError,
				map[string]any{
					"error": "Internal server error.",
				},
			)
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PrayerGroupHandler) RemovePrayerGroup(
	w http.ResponseWriter,
	r *http.Request,
) {
	userID := r.PathValue("userID")
	groupID := r.PathValue("groupID")

	if userID == "" || groupID == "" {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]any{
               "error": "userID and groupID are required",
			},
		)
	}

	if err := h.remove.Remove(
        r.Context(),
        appprayergroup.RemoveCommand{
            UserID:  domainid.UserID(userID),
            GroupID: domainid.PrayerGroupID(groupID),
        },
    ); err != nil {
		switch {
		case errors.Is(err, domainuser.ErrUserNotFound):
            writeJSON(
                w,
                http.StatusNotFound,
                map[string]any{
                    "error": "User not found.",
                },
            )
		case errors.Is(err, domainprayergroup.ErrPrayerGroupNotFound):
            writeJSON(
                w,
                http.StatusNotFound,
                map[string]any{
                    "error": "Prayer group not found.",
                },
            )
	    case errors.Is(err, domainprayergroup.ErrPrayerGroupNotAssigned):
            writeJSON(
                w,
                http.StatusConflict,
                map[string]any{
                    "error": "User is not assigned to the prayer group.",
                },
            )
        default:
            writeJSON(
                w,
                http.StatusInternalServerError,
                map[string]any{
                    "error": "Internal server error.",
                },
            )
        }
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *PrayerGroupHandler) BlockPrayerGroup(
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

	userID := r.PathValue("userID")
	groupID := r.PathValue("groupID")

	if userID == "" || groupID == "" {
		writeJSON(
			w,
			http.StatusBadRequest,
			map[string]any{
				"error": "userID and groupID are required",
			},
		)
		return
	}

	err := h.block.Block(
		r.Context(),
		appprayergroup.BlockCommand{
			ActorExternalID: claims.Subject,
			UserID:          domainid.UserID(userID),
			GroupID:         domainid.PrayerGroupID(groupID),
		},
	)

	if err != nil {
		switch {
		case errors.Is(err, appprayergroup.ErrUnauthorized):
			writeJSON(
				w,
				http.StatusUnauthorized,
				map[string]any{
					"error": "unauthorized",
				},
			)

		case errors.Is(err, appprayergroup.ErrActorNotFound):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "actor not found",
				},
			)

		case errors.Is(err, appprayergroup.ErrForbidden):
			writeJSON(
				w,
				http.StatusForbidden,
				map[string]any{
					"error": "forbidden",
				},
			)

		case errors.Is(err, domainuser.ErrUserNotFound):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "user not found",
				},
			)

		case errors.Is(err, domainprayergroup.ErrPrayerGroupNotFound):
			writeJSON(
				w,
				http.StatusNotFound,
				map[string]any{
					"error": "prayer group not found",
				},
			)

		case errors.Is(err, domainprayergroup.ErrPrayerGroupNotAssigned):
			writeJSON(
				w,
				http.StatusConflict,
				map[string]any{
					"error": "user is not assigned to the prayer group",
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