package main

import (
	"fmt"
	"log"
	"net/http"

	"prayer-api/internal/auth"
	"prayer-api/internal/config"
	apphttp "prayer-api/internal/http"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	auth.Configure(cfg.ClerkSecretKey)

	router := apphttp.NewRouter()

	addr := ":" + cfg.Port

	fmt.Printf("Prayer API listening on http://localhost%s\n", addr)

	if err := http.ListenAndServe(addr, router); err != nil {
		log.Fatal(err)
	}
}
