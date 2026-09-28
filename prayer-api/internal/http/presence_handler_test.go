
package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/clerk/clerk-sdk-go/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	apppresence "prayer-api/internal/application/presence"
	domain "prayer-api/internal/domain/presence"
	domainid "prayer-api/internal/domain/identity"
)

type mockHeartbeatService struct {
	executeFn func(
		context.Context,
		domainid.PrayerSessionID,
		domainid.UserID,
	) (*domain.Presence, error)
}

func (m *mockHeartbeatService) Execute(
	ctx context.Context,
	sessionID domainid.PrayerSessionID,
	userID domainid.UserID,
) (*domain.Presence, error) {
	return m.executeFn(ctx, sessionID, userID)
}

type mockListParticipantsService struct {
	executeFn func(
		context.Context,
		domainid.PrayerSessionID,
		domainid.UserID,
	) ([]domain.Presence, error)
}

func (m *mockListParticipantsService) Execute(
	ctx context.Context,
	sessionID domainid.PrayerSessionID,
	userID domainid.UserID,
) ([]domain.Presence, error) {
	return m.executeFn(ctx, sessionID, userID)
}

func presenceRequestWithClaims(
	method string,
	url string,
) *http.Request {
	req := httptest.NewRequest(method, url, nil)

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

func TestPresenceHandler_Heartbeat_Unauthorized(t *testing.T) {
	handler := &PresenceHandler{}

	req := httptest.NewRequest(
		http.MethodPost,
		"/api/prayer-sessions/session-1/heartbeat",
		nil,
	)
	req.SetPathValue("sessionID", "session-1")

	rec := httptest.NewRecorder()

	handler.Heartbeat(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestPresenceHandler_Heartbeat_MissingSessionID(t *testing.T) {
	mock := &mockHeartbeatService{
		executeFn: func(
			context.Context,
			domainid.PrayerSessionID,
			domainid.UserID,
		) (*domain.Presence, error) {
			t.Fatal("service should not be called")
			return nil, nil
		},
	}

	handler := &PresenceHandler{
		heartbeat: mock,
	}

	req := presenceRequestWithClaims(
		http.MethodPost,
		"/api/prayer-sessions//heartbeat",
	)
	req.SetPathValue("sessionID", "")

	rec := httptest.NewRecorder()

	handler.Heartbeat(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestPresenceHandler_Heartbeat_Success(t *testing.T) {
	lastSeen := time.Date(
		2026, time.September, 28,
		20, 30, 0, 0,
		time.UTC,
	)

	mock := &mockHeartbeatService{
		executeFn: func(
			ctx context.Context,
			sessionID domainid.PrayerSessionID,
			userID domainid.UserID,
		) (*domain.Presence, error) {
			assert.Equal(t, "session-1", string(sessionID))
			assert.Equal(t, "user-123", string(userID))

			return &domain.Presence{
				SessionID: "session-1",
				UserID:    "user-123",
				LastSeen:  lastSeen,
			}, nil
		},
	}

	handler := &PresenceHandler{
		heartbeat: mock,
	}

	req := presenceRequestWithClaims(
		http.MethodPost,
		"/api/prayer-sessions/session-1/heartbeat",
	)
	req.SetPathValue("sessionID", "session-1")

	rec := httptest.NewRecorder()

	handler.Heartbeat(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]any
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))

	assert.Equal(t, "session-1", response["sessionId"])
	assert.Equal(t, "active", response["status"])
	assert.NotEmpty(t, response["lastSeen"])
}

func TestPresenceHandler_Heartbeat_ServiceError(t *testing.T) {
	mock := &mockHeartbeatService{
		executeFn: func(
			context.Context,
			domainid.PrayerSessionID,
			domainid.UserID,
		) (*domain.Presence, error) {
			return nil, errors.New("database error")
		},
	}

	handler := &PresenceHandler{
		heartbeat: mock,
	}

	req := presenceRequestWithClaims(
		http.MethodPost,
		"/api/prayer-sessions/session-1/heartbeat",
	)
	req.SetPathValue("sessionID", "session-1")

	rec := httptest.NewRecorder()

	handler.Heartbeat(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestPresenceHandler_ListParticipants_Unauthorized(t *testing.T) {
	handler := &PresenceHandler{}

	req := httptest.NewRequest(
		http.MethodGet,
		"/api/prayer-sessions/session-1/participants",
		nil,
	)
	req.SetPathValue("sessionID", "session-1")

	rec := httptest.NewRecorder()

	handler.ListParticipants(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestPresenceHandler_ListParticipants_MissingSessionID(t *testing.T) {
	mock := &mockListParticipantsService{
		executeFn: func(
			context.Context,
			domainid.PrayerSessionID,
			domainid.UserID,
		) ([]domain.Presence, error) {
			t.Fatal("service should not be called")
			return nil, nil
		},
	}

	handler := &PresenceHandler{
		listParticipants: mock,
	}

	req := presenceRequestWithClaims(
		http.MethodGet,
		"/api/prayer-sessions//participants",
	)
	req.SetPathValue("sessionID", "")

	rec := httptest.NewRecorder()

	handler.ListParticipants(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestPresenceHandler_ListParticipants_Success(t *testing.T) {
	lastSeen1 := time.Date(
		2026, time.September, 28,
		20, 30, 0, 0,
		time.UTC,
	)
	lastSeen2 := lastSeen1.Add(-time.Minute)

	mock := &mockListParticipantsService{
		executeFn: func(
			ctx context.Context,
			sessionID domainid.PrayerSessionID,
			userID domainid.UserID,
		) ([]domain.Presence, error) {
			assert.Equal(t, "session-1", string(sessionID))
			assert.Equal(t, "user-123", string(userID))

			return []domain.Presence{
				{
					SessionID: "session-1",
					UserID:    "user-123",
					LastSeen:  lastSeen1,
				},
				{
					SessionID: "session-1",
					UserID:    "user-456",
					LastSeen:  lastSeen2,
				},
			}, nil
		},
	}

	handler := &PresenceHandler{
		listParticipants: mock,
	}

	req := presenceRequestWithClaims(
		http.MethodGet,
		"/api/prayer-sessions/session-1/participants",
	)
	req.SetPathValue("sessionID", "session-1")

	rec := httptest.NewRecorder()

	handler.ListParticipants(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

	var response map[string]any
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &response))

	participants, ok := response["participants"].([]any)
	require.True(t, ok)
	require.Len(t, participants, 2)

	first, ok := participants[0].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "user-123", first["userId"])

	second, ok := participants[1].(map[string]any)
	require.True(t, ok)
	assert.Equal(t, "user-456", second["userId"])
}

func TestPresenceHandler_ListParticipants_ServiceError(t *testing.T) {
	mock := &mockListParticipantsService{
		executeFn: func(
			context.Context,
			domainid.PrayerSessionID,
			domainid.UserID,
		) ([]domain.Presence, error) {
			return nil, errors.New("database error")
		},
	}

	handler := &PresenceHandler{
		listParticipants: mock,
	}

	req := presenceRequestWithClaims(
		http.MethodGet,
		"/api/prayer-sessions/session-1/participants",
	)
	req.SetPathValue("sessionID", "session-1")

	rec := httptest.NewRecorder()

	handler.ListParticipants(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)
}