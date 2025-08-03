package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Clear environment variables for testing
	os.Unsetenv("POKEAPI_BASE_URL")
	os.Unsetenv("OPENAI_API_KEY")
	os.Unsetenv("PORT")
	os.Unsetenv("ENVIRONMENT")

	cfg := New()

	assert.NotNil(t, cfg)
	assert.Equal(t, "https://pokeapi.co/api/v2", cfg.PokeAPIBaseURL)
	assert.Equal(t, "", cfg.OpenAIAPIKey)
	assert.Equal(t, "8080", cfg.ServerPort)
	assert.Equal(t, "development", cfg.Environment)
}

func TestNew_WithEnvironmentVariables(t *testing.T) {
	// Set environment variables
	os.Setenv("POKEAPI_BASE_URL", "https://test-api.com")
	os.Setenv("OPENAI_API_KEY", "test-key")
	os.Setenv("PORT", "3000")
	os.Setenv("ENVIRONMENT", "production")

	cfg := New()

	assert.NotNil(t, cfg)
	assert.Equal(t, "https://test-api.com", cfg.PokeAPIBaseURL)
	assert.Equal(t, "test-key", cfg.OpenAIAPIKey)
	assert.Equal(t, "3000", cfg.ServerPort)
	assert.Equal(t, "production", cfg.Environment)

	// Clean up
	os.Unsetenv("POKEAPI_BASE_URL")
	os.Unsetenv("OPENAI_API_KEY")
	os.Unsetenv("PORT")
	os.Unsetenv("ENVIRONMENT")
}

func TestGetEnv(t *testing.T) {
	// Test with environment variable set
	os.Setenv("TEST_VAR", "test_value")
	value := getEnv("TEST_VAR", "default_value")
	assert.Equal(t, "test_value", value)

	// Test with environment variable not set
	os.Unsetenv("TEST_VAR")
	value = getEnv("TEST_VAR", "default_value")
	assert.Equal(t, "default_value", value)

	// Test with empty environment variable
	os.Setenv("EMPTY_VAR", "")
	value = getEnv("EMPTY_VAR", "default_value")
	assert.Equal(t, "default_value", value)
} 