package main

import (
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
	appprayergroup "prayer-api/internal/application/prayergroup"
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

	users := memory.NewUserRepository()
	roles := memory.NewRoleRepository()
	groups := memory.NewPrayerGroupRepository()
	ids := memory.NewIDGenerator()

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

    prayerGroupHandler := httpapi.NewPrayerGroupHandler(
	    prayerGroupCreateService,
	    prayerGroupListService,
	    prayerGroupGetService,
	    prayerGroupUpdateService,
		prayerGroupDeleteService,
    )

	userHandler := httpapi.NewUserHandler(
		profileService,
	)

	userRoleHandler := httpapi.NewUserRoleHandler(
		userRoleService,
	)

	router := httpapi.NewRouter(authHandler,
		userHandler,
		userRoleHandler,
		prayerGroupHandler,
	)

	address := ":" + cfg.Port
	log.Printf("Prayer API listening on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
