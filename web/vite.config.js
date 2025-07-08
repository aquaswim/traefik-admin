import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import devManifest from 'vite-plugin-dev-manifest'

// https://vite.dev/config/
export default defineConfig({
  server: {
    port: 5173,
  },
  plugins: [
      react(),
      devManifest(),
  ],
  build: {
    manifest: true,
    rollupOptions: {
      input: "./src/main.jsx"
    },
  }
})
