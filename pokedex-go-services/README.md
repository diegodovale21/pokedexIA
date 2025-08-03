# PokedexIA Go Services

Backend services for the PokedexIA application, built with Go.

## Structure

```
pokedex-go-services/
└── packages/
    ├── api/        # Pokémon API service
    └── api-gpt/    # GPT integration service (future)
```

## Services

### API Service (`packages/api/`)
- Pokémon data integration with PokeAPI
- RESTful endpoints for Pokémon search
- Data transformation and caching

### GPT API Service (`packages/api-gpt/`) - Coming Soon
- OpenAI GPT integration
- Pokémon explanation generation
- Intelligent content creation

## Development

Each package can be developed and deployed independently. They follow the same Go project structure and conventions.

## Running Services

### Pokémon API
```bash
cd packages/api
go run main.go
```

### GPT API (when implemented)
```bash
cd packages/api-gpt
go run main.go
```

## Architecture

This monorepo structure allows for:
- Independent service development
- Shared utilities and types
- Easy deployment and scaling
- Clear separation of concerns 