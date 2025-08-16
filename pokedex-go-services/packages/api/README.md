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
   cd pokedex-go-services/packages/api
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

## API Endpoints

### Base URL

```
http://localhost:8080/api/v1
```

### Health Check

- **GET** `/api/v1/health` - Check if the API is working

### Pokémon Endpoints

#### Get Pokémon by ID

- **GET** `/api/v1/pokemon/id/{id}`
- **Description**: Retrieve a specific Pokémon by its unique ID number
- **Parameters**:
  - `id` (path): Pokémon ID (1-1025)
- **Example**: `GET /api/v1/pokemon/id/25`
- **Response**: Single Pokémon data

#### Get Pokémon by Name

- **GET** `/api/v1/pokemon/name/{name}`
- **Description**: Retrieve a specific Pokémon by its name
- **Parameters**:
  - `name` (path): Pokémon name (case-insensitive)
- **Example**: `GET /api/v1/pokemon/name/pikachu`
- **Response**: Single Pokémon data

#### Search Pokémon

- **GET** `/api/v1/pokemon/search`
- **Description**: Intelligent search that automatically detects if the query is an ID or name
- **Parameters**:
  - `q` (query): Search query (ID number or Pokémon name)
- **Examples**:
  - `GET /api/v1/pokemon/search?q=25` (search by ID)
  - `GET /api/v1/pokemon/search?q=pikachu` (search by name)
- **Response**: Single Pokémon data

### Root Endpoint

- **GET** `/` - API information and available endpoints

## API Response Format

### Success Response

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

### Error Response

```json
{
  "error": "Error message description"
}
```

## Usage Examples

### cURL Examples

```bash
# Health check
curl http://localhost:8080/api/v1/health

# Search Pokémon by ID
curl http://localhost:8080/api/v1/pokemon/id/25

# Search Pokémon by name
curl http://localhost:8080/api/v1/pokemon/name/pikachu

# Intelligent search (ID or name)
curl "http://localhost:8080/api/v1/pokemon/search?q=25"
curl "http://localhost:8080/api/v1/pokemon/search?q=pikachu"

# Get API information
curl http://localhost:8080/
```

### JavaScript/TypeScript Examples

```typescript
// Search by ID
const response = await fetch('http://localhost:8080/api/v1/pokemon/id/25');
const pokemon = await response.json();

// Search by name
const response = await fetch(
  'http://localhost:8080/api/v1/pokemon/name/pikachu'
);
const pokemon = await response.json();

// Intelligent search
const response = await fetch(
  'http://localhost:8080/api/v1/pokemon/search?q=pikachu'
);
const pokemon = await response.json();
```

## Data Types

### Pokémon Response Structure

```typescript
interface PokemonResponse {
  id: number; // Unique Pokémon ID (1-1025)
  name: string; // Pokémon name
  types: string[]; // Array of type names
  stats: {
    hp: number; // Hit Points
    attack: number; // Attack stat
    defense: number; // Defense stat
    special_attack: number; // Special Attack stat
    special_defense: number; // Special Defense stat
    speed: number; // Speed stat
  };
  image_url: string; // Official Pokémon sprite URL
  height: number; // Height in decimeters
  weight: number; // Weight in hectograms
  abilities: string[]; // Array of ability names
}
```

## Environment Variables

| Variable           | Description           | Default Value               | Required                 |
| ------------------ | --------------------- | --------------------------- | ------------------------ |
| `PORT`             | Server port           | `8080`                      | No                       |
| `ENVIRONMENT`      | Execution environment | `development`               | No                       |
| `GIN_MODE`         | Gin framework mode    | `debug`                     | No                       |
| `POKEAPI_BASE_URL` | PokeAPI base URL      | `https://pokeapi.co/api/v2` | No                       |
| `OPENAI_API_KEY`   | OpenAI API key        | ``                          | No (for future features) |

## Technologies Used

- **Go 1.21+** - Main programming language
- **Gin** - High-performance HTTP web framework
- **PokeAPI** - Public Pokémon database API
- **godotenv** - Environment variables management
- **net/http** - Standard HTTP client for API calls

## Error Handling

The API returns appropriate HTTP status codes:

- **200 OK** - Successful request
- **400 Bad Request** - Invalid parameters (ID out of range, empty name)
- **500 Internal Server Error** - Server-side errors (PokeAPI unavailable, parsing errors)

## Rate Limiting & Performance

- **PokeAPI Integration**: Direct integration with official Pokémon database
- **Timeout**: 10 seconds for external API calls
- **CORS**: Enabled for cross-origin requests
- **Response Time**: Typically under 500ms for successful requests

## Running Tests

### Run All Tests

```bash
cd pokedex-go-services/packages/api
go test -v ./...
```

### Run Tests with Coverage

```bash
cd pokedex-go-services/packages/api
./scripts/test.sh
```

The HTML coverage report will be generated as `coverage.html` in the backend folder.

## Future Enhancements

- [ ] **OpenAI Integration**: AI-powered Pokémon explanations
- [ ] **Caching Layer**: Redis cache for improved performance
- [ ] **Rate Limiting**: API usage limits and quotas
- [ ] **Advanced Search**: Filter by type, stats, abilities
- [ ] **Bulk Operations**: Multiple Pokémon retrieval
- [ ] **User Authentication**: Personalized experiences
- [ ] **Analytics**: Usage statistics and monitoring

## Contributing

Please read the main [CONTRIBUTING.md](../../../CONTRIBUTING.md) file for contribution guidelines.

## License

This project is under MIT license.
