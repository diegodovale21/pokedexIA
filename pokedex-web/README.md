# PokedexIA Frontend

Frontend application for PokedexIA built with Next.js, TypeScript, and SCSS.

## 🚀 Technologies

- **Next.js 14+** - React framework with App Router
- **TypeScript** - Static typing
- **SCSS** - CSS preprocessor with variables and mixins
- **ESLint** - Code quality and consistency

## 📁 Project Structure

```
pokedex-web/
├── src/
│   ├── app/              # Next.js App Router
│   │   ├── layout.tsx    # Root layout
│   │   └── page.tsx      # Home page
│   ├── components/       # Reusable components
│   ├── styles/           # SCSS styles
│   │   └── globals.scss  # Global styles and variables
│   ├── types/            # TypeScript type definitions
│   └── services/         # API services
├── package.json
├── tsconfig.json
├── next.config.js
└── README.md
```

## 🛠️ How to Run

### Prerequisites

- **asdf** - Version manager for Node.js and pnpm
- Node.js 20.9.0 (managed by asdf)
- pnpm 8.10.2 (managed by asdf)

### Installation

1. Install Node.js and pnpm using asdf:

   ```bash
   # Install asdf if you don't have it
   # https://asdf-vm.com/guide/getting-started.html

   # Install Node.js plugin
   asdf plugin add nodejs

   # Install pnpm plugin
   asdf plugin add pnpm

   # Install Node.js version
   asdf install nodejs 20.9.0

   # Install pnpm version
   asdf install pnpm 8.10.2

   # Set local versions
   asdf local nodejs 20.9.0
   asdf local pnpm 8.10.2
   ```

2. Install dependencies:

   ```bash
   pnpm install
   ```

3. Run the development server:

   ```bash
   pnpm dev
   ```

4. Open [http://localhost:3000](http://localhost:3000) in your browser.

## 📝 Available Scripts

- `pnpm dev` - Start development server
- `pnpm build` - Build for production
- `pnpm start` - Start production server
- `pnpm lint` - Run ESLint
- `pnpm type-check` - Run TypeScript type checking
- `pnpm setup:asdf` - Setup Node.js and pnpm versions with asdf

## 🎨 Styling

This project uses **SCSS** instead of Tailwind CSS:

- **Variables**: Colors, breakpoints, spacing
- **Mixins**: Responsive design, flexbox utilities
- **Nested styles**: Better organization and readability
- **Utility classes**: Common spacing and layout classes

### Breakpoints

- Mobile: `max-width: 768px`
- Tablet: `768px - 1024px`
- Desktop: `min-width: 1024px`

## 🔗 Backend Integration

The frontend connects to the PokedexIA backend API running on `http://localhost:8080`.

## 🚧 Development Status

- [x] Project setup
- [x] Basic layout structure
- [x] SCSS configuration
- [ ] API integration
- [ ] Pokémon search functionality
- [ ] Responsive design
- [ ] Component library

## 📚 Next Steps

1. **API Integration**: Connect to backend services
2. **Components**: Build reusable UI components
3. **Search**: Implement Pokémon search functionality
4. **Styling**: Complete responsive design
5. **Testing**: Add unit and integration tests
