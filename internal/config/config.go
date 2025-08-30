package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	KibelaToken string
	KibelaTeam  string
}

func Load() (*Config, error) {
	// Try to load .env file if it exists
	_ = godotenv.Load()

	config := &Config{
		KibelaToken: os.Getenv("KIBELA_TOKEN"),
		KibelaTeam:  os.Getenv("KIBELA_TEAM"),
	}

	if config.KibelaToken == "" {
		return nil, fmt.Errorf("KIBELA_TOKEN environment variable is required")
	}
	if config.KibelaTeam == "" {
		return nil, fmt.Errorf("KIBELA_TEAM environment variable is required")
	}

	return config, nil
}