package services

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"pokedexia-backend/internal/config"
	"pokedexia-backend/internal/types"
)

// PokeAPIService represents the service for integrating with the PokeAPI
type PokeAPIService struct {
	baseURL    string
	httpClient *http.Client
}

// NewPokeAPIService creates a new instance of the service
func NewPokeAPIService(cfg *config.Config) *PokeAPIService {
	return &PokeAPIService{
		baseURL: cfg.PokeAPIBaseURL,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetPokemonByID searches for a Pokémon by ID
func (s *PokeAPIService) GetPokemonByID(id int) (*types.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%d", s.baseURL, id)
	
	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var pokemon types.Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return nil, fmt.Errorf("error deserializing JSON: %w", err)
	}

	return &pokemon, nil
}

// GetPokemonByName searches for a Pokémon by name
func (s *PokeAPIService) GetPokemonByName(name string) (*types.Pokemon, error) {
	url := fmt.Sprintf("%s/pokemon/%s", s.baseURL, strings.ToLower(name))
	
	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	var pokemon types.Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return nil, fmt.Errorf("error deserializing JSON: %w", err)
	}

	return &pokemon, nil
}

// TransformPokemonToResponse transforms the Pokémon from the API to the response format
func (s *PokeAPIService) TransformPokemonToResponse(pokemon *types.Pokemon) *types.PokemonResponse {
	// Extract types
	pokemonTypes := make([]string, len(pokemon.Types))
	for i, t := range pokemon.Types {
		pokemonTypes[i] = t.Type.Name
	}

	// Extract abilities
	abilities := make([]string, len(pokemon.Abilities))
	for i, a := range pokemon.Abilities {
		abilities[i] = a.Ability.Name
	}

	// Organize stats
	stats := types.Stats{}
	for _, stat := range pokemon.Stats {
		switch stat.Stat.Name {
		case "hp":
			stats.HP = stat.BaseStat
		case "attack":
			stats.Attack = stat.BaseStat
		case "defense":
			stats.Defense = stat.BaseStat
		case "special-attack":
			stats.SpecialAttack = stat.BaseStat
		case "special-defense":
			stats.SpecialDefense = stat.BaseStat
		case "speed":
			stats.Speed = stat.BaseStat
		}
	}

	return &types.PokemonResponse{
		ID:        pokemon.ID,
		Name:      pokemon.Name,
		Types:     pokemonTypes,
		Stats:     stats,
		ImageURL:  pokemon.Sprites.FrontDefault,
		Height:    pokemon.Height,
		Weight:    pokemon.Weight,
		Abilities: abilities,
	}
}

// ValidatePokemonID validates if the Pokémon ID is valid
func (s *PokeAPIService) ValidatePokemonID(idStr string) (int, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, fmt.Errorf("invalid ID: %s", idStr)
	}

	if id < 1 || id > 1025 { // Current limit of the PokeAPI
		return 0, fmt.Errorf("ID must be between 1 and 1025")
	}

	return id, nil
} 