package config

import (
	"os"
)

// Config represents the application configuration
type Config struct {
	PokeAPIBaseURL string
	OpenAIAPIKey   string
	ServerPort     string
	Environment    string
}

// New creates a new instance of Config
func New() *Config {
	return &Config{
		PokeAPIBaseURL: getEnv("POKEAPI_BASE_URL", "https://pokeapi.co/api/v2"),
		OpenAIAPIKey:   getEnv("OPENAI_API_KEY", ""),
		ServerPort:     getEnv("PORT", "8080"),
		Environment:    getEnv("ENVIRONMENT", "development"),
	}
}

// getEnv returns the value of the environment variable or the default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
} 