package httpapi

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"status": "ok",
		})
	})

	request := httptest.NewRequest(
		http.MethodGet,
		"/api/health",
		nil,
	)

	response := httptest.NewRecorder()

	mux.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", response.Code)
	}

	expected := "{\"status\":\"ok\"}\n"

	if response.Body.String() != expected {
		t.Fatalf("expected %q, got %q", expected, response.Body.String())
	}
}
