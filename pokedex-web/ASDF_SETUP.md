# 🚀 Configuração do asdf para PokedexIA

Este projeto usa o **asdf** para gerenciar as versões do Node.js e pnpm.

## 📋 Pré-requisitos

### 1. Instalar o asdf

**macOS (usando Homebrew):**

```bash
brew install asdf
```

**Linux:**

```bash
git clone https://github.com/asdf-vm/asdf.git ~/.asdf --branch v0.13.1
echo '. "$HOME/.asdf/asdf.sh"' >> ~/.bashrc
echo '. "$HOME/.asdf/completions/asdf.bash"' >> ~/.bashrc
source ~/.bashrc
```

**Windows (usando Chocolatey):**

```bash
choco install asdf
```

### 2. Adicionar plugins necessários

```bash
# Plugin para Node.js
asdf plugin add nodejs

# Plugin para pnpm
asdf plugin add pnpm
```

## 🎯 Configuração do Projeto

### 1. Instalar versões específicas

```bash
cd pokedex-web

# Instalar Node.js 20.9.0
asdf install nodejs 20.9.0

# Instalar pnpm 8.10.2
asdf install pnpm 8.10.2

# Definir versões locais do projeto
asdf local nodejs 20.9.0
asdf local pnpm 8.10.2
```

### 2. Verificar instalação

```bash
# Verificar versão do Node.js
node --version  # Deve mostrar v20.9.0

# Verificar versão do pnpm
pnpm --version  # Deve mostrar 8.10.2
```

### 3. Instalar dependências

```bash
pnpm install
```

## 🔧 Scripts Úteis

### Setup automático

```bash
pnpm setup:asdf
```

### Verificar versões

```bash
asdf current
```

### Listar versões disponíveis

```bash
asdf list all nodejs
```

## 📚 Recursos Adicionais

- [Documentação oficial do asdf](https://asdf-vm.com/)
- [Plugin Node.js](https://github.com/asdf-vm/asdf-nodejs)
- [Plugin pnpm](https://github.com/asdf-vm/asdf-pnpm)

## 🚨 Solução de Problemas

### Erro: "command not found: asdf"

- Verifique se o asdf está no PATH
- Execute `source ~/.bashrc` ou reinicie o terminal

### Erro: "Plugin not found"

- Execute `asdf plugin add nodejs` primeiro

### Erro: "Version not found"

- Execute `asdf list all nodejs` para ver versões disponíveis
- Use uma versão mais recente se necessário
