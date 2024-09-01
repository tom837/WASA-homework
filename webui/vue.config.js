module.exports = {
    devServer: {
      proxy: {
        '/api': {
          target: 'http://localhost:3000', // Backend server
          changeOrigin: true,
          secure: false,
        },
      },
    },
  };
  