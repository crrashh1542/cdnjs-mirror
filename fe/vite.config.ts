import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
   plugins: [
      react()
   ],
   resolve: {},
   server: {
      port: 23646,
      host: true
   },
   build: {
      rollupOptions: {
         output: {
            hashCharacters: 'hex',
            assetFileNames: '_assets/[name].[hash].[ext]',
            chunkFileNames: '_assets/[name].[hash].js',
            entryFileNames: '_assets/[name].[hash].js',
            manualChunks: {
               vendor_highlight: ['highlight.js', 'highlight.js/lib/languages/javascript.js'],
               vendor_react: ['react-dom']
            }
         }
      }
   },
   css: {
      preprocessorOptions: {
         less: {
            javascriptEnabled: true
         }
      }
   }
})
