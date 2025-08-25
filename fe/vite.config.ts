import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

// https://vitejs.dev/config/
export default defineConfig({
   plugins: [
      react()
   ],
   resolve: {},
   server: {
      port: 23656,
      host: true,
      proxy: {
         "/getSite": {
            target: "http://localhost:23657",
            changeOrigin: true
         }
      }
   },
   build: {
      rollupOptions: {
         output: {
            hashCharacters: 'hex',
            assetFileNames: '_assets/[name]-[hash].[ext]',
            chunkFileNames: '_assets/[name]-[hash].js',
            entryFileNames: '_assets/[name]-[hash].js',
            manualChunks(i: String) {
               if (i.includes('react-dom')) {
                  return 'vendor_react-dom'
               } else if(i.includes('react-syntax-highlighter')) {
                  return 'vendor_highlight'
               } else if(i.includes('react')) {
                  return 'vendor_react'
               }
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
