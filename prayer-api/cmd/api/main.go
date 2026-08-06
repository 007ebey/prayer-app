package main

import (
	"log"
	"net/http"
	"os"

	"github.com/clerk/clerk-sdk-go/v2"

	appauth "prayer-api/internal/application/auth"
	httpapi "prayer-api/internal/http"
	"prayer-api/internal/repository/memory"
)

func main() {
	clerkSecret := os.Getenv("CLERK_SECRET_KEY")

	if clerkSecret == "" {
		log.Fatal("CLERK_SECRET_KEY is required")
	}

	clerk.SetKey(clerkSecret)

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

	address := ":8080"
	log.Printf("Prayer API listening on %s", address)

	if err := http.ListenAndServe(address, router); err != nil {
		log.Fatal(err)
	}
}
