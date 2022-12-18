const path = require("path");

module.exports = {
  target: "web",
  entry: "./src/index.js",
  output: {
    filename: "main.js",
    path: path.resolve(__dirname, "public"),
  },
  mode: "development", // or production
  module: {
    rules: [
      {
        test: /\.css$/i,
        loader: "css-loader",
        options: {
          url: true,
        },
      },
    ],
  },
  performance: {
    // we dont want the wasm blob to generate warnings
    hints: false,
    maxEntrypointSize: 512000,
    maxAssetSize: 512000,
  },
};
