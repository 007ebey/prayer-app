package main

import (
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"
    "time"
	appauth "prayer-api/internal/application/auth"
	appprayergroup "prayer-api/internal/application/prayergroup"
	appprayersession "prayer-api/internal/application/prayersession"
	presence "prayer-api/internal/application/presence"
	"prayer-api/internal/application/userprofile"
	"prayer-api/internal/application/userrole"
	identityauth "prayer-api/internal/auth"
	"prayer-api/internal/config"
	httpapi "prayer-api/internal/http"
	"prayer-api/internal/repository/memory"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	clerk.SetKey(cfg.ClerkSecretKey)

	users := memory.NewUserRepository(cfg)
	roles := memory.NewRoleRepository()
	groups := memory.NewPrayerGroupRepository()
	ids := memory.NewIDGenerator()
	prayerSessions := memory.NewPrayerSessionRepository()
	prayerPoints := memory.NewPrayerPointRepository()
	heartBeats := memory.NewPresenceRepository()

	// claims, ok := clerk.SessionClaimsFromContext(r.Context())

	loginService := appauth.NewService(
		users,
		roles,
		groups,
		ids,
	)

	profileService := userprofile.NewService(
		users,
		roles,
		groups,
	)

	userRoleService := userrole.NewService(
		users,
		roles,
	)

	sessionValidator := presence.NewSessionValidator(
		prayerSessions,
		users,
	)

	identityProvider := identityauth.NewClerkProvider()

	authHandler := httpapi.NewAuthHandler(
		loginService,
		identityProvider,
	)

	prayerGroupCreateService := appprayergroup.NewCreateService(
		users,
		roles,
		groups,
		ids,
	)

	prayerGroupListService := appprayergroup.NewListService(
	    users,
	    groups,
    )

    prayerGroupGetService := appprayergroup.NewGetService(
	    users,
	    groups,
    )

    prayerGroupUpdateService := appprayergroup.NewUpdateService(
	    groups,
    )

	prayerGroupDeleteService := appprayergroup.NewDeleteService(
		groups,
		users,
		roles,
	)

	prayerGroupAssignService := appprayergroup.NewAssignPrayerGroupService(
       users,
	   groups,
	   roles,
	)

	prayerGroupRemoveService := appprayergroup.NewRemovePrayerGroupService(
		users,
		groups,
	)

	prayerGroupBlockService := appprayergroup.NewBlockPrayerGroupService(
		users,
	    groups,
	    roles,
	)

	heartbeatService	 := presence.NewHeartbeatService(
		heartBeats,
		sessionValidator,
	)

	listParticipantsService := presence.NewListParticipantsService(
		heartBeats,
		users,
		5*time.Minute,
	)

    prayerGroupHandler := httpapi.NewPrayerGroupHandler(
	    prayerGroupCreateService,
	    prayerGroupListService,
	    prayerGroupGetService,
	    prayerGroupUpdateService,
		prayerGroupDeleteService,
		prayerGroupAssignService,
		prayerGroupRemoveService,
		prayerGroupBlockService,
    )

    // ─────────────────────────────────────────────
	// Prayer Sessions
	// ─────────────────────────────────────────────

	prayerSessionService := appprayersession.NewService(
		prayerSessions,
		users,
		prayerPoints,
	)

	prayerSessionHandler := httpapi.NewPrayerSessionHandler(
		prayerSessionService,
		prayerSessionService,
		prayerSessionService,
		prayerSessionService,
	)

	userHandler := httpapi.NewUserHandler(
		profileService,
	)

	userRoleHandler := httpapi.NewUserRoleHandler(
		userRoleService,
	)

	presenceHandler := httpapi.NewPresenceHandler(
	   heartbeatService,
	   listParticipantsService,
	)

	router := httpapi.NewRouter(
		authHandler,
		userHandler,
		userRoleHandler,
		prayerGroupHandler,
		prayerSessionHandler,
		presenceHandler,
	)

	address := ":" + cfg.Port
	log.Printf("Prayer API listening on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
