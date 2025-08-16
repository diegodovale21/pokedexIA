import type { Metadata } from 'next';
import '@/styles/globals.scss';

export const metadata: Metadata = {
  title: 'PokedexIA - Pokémon com IA',
  description:
    'Uma aplicação moderna de Pokédex que combina dados oficiais de Pokémon com inteligência artificial',
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang='pt-BR'>
      <body>{children}</body>
    </html>
  );
}
