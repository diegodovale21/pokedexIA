package api

import (
	"github.com/gin-gonic/gin"
	"pokedexia-backend/internal/config"
	"pokedexia-backend/internal/handlers"
)

// SetupRoutes configures all the API routes
func SetupRoutes(router *gin.Engine, cfg *config.Config) {
	// Create the handlers
	pokemonHandler := handlers.NewPokemonHandler(cfg)

	// API routes group
	api := router.Group("/api/v1")
	{
		// Health check
		api.GET("/health", pokemonHandler.HealthCheck)

		// Pokemon routes
		pokemon := api.Group("/pokemon")
		{
			pokemon.GET("/id/:id", pokemonHandler.GetPokemonByID)
			pokemon.GET("/name/:name", pokemonHandler.GetPokemonByName)
			pokemon.GET("/search", pokemonHandler.SearchPokemon)
		}
	}

	// Root route
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to PokedexIA API!",
			"version": "1.0.0",
			"endpoints": gin.H{
				"health": "/api/v1/health",
				"pokemon_by_id": "/api/v1/pokemon/id/:id",
				"pokemon_by_name": "/api/v1/pokemon/name/:name",
				"search_pokemon": "/api/v1/pokemon/search?q=:query",
			},
		})
	})
} 