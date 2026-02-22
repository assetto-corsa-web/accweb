import { fileURLToPath, URL } from 'node:url'

import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import vueJsx from '@vitejs/plugin-vue-jsx'
import vueDevTools from 'vite-plugin-vue-devtools'
import ui from '@nuxt/ui/vite'

// https://vite.dev/config/
export default defineConfig({
  plugins: [
    vue(),
    vueJsx(),
    vueDevTools(),
    ui({
      ui: {
        colors: {
          neutral: 'zinc'
        },
        select: {
          slots: {
            content: 'min-w-fit'
          }
        }
      }
    })
  ],
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url))
    },
  },
  server: {
    host: '0.0.0.0',
    hmr: {
      overlay: true, // Mostra erros na tela em vez de apenas quebrar
    },
    watch: {
      // Se estiver usando Windows/WSL2, o 'usePolling' pode resolver atrasos
      usePolling: true,
      interval: 100,
    }

  }
})
