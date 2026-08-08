package main

import (
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
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

	userHandler := httpapi.NewUserHandler(
		profileService,
	)

	userRoleHandler := httpapi.NewUserRoleHandler(
		userRoleService,
	)

	router := httpapi.NewRouter(
		authHandler,
		userHandler,
		userRoleHandler,
	)

	address := ":" + cfg.Port
	log.Printf("Prayer API listening on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
