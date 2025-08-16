# PokedexIA - Pokémon with AI

A modern Pokédex application that combines official Pokémon data with artificial intelligence to generate personalized explanations about each Pokémon.

## 🏗️ Project Structure

```
pokedexIA/
├── pokedex-web/              # Frontend (Next.js + TypeScript)
├── pokedex-go-services/      # Backend Services (Go)
│   └── packages/
│       ├── api/              # Pokémon API
│       └── api-gpt/          # GPT API (future)
└── README.md
```

## 🚀 Technologies

### Frontend

- **Next.js 14+** - React framework
- **TypeScript** - Static typing
- **SCSS** - CSS preprocessor with variables and mixins

### Backend

- **Go 1.21+** - Main language
- **Gin** - Web framework
- **PokeAPI** - Pokémon data
- **OpenAI API** - AI explanations (future)

## 📋 Features

### ✅ Implemented

- [x] REST API for Pokémon search
- [x] PokeAPI integration
- [x] Search by ID and name
- [x] Unit tests
- [x] CI/CD with GitHub Actions
- [x] Frontend project setup (Next.js + TypeScript + SCSS)

### 🚧 In Development

- [ ] Frontend API integration
- [ ] Responsive web interface (Portuguese UI)
- [ ] OpenAI GPT integration
- [ ] Personalized explanation generation
- [ ] Cache and optimizations

## 🛠️ How to Run

### Backend (Pokémon API)

```bash
cd pokedex-go-services/packages/api
go mod tidy
go run main.go
```

### Frontend (Web)

```bash
cd pokedex-web

# Install Node.js and pnpm using asdf
asdf plugin add nodejs
asdf plugin add pnpm
asdf install nodejs 20.9.0
asdf install pnpm 8.10.2
asdf local nodejs 20.9.0
asdf local pnpm 8.10.2

# Install dependencies and run
pnpm install
pnpm dev
```

The frontend will be available at `http://localhost:3000`

## 🧪 Tests

### Backend

```bash
cd pokedex-go-services/packages/api
go test -v ./...
```

### Frontend

```bash
cd pokedex-web
pnpm lint
pnpm type-check
```

## 📚 Documentation

- [Backend API](pokedex-go-services/packages/api/README.md)
- [Frontend](pokedex-web/README.md)

## 🤝 Contributing

This project follows a Pull Request workflow. To contribute:

1. **Fork** the project
2. **Clone** your fork locally
3. **Create a branch** for your feature: `git checkout -b feature/new-feature`
4. **Develop** following project standards
5. **Commit** using [Conventional Commits](https://www.conventionalcommits.org/)
6. **Push** to your branch
7. **Open a Pull Request** to `develop`

📖 **Read the [Contributing Guide](CONTRIBUTING.md)** for complete details about:

- Commit conventions
- Code standards
- Development workflow
- Issue and PR templates

## 📄 License

This project is under MIT license.
