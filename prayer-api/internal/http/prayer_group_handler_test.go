package httpapi

import (
	
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	appprayergroup "prayer-api/internal/application/prayergroup"
	domain "prayer-api/internal/domain/prayergroup"
	domainuser "prayer-api/internal/domain/user"
)

type mockCreateService struct {
	createFn func(
		context.Context,
		appprayergroup.CreateCommand,
	) (*appprayergroup.CreateResult, error)
}

func (m *mockCreateService) Create(
	ctx context.Context,
	cmd appprayergroup.CreateCommand,
) (*appprayergroup.CreateResult, error) {
	return m.createFn(ctx, cmd)
}

type mockBlockService struct {
	blockFn func(
		context.Context,
		appprayergroup.BlockCommand,
	) error
}

func (m *mockBlockService) Block(
	ctx context.Context,
	cmd appprayergroup.BlockCommand,
) error {
	return m.blockFn(ctx, cmd)
}

func requestWithClaims(body string) *http.Request {
	req := httptest.NewRequest(
		http.MethodPost,
		"/groups",
		bytes.NewBufferString(body),
	)

    claims := &clerk.SessionClaims{
        RegisteredClaims: clerk.RegisteredClaims{
        	Subject: "user-123",
        },
    }

	ctx := clerk.ContextWithSessionClaims(
		req.Context(),
		claims,
	)

	return req.WithContext(ctx)
}

