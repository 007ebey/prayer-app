package httpapi

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
    "io"
	appprayersession "prayer-api/internal/application/prayersession"
	domainid "prayer-api/internal/domain/identity"
	domain "prayer-api/internal/domain/prayersession"
	domainuser "prayer-api/internal/domain/user"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ------------------------------------------------------------
// Create service mock
// ------------------------------------------------------------

type createPrayerSessionServiceMock struct {
	createFn func(
		ctx context.Context,
		cmd appprayersession.CreateCommand,
	) (*domain.PrayerSession, error)

	called bool
	cmd    appprayersession.CreateCommand
}

func (m *createPrayerSessionServiceMock) Create(
	ctx context.Context,
	cmd appprayersession.CreateCommand,
) (*domain.PrayerSession, error) {
	m.called = true
	m.cmd = cmd

	if m.createFn != nil {
		return m.createFn(ctx, cmd)
	}

	return nil, nil
}

// ------------------------------------------------------------
// Update service mock
// ------------------------------------------------------------

type updatePrayerSessionServiceMock struct {
	updateFn func(
		ctx context.Context,
		cmd appprayersession.UpdateCommand,
	) (*domain.PrayerSession, error)

	called bool
	cmd    appprayersession.UpdateCommand
}

func (m *updatePrayerSessionServiceMock) Update(
	ctx context.Context,
	cmd appprayersession.UpdateCommand,
) (*domain.PrayerSession, error) {
	m.called = true
	m.cmd = cmd

	if m.updateFn != nil {
		return m.updateFn(ctx, cmd)
	}

	return nil, nil
}

// ------------------------------------------------------------
// Delete service mock
// ------------------------------------------------------------

type deletePrayerSessionServiceMock struct {
	deleteFn func(
		ctx context.Context,
		userID domainid.UserID,
		sessionID domainid.PrayerSessionID,
	) error

	called   bool
	userID   domainid.UserID
	sessionID domainid.PrayerSessionID
}

func (m *deletePrayerSessionServiceMock) Delete(
	ctx context.Context,
	userID domainid.UserID,
	sessionID domainid.PrayerSessionID,
) error {
	m.called = true
	m.userID = userID
	m.sessionID = sessionID

	if m.deleteFn != nil {
		return m.deleteFn(ctx, userID, sessionID)
	}

	return nil
}

// ------------------------------------------------------------
// List service mock
// ------------------------------------------------------------

type listPrayerSessionServiceMock struct {
	listFn func(
		ctx context.Context,
		userID domainid.UserID,
	) ([]domain.PrayerSession, error)

	called bool
	userID domainid.UserID
}

func (m *listPrayerSessionServiceMock) ListForUser(
	ctx context.Context,
	userID domainid.UserID,
) ([]domain.PrayerSession, error) {
	m.called = true
	m.userID = userID

	if m.listFn != nil {
		return m.listFn(ctx, userID)
	}

	return nil, nil
}

// ------------------------------------------------------------
// Helpers
// ------------------------------------------------------------

// Reuses the authenticatedRequest helper already defined in
// auth_handler_test.go.
//
// That shared helper is responsible for putting the Clerk
// authentication context into the request.
//
// We only need to change method/body for prayer-session tests.
func authenticatedPrayerSessionRequest(
	method string,
	target string,
	body []byte,
) *http.Request {
	req := authenticatedRequest("user-123")

	req.Method = method
	req.URL.Path = target

	if body != nil {
		req.Body = io.NopCloser(bytes.NewReader(body))
		req.ContentLength = int64(len(body))
	}

	return req
}
// Small local adapter so we don't need to duplicate Clerk auth logic.
type nopCloser struct {
	*bytes.Reader
}

func (n nopCloser) Close() error {
	return nil
}

func ioNopCloser(r *bytes.Reader) nopCloser {
	return nopCloser{Reader: r}
}

func decodeJSONBody(
	t *testing.T,
	rec *httptest.ResponseRecorder,
) map[string]any {
	t.Helper()

	var body map[string]any

	err := json.NewDecoder(rec.Body).Decode(&body)
	require.NoError(t, err)

	return body
}

// ------------------------------------------------------------
// Create
// ------------------------------------------------------------

