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
    // devServer: {
    //     proxy: "http://localhost:8090"
    // },
}