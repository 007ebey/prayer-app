package main

import (
	"log"
	"net/http"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
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

	identityProvider := identityauth.NewClerkProvider()

	authHandler := httpapi.NewAuthHandler(
		loginService,
		identityProvider,
	)

	router := httpapi.NewRouter(authHandler)

	address := ":" + cfg.Port
	log.Printf("Prayer API listening on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
