module.exports = {
  configureWebpack: {
    // devServer: {
    // headers: { "Access-Control-Allow-Origin": "*" },
    devServer: {
      proxy: {
        "/api": {
          target: process.env.VUE_APP_BACKEND_URL,
          pathRewrite: { "^/api": "" },
          changeOrigin: true,
          secure: false,
        },
      },
    },
  },
  // },
};