func TestPrayerSessionHandler_Create_Unauthorized(t *testing.T) {
	createMock := &createPrayerSessionServiceMock{}

	handler := NewPrayerSessionHandler(
		createMock,
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		&listPrayerSessionServiceMock{},
	)

	req := httptest.NewRequest(
		http.MethodPost,
		"/prayer-sessions",
		bytes.NewBufferString(`{}`),
	)

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)

	body := decodeJSONBody(t, rec)

	assert.Equal(t, "unauthorized", body["error"])
	assert.False(t, createMock.called)
}

func TestPrayerSessionHandler_Create_InvalidJSON(t *testing.T) {
	createMock := &createPrayerSessionServiceMock{}

	handler := NewPrayerSessionHandler(
		createMock,
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		&listPrayerSessionServiceMock{},
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodPost,
		"/prayer-sessions",
		[]byte(`{"title":`),
	)

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	body := decodeJSONBody(t, rec)

	assert.Equal(t, "invalid request body", body["error"])
	assert.False(t, createMock.called)
}

func TestPrayerSessionHandler_Create_InvalidDate(t *testing.T) {
	createMock := &createPrayerSessionServiceMock{}

	handler := NewPrayerSessionHandler(
		createMock,
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		&listPrayerSessionServiceMock{},
	)

	body := []byte(`{
		"prayerGroupID": "group-1",
		"title": "Morning Prayer",
		"date": "not-a-date",
		"time": "07:00",
		"duration": 30
	}`)

	req := authenticatedPrayerSessionRequest(
		http.MethodPost,
		"/prayer-sessions",
		body,
	)

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "invalid date", response["error"])
	assert.False(t, createMock.called)
}

func TestPrayerSessionHandler_Create_Success(t *testing.T) {
	expected := &domain.PrayerSession{
		ID:            domainid.PrayerSessionID("session-1"),
		PrayerGroupID: domainid.PrayerGroupID("group-1"),
		Title:         "Morning Prayer",
		Date:          time.Date(2026, 9, 20, 0, 0, 0, 0, time.UTC),
		Time:          "07:00",
		Duration:      30,
		PrayerPointIDs: []domainid.PrayerPointID{
			domainid.PrayerPointID("point-1"),
			domainid.PrayerPointID("point-2"),
		},
	}

	createMock := &createPrayerSessionServiceMock{
		createFn: func(
			ctx context.Context,
			cmd appprayersession.CreateCommand,
		) (*domain.PrayerSession, error) {
			return expected, nil
		},
	}

	handler := NewPrayerSessionHandler(
		createMock,
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		&listPrayerSessionServiceMock{},
	)

	body := []byte(`{
		"prayerGroupID": "group-1",
		"title": " Morning Prayer ",
		"date": "2026-09-20",
		"time": "07:00",
		"duration": 30,
		"prayerPointIDs": ["point-1", "point-2"]
	}`)

	req := authenticatedPrayerSessionRequest(
		http.MethodPost,
		"/prayer-sessions",
		body,
	)

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.True(t, createMock.called)

	assert.Equal(
		t,
		domainid.UserID("user-123"),
		createMock.cmd.ActorID,
	)

	assert.Equal(
		t,
		domainid.PrayerGroupID("group-1"),
		createMock.cmd.PrayerGroupID,
	)

	assert.Equal(t, " Morning Prayer ", createMock.cmd.Title)
	assert.Equal(t, "07:00", createMock.cmd.Time)
	assert.Equal(t, 30, createMock.cmd.Duration)

	assert.Equal(
		t,
		time.Date(2026, 9, 20, 0, 0, 0, 0, time.UTC),
		createMock.cmd.Date,
	)

	assert.Equal(
		t,
		[]domainid.PrayerPointID{
			domainid.PrayerPointID("point-1"),
			domainid.PrayerPointID("point-2"),
		},
		createMock.cmd.PrayerPointIDs,
	)

	response := decodeJSONBody(t, rec)

	assert.NotNil(t, response["prayerSession"])
}

func TestPrayerSessionHandler_Create_UserNotFound(t *testing.T) {
	createMock := &createPrayerSessionServiceMock{
		createFn: func(
			ctx context.Context,
			cmd appprayersession.CreateCommand,
		) (*domain.PrayerSession, error) {
			return nil, appprayersession.ErrUserNotFound
		},
	}

	handler := NewPrayerSessionHandler(
		createMock,
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		&listPrayerSessionServiceMock{},
	)

	body := []byte(`{
		"prayerGroupID": "group-1",
		"title": "Morning Prayer",
		"date": "2026-09-20",
		"time": "07:00",
		"duration": 30
	}`)

	req := authenticatedPrayerSessionRequest(
		http.MethodPost,
		"/prayer-sessions",
		body,
	)

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "user not found", response["error"])
}