func TestPrayerGroupHandler_Create_Unauthorized(t *testing.T) {
	handler := &PrayerGroupHandler{}

	req := httptest.NewRequest(
		http.MethodPost,
		"/groups",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.Create(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Create_InvalidJSON(t *testing.T) {

	handler := &PrayerGroupHandler{}

	req := requestWithClaims("{")
	rec := httptest.NewRecorder()

	handler.Create(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusBadRequest,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Create_Success(t *testing.T) {

	mock := &mockCreateService{
		createFn: func(
			ctx context.Context,
			cmd appprayergroup.CreateCommand,
		) (*appprayergroup.CreateResult, error) {

			assert.Equal(t, "user-123", cmd.ActorExternalID)
			assert.Equal(t, "Young Adults", cmd.Name)
			assert.Equal(t, "Description", cmd.Description)

			return &appprayergroup.CreateResult{
				PrayerGroup: &domain.PrayerGroup{
					ID:          "group-1",
					Name:        "Young Adults",
					Description: "Description",
					Status:      domain.StatusActive,
				},
			}, nil
		},
	}

	handler := &PrayerGroupHandler{
		create: mock,
	}

	body := `{
		"name":"Young Adults",
		"description":"Description"
	}`

	req := requestWithClaims(body)
	rec := httptest.NewRecorder()

	handler.Create(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusCreated,
		rec.Code,
	)

	var response map[string]any

	require.NoError(
		t,
		json.Unmarshal(
			rec.Body.Bytes(),
			&response,
		),
	)

	group := response["prayerGroup"].(map[string]any)

	assert.Equal(
		t,
		"group-1",
		group["id"],
	)
}

func TestPrayerGroupHandler_Create_ServiceUnauthorized(t *testing.T) {

	mock := &mockCreateService{
		createFn: func(
			context.Context,
			appprayergroup.CreateCommand,
		) (*appprayergroup.CreateResult, error) {
			return nil, appprayergroup.ErrUnauthorized
		},
	}

	handler := &PrayerGroupHandler{
		create: mock,
	}

	req := requestWithClaims(`{"name":"A"}`)
	rec := httptest.NewRecorder()

	handler.Create(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Create_ActorNotFound(t *testing.T) {

	mock := &mockCreateService{
		createFn: func(
			context.Context,
			appprayergroup.CreateCommand,
		) (*appprayergroup.CreateResult, error) {
			return nil, appprayergroup.ErrActorNotFound
		},
	}

	handler := &PrayerGroupHandler{
		create: mock,
	}

	req := requestWithClaims(`{"name":"A"}`)
	rec := httptest.NewRecorder()

	handler.Create(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusNotFound,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Create_Forbidden(t *testing.T) {

	mock := &mockCreateService{
		createFn: func(
			context.Context,
			appprayergroup.CreateCommand,
		) (*appprayergroup.CreateResult, error) {
			return nil, appprayergroup.ErrForbidden
		},
	}

	handler := &PrayerGroupHandler{
		create: mock,
	}

	req := requestWithClaims(`{"name":"A"}`)
	rec := httptest.NewRecorder()

	handler.Create(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusForbidden,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Create_Duplicate(t *testing.T) {

	mock := &mockCreateService{
		createFn: func(
			context.Context,
			appprayergroup.CreateCommand,
		) (*appprayergroup.CreateResult, error) {
			return nil, appprayergroup.ErrPrayerGroupExists
		},
	}

	handler := &PrayerGroupHandler{
		create: mock,
	}

	req := requestWithClaims(`{"name":"A"}`)
	rec := httptest.NewRecorder()

	handler.Create(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusConflict,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Create_GenericError(t *testing.T) {

	mock := &mockCreateService{
		createFn: func(
			context.Context,
			appprayergroup.CreateCommand,
		) (*appprayergroup.CreateResult, error) {
			return nil, errors.New("validation failed")
		},
	}

	handler := &PrayerGroupHandler{
		create: mock,
	}

	req := requestWithClaims(`{"name":"A"}`)
	rec := httptest.NewRecorder()

	handler.Create(
		rec,
		req,
	)

	assert.Equal(
		t,
		http.StatusBadRequest,
		rec.Code,
	)

	assert.Contains(
		t,
		rec.Body.String(),
		"validation failed",
	)
}

func TestPrayerGroupHandler_Block_Unauthorized(t *testing.T) {
	handler := &PrayerGroupHandler{}

	req := httptest.NewRequest(
		http.MethodPatch,
		"/groups/group-1/users/user-1/block",
		nil,
	)

	req.SetPathValue("groupID", "group-1")
	req.SetPathValue("userID", "user-1")

	rec := httptest.NewRecorder()

	handler.BlockPrayerGroup(rec, req)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Block_Success(t *testing.T) {
	mock := &mockBlockService{
		blockFn: func(
			ctx context.Context,
			cmd appprayergroup.BlockCommand,
		) error {
			assert.Equal(
				t,
				"user-123",
				cmd.ActorExternalID,
			)

			assert.Equal(
				t,
				"group-1",
				string(cmd.GroupID),
			)

			assert.Equal(
				t,
				"user-1",
				string(cmd.UserID),
			)

			return nil
		},
	}

	handler := &PrayerGroupHandler{
		block: mock,
	}

	req := requestWithClaims("")
	req.Method = http.MethodPatch

	req.SetPathValue(
		"groupID",
		"group-1",
	)

	req.SetPathValue(
		"userID",
		"user-1",
	)

	rec := httptest.NewRecorder()

	handler.BlockPrayerGroup(rec, req)

	assert.Equal(
		t,
		http.StatusNoContent,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Block_ServiceUnauthorized(t *testing.T) {
	mock := &mockBlockService{
		blockFn: func(
			context.Context,
			appprayergroup.BlockCommand,
		) error {
			return appprayergroup.ErrUnauthorized
		},
	}

	handler := &PrayerGroupHandler{
		block: mock,
	}

	req := requestWithClaims("")
	req.SetPathValue("groupID", "group-1")
	req.SetPathValue("userID", "user-1")

	rec := httptest.NewRecorder()

	handler.BlockPrayerGroup(rec, req)

	assert.Equal(
		t,
		http.StatusUnauthorized,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Block_ActorNotFound(t *testing.T) {
	mock := &mockBlockService{
		blockFn: func(
			context.Context,
			appprayergroup.BlockCommand,
		) error {
			return appprayergroup.ErrActorNotFound
		},
	}

	handler := &PrayerGroupHandler{
		block: mock,
	}

	req := requestWithClaims("")
	req.SetPathValue("groupID", "group-1")
	req.SetPathValue("userID", "user-1")

	rec := httptest.NewRecorder()

	handler.BlockPrayerGroup(rec, req)

	assert.Equal(
		t,
		http.StatusNotFound,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Block_Forbidden(t *testing.T) {
	mock := &mockBlockService{
		blockFn: func(
			context.Context,
			appprayergroup.BlockCommand,
		) error {
			return appprayergroup.ErrForbidden
		},
	}

	handler := &PrayerGroupHandler{
		block: mock,
	}

	req := requestWithClaims("")
	req.SetPathValue("groupID", "group-1")
	req.SetPathValue("userID", "user-1")

	rec := httptest.NewRecorder()

	handler.BlockPrayerGroup(rec, req)

	assert.Equal(
		t,
		http.StatusForbidden,
		rec.Code,
	)
}

func TestPrayerGroupHandler_Block_UserNotFound(t *testing.T) {
	mock := &mockBlockService{
		blockFn: func(
			context.Context,
			appprayergroup.BlockCommand,
		) error {
			return domainuser.ErrUserNotFound
		},
	}

	handler := &PrayerGroupHandler{
		block: mock,
	}

	req := requestWithClaims("")
	req.SetPathValue("groupID", "group-1")
	req.SetPathValue("userID", "user-1")

	rec := httptest.NewRecorder()

	handler.BlockPrayerGroup(rec, req)

	assert.Equal(
		t,
		http.StatusNotFound,
		rec.Code,
	)
}