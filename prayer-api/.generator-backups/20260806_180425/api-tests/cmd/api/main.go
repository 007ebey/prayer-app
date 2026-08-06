package main

import (
	"log"
	"net/http"
	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
	httpapi "prayer-api/internal/http"
	"prayer-api/internal/config"
	"prayer-api/internal/repository/memory"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Configure Clerk.
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

	authHandler := httpapi.NewAuthHandler(loginService)
	router := httpapi.NewRouter(authHandler)

	address := ":8181"
	log.Printf("Prayer API listening on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
