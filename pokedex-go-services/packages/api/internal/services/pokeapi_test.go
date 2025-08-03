package services

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"pokedexia-backend/internal/config"
	"pokedexia-backend/internal/types"
)

func TestNewPokeAPIService(t *testing.T) {
	cfg := &config.Config{
		PokeAPIBaseURL: "https://pokeapi.co/api/v2",
	}

	service := NewPokeAPIService(cfg)

	if service == nil {
		t.Fatal("Expected service to be created, got nil")
	}

	if service.baseURL != cfg.PokeAPIBaseURL {
		t.Errorf("Expected baseURL to be %s, got %s", cfg.PokeAPIBaseURL, service.baseURL)
	}

	if service.httpClient == nil {
		t.Fatal("Expected httpClient to be initialized, got nil")
	}
}

func TestValidatePokemonID(t *testing.T) {
	cfg := &config.Config{PokeAPIBaseURL: "https://pokeapi.co/api/v2"}
	service := NewPokeAPIService(cfg)

	tests := []struct {
		name    string
		idStr   string
		wantID  int
		wantErr bool
	}{
		{"valid ID", "25", 25, false},
		{"valid ID min", "1", 1, false},
		{"valid ID max", "1025", 1025, false},
		{"invalid ID - zero", "0", 0, true},
		{"invalid ID - negative", "-1", 0, true},
		{"invalid ID - too high", "1026", 0, true},
		{"invalid ID - not number", "abc", 0, true},
		{"invalid ID - empty", "", 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := service.ValidatePokemonID(tt.idStr)

			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePokemonID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotID != tt.wantID {
				t.Errorf("ValidatePokemonID() = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func TestGetPokemonByID_Success(t *testing.T) {
	// Create a mock server
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/pokemon/25" {
			t.Errorf("Expected to request '/pokemon/25', got: %s", r.URL.Path)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{
			"id": 25,
			"name": "pikachu",
			"height": 4,
			"weight": 60,
			"types": [
				{
					"slot": 1,
					"type": {
						"name": "electric",
						"url": "https://pokeapi.co/api/v2/type/13/"
					}
				}
			],
			"stats": [
				{
					"base_stat": 35,
					"effort": 0,
					"stat": {
						"name": "hp",
						"url": "https://pokeapi.co/api/v2/stat/1/"
					}
				},
				{
					"base_stat": 55,
					"effort": 0,
					"stat": {
						"name": "attack",
						"url": "https://pokeapi.co/api/v2/stat/2/"
					}
				}
			],
			"abilities": [
				{
					"ability": {
						"name": "static",
						"url": "https://pokeapi.co/api/v2/ability/9/"
					},
					"is_hidden": false,
					"slot": 1
				}
			],
			"sprites": {
				"front_default": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png"
			}
		}`))
	}))
	defer server.Close()

	cfg := &config.Config{PokeAPIBaseURL: server.URL}
	service := NewPokeAPIService(cfg)

	pokemon, err := service.GetPokemonByID(25)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if pokemon == nil {
		t.Fatal("Expected pokemon to be returned, got nil")
	}

	if pokemon.ID != 25 {
		t.Errorf("Expected ID 25, got %d", pokemon.ID)
	}

	if pokemon.Name != "pikachu" {
		t.Errorf("Expected name 'pikachu', got %s", pokemon.Name)
	}

	if len(pokemon.Types) != 1 {
		t.Errorf("Expected 1 type, got %d", len(pokemon.Types))
	}

	if pokemon.Types[0].Type.Name != "electric" {
		t.Errorf("Expected type 'electric', got %s", pokemon.Types[0].Type.Name)
	}
}

func TestGetPokemonByID_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"detail": "Not found."}`))
	}))
	defer server.Close()

	cfg := &config.Config{PokeAPIBaseURL: server.URL}
	service := NewPokeAPIService(cfg)

	pokemon, err := service.GetPokemonByID(99999)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if pokemon != nil {
		t.Errorf("Expected nil pokemon, got %v", pokemon)
	}
}

func TestTransformPokemonToResponse(t *testing.T) {
	pokemon := &types.Pokemon{
		ID:     25,
		Name:   "pikachu",
		Height: 4,
		Weight: 60,
		Types: []types.Type{
			{
				Slot: 1,
				Type: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "electric",
					URL:  "https://pokeapi.co/api/v2/type/13/",
				},
			},
		},
		Stats: []types.Stat{
			{
				BaseStat: 35,
				Effort:   0,
				Stat: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "hp",
					URL:  "https://pokeapi.co/api/v2/stat/1/",
				},
			},
			{
				BaseStat: 55,
				Effort:   0,
				Stat: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "attack",
					URL:  "https://pokeapi.co/api/v2/stat/2/",
				},
			},
		},
		Abilities: []types.Ability{
			{
				Ability: struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				}{
					Name: "static",
					URL:  "https://pokeapi.co/api/v2/ability/9/",
				},
				IsHidden: false,
				Slot:     1,
			},
		},
		Sprites: types.Sprites{
			FrontDefault: "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png",
		},
	}

	cfg := &config.Config{PokeAPIBaseURL: "https://pokeapi.co/api/v2"}
	service := NewPokeAPIService(cfg)

	response := service.TransformPokemonToResponse(pokemon)

	if response == nil {
		t.Fatal("Expected response to be returned, got nil")
	}

	if response.ID != 25 {
		t.Errorf("Expected ID 25, got %d", response.ID)
	}

	if response.Name != "pikachu" {
		t.Errorf("Expected name 'pikachu', got %s", response.Name)
	}

	if len(response.Types) != 1 {
		t.Errorf("Expected 1 type, got %d", len(response.Types))
	}

	if response.Types[0] != "electric" {
		t.Errorf("Expected type 'electric', got %s", response.Types[0])
	}

	if response.Stats.HP != 35 {
		t.Errorf("Expected HP 35, got %d", response.Stats.HP)
	}

	if response.Stats.Attack != 55 {
		t.Errorf("Expected Attack 55, got %d", response.Stats.Attack)
	}

	if len(response.Abilities) != 1 {
		t.Errorf("Expected 1 ability, got %d", len(response.Abilities))
	}

	if response.Abilities[0] != "static" {
		t.Errorf("Expected ability 'static', got %s", response.Abilities[0])
	}

	if response.ImageURL != "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png" {
		t.Errorf("Expected correct image URL, got %s", response.ImageURL)
	}
} 