func TestPrayerSessionHandler_Create_Forbidden(t *testing.T) {
	createMock := &createPrayerSessionServiceMock{
		createFn: func(
			ctx context.Context,
			cmd appprayersession.CreateCommand,
		) (*domain.PrayerSession, error) {
			return nil, appprayersession.ErrForbidden
		},
	}

	handler := NewPrayerSessionHandler(
		createMock,
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		&listPrayerSessionServiceMock{},
	)

	body := []byte(`{
		"prayerGroupID": "group-1",
		"title": "Morning Prayer",
		"date": "2026-09-20",
		"time": "07:00",
		"duration": 30
	}`)

	req := authenticatedPrayerSessionRequest(
		http.MethodPost,
		"/prayer-sessions",
		body,
	)

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusForbidden, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "forbidden", response["error"])
}

func TestPrayerSessionHandler_Create_InternalError(t *testing.T) {
	createMock := &createPrayerSessionServiceMock{
		createFn: func(
			ctx context.Context,
			cmd appprayersession.CreateCommand,
		) (*domain.PrayerSession, error) {
			return nil, errors.New("repository failure")
		},
	}

	handler := NewPrayerSessionHandler(
		createMock,
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		&listPrayerSessionServiceMock{},
	)

	body := []byte(`{
		"prayerGroupID": "group-1",
		"title": "Morning Prayer",
		"date": "2026-09-20",
		"time": "07:00",
		"duration": 30
	}`)

	req := authenticatedPrayerSessionRequest(
		http.MethodPost,
		"/prayer-sessions",
		body,
	)

	rec := httptest.NewRecorder()

	handler.Create(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "internal server error", response["error"])
}

// ------------------------------------------------------------
// Delete
// ------------------------------------------------------------

func TestPrayerSessionHandler_Delete_Unauthorized(t *testing.T) {
	deleteMock := &deletePrayerSessionServiceMock{}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		deleteMock,
		&listPrayerSessionServiceMock{},
	)

	req := httptest.NewRequest(
		http.MethodDelete,
		"/prayer-sessions/session-123",
		nil,
	)

	req.SetPathValue("sessionID", "session-123")

	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, deleteMock.called)
}

func TestPrayerSessionHandler_Delete_MissingSessionID(t *testing.T) {
	deleteMock := &deletePrayerSessionServiceMock{}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		deleteMock,
		&listPrayerSessionServiceMock{},
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodDelete,
		"/prayer-sessions/",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusBadRequest, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "sessionID is required", response["error"])
	assert.False(t, deleteMock.called)
}

func TestPrayerSessionHandler_Delete_Success(t *testing.T) {
	deleteMock := &deletePrayerSessionServiceMock{}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		deleteMock,
		&listPrayerSessionServiceMock{},
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodDelete,
		"/prayer-sessions/session-123",
		nil,
	)

	req.SetPathValue("sessionID", "session-123")

	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusNoContent, rec.Code)
	assert.True(t, deleteMock.called)

	assert.Equal(
		t,
		domainid.UserID("user-123"),
		deleteMock.userID,
	)

	assert.Equal(
		t,
		domainid.PrayerSessionID("session-123"),
		deleteMock.sessionID,
	)

	assert.Empty(t, rec.Body.String())
}

func TestPrayerSessionHandler_Delete_NotFound(t *testing.T) {
	deleteMock := &deletePrayerSessionServiceMock{
		deleteFn: func(
			ctx context.Context,
			userID domainid.UserID,
			sessionID domainid.PrayerSessionID,
		) error {
			return appprayersession.ErrNotFound
		},
	}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		deleteMock,
		&listPrayerSessionServiceMock{},
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodDelete,
		"/prayer-sessions/session-123",
		nil,
	)

	req.SetPathValue("sessionID", "session-123")

	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "prayer session not found", response["error"])
}

func TestPrayerSessionHandler_Delete_UserNotFound(t *testing.T) {
	deleteMock := &deletePrayerSessionServiceMock{
		deleteFn: func(
			ctx context.Context,
			userID domainid.UserID,
			sessionID domainid.PrayerSessionID,
		) error {
			return appprayersession.ErrUserNotFound
		},
	}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		deleteMock,
		&listPrayerSessionServiceMock{},
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodDelete,
		"/prayer-sessions/session-123",
		nil,
	)

	req.SetPathValue("sessionID", "session-123")

	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "user not found", response["error"])
}

