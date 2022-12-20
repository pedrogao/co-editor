import { defineConfig } from "vite";
import { resolve } from "path";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [],
  publicDir: "static",
  build: {
    rollupOptions: {
      input: {
        main: resolve(__dirname, "index.html"),
        essay: resolve(__dirname, "essay/index.html"),
      },
    },
  },
});
