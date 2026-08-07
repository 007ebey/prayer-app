package main

import (
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
	"prayer-api/internal/application/userprofile"
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

	identityProvider := identityauth.NewClerkProvider()

	authHandler := httpapi.NewAuthHandler(
		loginService,
		identityProvider,
	)

	userHandler := httpapi.NewUserHandler(
		profileService,
	)

	router := httpapi.NewRouter(
		authHandler,
		userHandler,
	)

	address := ":" + cfg.Port
	log.Printf("Prayer API listening on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
