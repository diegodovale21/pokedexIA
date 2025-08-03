# PokedexIA - Pokémon com IA

Uma aplicação moderna de Pokédex que combina dados oficiais de Pokémon com inteligência artificial para gerar explicações personalizadas sobre cada Pokémon.

## 🏗️ Estrutura do Projeto

```
pokedexIA/
├── pokedex-web/              # Frontend (Next.js + TypeScript)
├── pokedex-go-services/      # Backend Services (Go)
│   └── packages/
│       ├── api/              # API do Pokémon
│       └── api-gpt/          # API do GPT (futuro)
└── README.md
```

## 🚀 Tecnologias

### Frontend
- **Next.js 14+** - Framework React
- **TypeScript** - Tipagem estática
- **Tailwind CSS** - Estilização

### Backend
- **Go 1.21+** - Linguagem principal
- **Gin** - Framework web
- **PokeAPI** - Dados dos Pokémon
- **OpenAI API** - Explicações com IA (futuro)

## 📋 Funcionalidades

### ✅ Implementadas
- [x] API REST para busca de Pokémon
- [x] Integração com PokeAPI
- [x] Busca por ID e nome
- [x] Testes unitários
- [x] CI/CD com GitHub Actions

### 🚧 Em Desenvolvimento
- [ ] Interface web responsiva
- [ ] Integração com OpenAI GPT
- [ ] Geração de explicações personalizadas
- [ ] Cache e otimizações

## 🛠️ Como Executar

### Backend (API do Pokémon)
```bash
cd pokedex-go-services/packages/api
go mod tidy
go run main.go
```

### Frontend (Web)
```bash
cd pokedex-web
npm install
npm run dev
```

## 🧪 Testes

### Backend
```bash
cd pokedex-go-services/packages/api
go test -v ./...
```

## 📚 Documentação

- [Backend API](pokedex-go-services/packages/api/README.md)
- [Frontend](pokedex-web/README.md)

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

## 📄 Licença

Este projeto está sob a licença MIT.
