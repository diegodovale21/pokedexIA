package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"pokedexia-backend/internal/config"
	"pokedexia-backend/internal/types"
	"pokedexia-backend/internal/services"
)

// PokemonHandler represents the handler for Pokémon endpoints
type PokemonHandler struct {
	pokeAPIService *services.PokeAPIService
}

// NewPokemonHandler creates a new instance of the handler
func NewPokemonHandler(cfg *config.Config) *PokemonHandler {
	return &PokemonHandler{
		pokeAPIService: services.NewPokeAPIService(cfg),
	}
}

// GetPokemonByID searches for a Pokémon by ID
func (h *PokemonHandler) GetPokemonByID(c *gin.Context) {
	idStr := c.Param("id")
	
	// Validate the ID
	id, err := h.pokeAPIService.ValidatePokemonID(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Search for the Pokémon
	pokemon, err := h.pokeAPIService.GetPokemonByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar Pokémon: " + err.Error(),
		})
		return
	}

	// Transform to the response format
	response := h.pokeAPIService.TransformPokemonToResponse(pokemon)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// GetPokemonByName searches for a Pokémon by name
func (h *PokemonHandler) GetPokemonByName(c *gin.Context) {
	name := c.Param("name")
	
	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Nome do Pokémon é obrigatório",
		})
		return
	}

	// Search for the Pokémon
	pokemon, err := h.pokeAPIService.GetPokemonByName(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar Pokémon: " + err.Error(),
		})
		return
	}

	// Transform to the response format
	response := h.pokeAPIService.TransformPokemonToResponse(pokemon)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// SearchPokemon searches for a Pokémon by ID or name
func (h *PokemonHandler) SearchPokemon(c *gin.Context) {
	query := c.Query("q")
	
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Parâmetro 'q' é obrigatório",
		})
		return
	}

	var pokemon *types.Pokemon
	var err error

		// Try to convert to ID first
	if id, err := strconv.Atoi(query); err == nil {
		// It's a number, search by ID
		if id < 1 || id > 1025 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "ID deve estar entre 1 e 1025",
			})
			return
		}
		pokemon, err = h.pokeAPIService.GetPokemonByID(id)
	} else {
		// It's a name, search by name
		pokemon, err = h.pokeAPIService.GetPokemonByName(query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Erro ao buscar Pokémon: " + err.Error(),
		})
		return
	}

	// Transform to the response format
	response := h.pokeAPIService.TransformPokemonToResponse(pokemon)

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    response,
	})
}

// HealthCheck checks if the API is working
func (h *PokemonHandler) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": "PokedexIA API está funcionando",
	})
} 