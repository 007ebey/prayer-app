package httpapi

import (
	"net/http"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

func NewRouter(auth *AuthHandler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(
		"POST /api/auth/login",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(auth.Login),
		),
	)

	mux.HandleFunc(
		"GET /api/health",
		func(w http.ResponseWriter, r *http.Request) {
			writeJSON(w, http.StatusOK, map[string]any{
				"status": "ok",
			})
		},
	)

	return mux
}