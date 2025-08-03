# PokedexIA Backend

Backend for the PokedexIA application developed in Go, responsible for integrating with the PokeAPI and providing Pokémon data.

## Project Structure

```
backend/
├── main.go                 # Application entry point
├── go.mod                  # Go dependencies
├── env.example             # Environment variables example
├── internal/
│   ├── api/               # Route configuration
│   ├── config/            # Application configuration
│   ├── handlers/          # HTTP handlers
│   ├── types/             # Data types
│   └── services/          # Business services
└── README.md              # This file
```

## How to Run

### Prerequisites
- Go 1.21 or higher

### Installation

1. Clone the repository
2. Navigate to the backend folder:
   ```bash
   cd backend
   ```

3. Install dependencies:
   ```bash
   go mod tidy
   ```

4. Configure environment variables:
   ```bash
   cp env.example .env
   # Edit the .env file with your settings
   ```

5. Run the application:
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`

## Available Endpoints

### Health Check
- **GET** `/api/v1/health` - Check if the API is working

### Pokémon
- **GET** `/api/v1/pokemon/id/{id}` - Search Pokémon by ID (1-1025)
- **GET** `/api/v1/pokemon/name/{name}` - Search Pokémon by name
- **GET** `/api/v1/pokemon/search?q={query}` - Search Pokémon by ID or name

### Usage Examples

```bash
# Search Pokémon by ID
curl http://localhost:8080/api/v1/pokemon/id/25

# Search Pokémon by name
curl http://localhost:8080/api/v1/pokemon/name/pikachu

# Search Pokémon by query (ID or name)
curl http://localhost:8080/api/v1/pokemon/search?q=25
curl http://localhost:8080/api/v1/pokemon/search?q=pikachu
```

## API Response

The API returns data in the following format:

```json
{
  "success": true,
  "data": {
    "id": 25,
    "name": "pikachu",
    "types": ["electric"],
    "stats": {
      "hp": 35,
      "attack": 55,
      "defense": 40,
      "special_attack": 50,
      "special_defense": 50,
      "speed": 90
    },
    "image_url": "https://raw.githubusercontent.com/PokeAPI/sprites/master/sprites/pokemon/25.png",
    "height": 4,
    "weight": 60,
    "abilities": ["static", "lightning-rod"]
  }
}
```

## Environment Variables

- `PORT` - Server port (default: 8080)
- `ENVIRONMENT` - Execution environment (development/production)
- `GIN_MODE` - Gin mode (debug/release)
- `POKEAPI_BASE_URL` - PokeAPI base URL
- `OPENAI_API_KEY` - OpenAI API key (for future features)

## Technologies Used

- **Go 1.21+** - Main language
- **Gin** - Web framework
- **PokeAPI** - Public Pokémon API
- **godotenv** - Environment variables management 

## Running Tests

To run all unit tests:

```bash
cd backend
go test -v ./...
```

To run tests with coverage and generate a report:

```bash
cd backend
./scripts/test.sh
```

The HTML coverage report will be generated as `coverage.html` in the backend folder. 