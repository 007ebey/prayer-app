package httpapi

import (
	"net/http"

	clerkhttp "github.com/clerk/clerk-sdk-go/v2/http"
)

func NewRouter(
	auth *AuthHandler, 
	users *UserHandler,
	userRoles *UserRoleHandler,
	prayerGroupHandler *PrayerGroupHandler,
	prayerSessionHandler *PrayerSessionHandler,
	presenceHandler *PresenceHandler,
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

	mux.Handle(
		"PATCH /api/prayer-groups/{groupID}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(prayerGroupHandler.Update),
		),
	)

	mux.Handle(
		"DELETE /api/prayer-groups/{groupID}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(prayerGroupHandler.Delete),
		),
	)

	mux.Handle(
        "PUT /api/prayer-groups/{groupID}/users/{userID}",
        clerkhttp.RequireHeaderAuthorization()(
            http.HandlerFunc(prayerGroupHandler.AssignPrayerGroup),
        ),
    )

	mux.Handle(
        "DELETE /api/users/{userID}/prayer-groups/{groupID}",
        clerkhttp.RequireHeaderAuthorization()(
            http.HandlerFunc(prayerGroupHandler.RemovePrayerGroup),
        ),
    )

	mux.Handle(
        "POST /api/users/{userID}/prayer-groups/{groupID}/block",
        clerkhttp.RequireHeaderAuthorization()(
            http.HandlerFunc(prayerGroupHandler.BlockPrayerGroup),
        ),
    )

	mux.Handle(
		"GET /api/prayer-sessions",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(prayerSessionHandler.List),
		),
	)

	mux.Handle(
		"POST /api/prayer-sessions",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(prayerSessionHandler.Create),
		),
	)

	mux.Handle(
		"PATCH /api/prayer-sessions/{sessionID}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(prayerSessionHandler.Update),
		),
	)

	mux.Handle(
		"DELETE /api/prayer-sessions/{sessionID}",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(prayerSessionHandler.Delete),
		),
	)

	mux.Handle(
		"POST /api/prayer-sessions/{sessionID}/heartbeat",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(presenceHandler.Heartbeat),
		),
	)

	mux.Handle(
		"GET /api/prayer-sessions/{sessionID}/participants",
		clerkhttp.RequireHeaderAuthorization()(
			http.HandlerFunc(presenceHandler.ListParticipants),
		),
	)

	return corsMiddleware(mux)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		w.Header().Set(
			"Access-Control-Allow-Methods",
			"GET, POST, PUT, PATCH, DELETE, OPTIONS",
		)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization",
		)

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}