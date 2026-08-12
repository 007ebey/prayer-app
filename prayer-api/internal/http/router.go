package httpapi

import (
	"net/http"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

func NewRouter(auth *AuthHandler, users *UserHandler, userRoles *UserRoleHandler,
	prayerGroupHandler *PrayerGroupHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.Handle(
		"POST /api/auth/login",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(auth.Login),
		),
	)

	mux.Handle(
		"GET /api/users/{id}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(users.Get),
		),
	)

	mux.Handle(
		"POST /api/users/{userID}/roles/{roleID}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(userRoles.Assign),
		),
	)

	mux.Handle(
		"DELETE /api/users/{userID}/roles/{roleID}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(userRoles.Remove),
		),
	)

	mux.HandleFunc("GET /api/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]any{
			"status": "ok",
		})
	})

	mux.HandleFunc("POST /api/prayer-groups", prayerGroupHandler.Create)

	mux.HandleFunc("GET /api/prayer-groups", prayerGroupHandler.Get)

	mux.Handle(
		"GET /api/prayer-groups/{groupID}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(prayerGroupHandler.GetByID),
		),
	)
	
	return mux
}