func TestPrayerSessionHandler_Delete_Forbidden(t *testing.T) {
	deleteMock := &deletePrayerSessionServiceMock{
		deleteFn: func(
			ctx context.Context,
			userID domainid.UserID,
			sessionID domainid.PrayerSessionID,
		) error {
			return appprayersession.ErrForbidden
		},
	}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		deleteMock,
		&listPrayerSessionServiceMock{},
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodDelete,
		"/prayer-sessions/session-123",
		nil,
	)

	req.SetPathValue("sessionID", "session-123")

	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusForbidden, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "forbidden", response["error"])
}

func TestPrayerSessionHandler_Delete_InternalError(t *testing.T) {
	deleteMock := &deletePrayerSessionServiceMock{
		deleteFn: func(
			ctx context.Context,
			userID domainid.UserID,
			sessionID domainid.PrayerSessionID,
		) error {
			return errors.New("database failure")
		},
	}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		deleteMock,
		&listPrayerSessionServiceMock{},
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodDelete,
		"/prayer-sessions/session-123",
		nil,
	)

	req.SetPathValue("sessionID", "session-123")

	rec := httptest.NewRecorder()

	handler.Delete(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "internal server error", response["error"])
}

// ------------------------------------------------------------
// List
// ------------------------------------------------------------

func TestPrayerSessionHandler_List_Unauthorized(t *testing.T) {
	listMock := &listPrayerSessionServiceMock{}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		listMock,
	)

	req := httptest.NewRequest(
		http.MethodGet,
		"/prayer-sessions",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.List(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
	assert.False(t, listMock.called)
}

func TestPrayerSessionHandler_List_Success(t *testing.T) {
	expectedSessions := []domain.PrayerSession{
		{
			ID:            domainid.PrayerSessionID("session-1"),
			PrayerGroupID: domainid.PrayerGroupID("group-1"),
			Title:         "Morning Prayer",
			Date:          time.Date(2026, 9, 20, 0, 0, 0, 0, time.UTC),
			Time:          "07:00",
			Duration:      30,
		},
		{
			ID:            domainid.PrayerSessionID("session-2"),
			PrayerGroupID: domainid.PrayerGroupID("group-2"),
			Title:         "Evening Prayer",
			Date:          time.Date(2026, 9, 20, 0, 0, 0, 0, time.UTC),
			Time:          "19:30",
			Duration:      60,
		},
	}

	listMock := &listPrayerSessionServiceMock{
		listFn: func(
			ctx context.Context,
			userID domainid.UserID,
		) ([]domain.PrayerSession, error) {
			return expectedSessions, nil
		},
	}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		listMock,
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodGet,
		"/prayer-sessions",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.List(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.True(t, listMock.called)

	assert.Equal(
		t,
		domainid.UserID("user-123"),
		listMock.userID,
	)

	response := decodeJSONBody(t, rec)

	assert.NotNil(t, response["prayerSessions"])

	sessions, ok := response["prayerSessions"].([]any)
	require.True(t, ok)

	assert.Len(t, sessions, 2)
}

func TestPrayerSessionHandler_List_UserNotFound(t *testing.T) {
	listMock := &listPrayerSessionServiceMock{
		listFn: func(
			ctx context.Context,
			userID domainid.UserID,
		) ([]domain.PrayerSession, error) {
			return nil, domainuser.ErrUserNotFound
		},
	}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		listMock,
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodGet,
		"/prayer-sessions",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.List(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "user not found", response["error"])
}

func TestPrayerSessionHandler_List_InternalError(t *testing.T) {
	listMock := &listPrayerSessionServiceMock{
		listFn: func(
			ctx context.Context,
			userID domainid.UserID,
		) ([]domain.PrayerSession, error) {
			return nil, errors.New("repository failure")
		},
	}

	handler := NewPrayerSessionHandler(
		&createPrayerSessionServiceMock{},
		&updatePrayerSessionServiceMock{},
		&deletePrayerSessionServiceMock{},
		listMock,
	)

	req := authenticatedPrayerSessionRequest(
		http.MethodGet,
		"/prayer-sessions",
		nil,
	)

	rec := httptest.NewRecorder()

	handler.List(rec, req)

	assert.Equal(t, http.StatusInternalServerError, rec.Code)

	response := decodeJSONBody(t, rec)

	assert.Equal(t, "internal server error", response["error"])
}
