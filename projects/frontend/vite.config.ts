import { fileURLToPath, URL } from "node:url";

import { defineConfig } from "vite";
import vue from "@vitejs/plugin-vue";
import wasm from "vite-plugin-wasm";
import topLevelAwait from "vite-plugin-top-level-await";
import { NodeGlobalsPolyfillPlugin } from "@esbuild-plugins/node-globals-polyfill";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [topLevelAwait(), wasm(), vue()],
  worker: {
    format: "es",
    plugins: [topLevelAwait(), wasm()],
  },
  optimizeDeps: {
    // This is necessary because otherwise `vite dev` includes two separate
    // versions of the JS wrapper. This causes problems because the JS
    // wrapper has a module level variable to track JS side heap
    // allocations, initializing this twice causes horrible breakage
    exclude: ["@automerge/automerge-wasm"],
    esbuildOptions: {
      // Node.js global to browser globalThis
      define: {
        global: "globalThis",
      },
      // Enable esbuild polyfill plugins
      plugins: [
        NodeGlobalsPolyfillPlugin({
          buffer: true,
        }),
      ],
    },
  },
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
  // 本地运行配置，及反向代理配置
  server: {
    port: 4000,
    hmr: {
      overlay: false,
    },
    proxy: {
      "/api": {
        target: "http://localhost:3000",
        changeOrigin: true, // 开启跨域
        rewrite: (path) => path.replace(/\/api/, ""), // 路径重写
      },
    },
  },
});
