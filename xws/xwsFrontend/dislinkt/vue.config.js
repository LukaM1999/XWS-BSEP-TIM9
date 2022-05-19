const fs = require('fs');

module.exports = {
  configureWebpack: {
    devtool: 'source-map',
  },
  devServer: {
    host: 'localhost',
    https: {
      key: fs.readFileSync('../../../cert/server-key.pem'),
      cert: fs.readFileSync('../../../cert/server-cert.pem'),
    },
    //public: process.env.VUE_APP_BACKEND,
  },
  //publicPath: "/"
}
