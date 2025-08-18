/** @type {import('next').NextConfig} */
const nextConfig = {
  swcMinify: true,
  experimental: {
    forceSwcTransforms: true,
  },
  sassOptions: {
    includePaths: ['./src/styles'],
  },
};

module.exports = nextConfig;
