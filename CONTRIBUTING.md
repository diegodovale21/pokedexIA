# 🤝 Contributing Guide

Thank you for contributing to PokedexIA! This document contains guidelines for contributing to the project.

## 📋 Index

- [How to Contribute](#how-to-contribute)
- [Commit Conventions](#commit-conventions)
- [Development Workflow](#development-workflow)
- [Code Standards](#code-standards)

## 🚀 How to Contribute

### 1. Fork and Clone
```bash
# Fork the repository on GitHub
# Clone your fork
git clone https://github.com/YOUR_USERNAME/pokedexIA.git
cd pokedexIA
```

### 2. Set Up Upstream
```bash
git remote add upstream https://github.com/diegodovale21/pokedexIA.git
```

### 3. Create Feature Branch
```bash
git checkout -b feature/feature-name
```

### 4. Develop
- Make your changes
- Follow code standards
- Add tests when necessary

### 5. Commit and Push
```bash
git add .
git commit -m "feat: add new functionality"
git push origin feature/feature-name
```

### 6. Create Pull Request
- Go to GitHub
- Create a Pull Request to `develop`
- Use the PR template
- Wait for review

## 📝 Commit Conventions

We follow [Conventional Commits](https://www.conventionalcommits.org/):

### Commit Types
- `feat:` New feature
- `fix:` Bug fix
- `docs:` Documentation
- `style:` Formatting, semicolons, etc.
- `refactor:` Code refactoring
- `test:` Adding or fixing tests
- `chore:` Build tasks, dependencies, etc.

### Examples
```bash
feat: add Pokémon search by name
fix: fix ID validation error
docs: update README with new instructions
test: add tests for Pokémon service
refactor: reorganize folder structure
```

### Structure
```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

## 🔄 Development Workflow

### Branches
- `main`: Production (always stable)
- `develop`: Development (feature integration)
- `feature/*`: Features under development
- `hotfix/*`: Urgent fixes

### Flow
1. **Feature Branch** → `develop`
2. **Develop** → `main` (after testing)
3. **Hotfix** → `main` + `develop`

## 📏 Code Standards

### Go (Backend)
- Use `gofmt` for formatting
- Follow Go conventions
- Add tests for new features
- Use descriptive names for variables and functions

### TypeScript/JavaScript (Frontend)
- Use ESLint and Prettier
- Follow TypeScript conventions
- Use descriptive names
- Add types when possible

### General
- Keep commits small and focused
- Write clear commit messages
- Document important changes
- Test your changes

## 🧪 Tests

### Backend
```bash
cd pokedex-go-services/packages/api
go test -v ./...
```

### Frontend
```bash
cd pokedex-web
npm test
```

## 📚 Resources

- [Conventional Commits](https://www.conventionalcommits.org/)
- [GitHub Flow](https://guides.github.com/introduction/flow/)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)

## 🤝 Acknowledgments

Thank you for contributing to PokedexIA! 🎉 