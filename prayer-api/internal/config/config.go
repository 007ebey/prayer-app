package config

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	ClerkSecretKey string
	DefaultAdminEmail string
}

func Load() (Config, error) {
	// Load .env into process environment.
	// If .env doesn't exist, continue because production
	// may provide real environment variables.
	err := godotenv.Load()

	if err != nil {
		return Config{}, fmt.Errorf("failed to load .env: %w", err)
	}

	fmt.Println("CLERK_SECRET_KEY exists:", os.Getenv("CLERK_SECRET_KEY") != "")

	cfg := Config{
		Port:           os.Getenv("PORT"),
		ClerkSecretKey: os.Getenv("CLERK_SECRET_KEY"),
		DefaultAdminEmail: os.Getenv("DEFAULT_ADMIN_EMAIL"),
	}

	if cfg.ClerkSecretKey == "" {
		return Config{}, fmt.Errorf("CLERK_SECRET_KEY is required!")
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
