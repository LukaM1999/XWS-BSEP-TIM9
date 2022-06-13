const fs = require('fs');


module.exports = {
    configureWebpack: {
        devtool: 'source-map',
        optimization: { concatenateModules: false, providedExports: false, usedExports: false },
        resolve: {
            alias: {
                'vue$': 'vue/dist/vue.esm.js'
            }
        }

    },
    devServer: {
        host: 'localhost',
        https: {
            key: fs.readFileSync('../../cert/server-key.pem'),
            cert: fs.readFileSync('../../cert/server-cert.pem'),
        },
        //public: process.env.VUE_APP_BACKEND,
    },
}