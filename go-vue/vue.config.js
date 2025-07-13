module.exports = {
    devServer: {
        proxy: {
            '/v1/medical': {
                target: 'http://127.0.0.1:8885',
                changeOrigin: true
            },
            '/v1/data': {
                target: 'http://127.0.0.1:8888',
                changeOrigin: true
            },
            '/v1/patient': {
                target: 'http://127.0.0.1:8881',
                changeOrigin: true
            },
            '/v1/doctor': {
                target: 'http://127.0.0.1:8883',
                changeOrigin: true
            },
            '/v1/user': {
                target: 'http://127.0.0.1:8884',
                changeOrigin: true
            }
        }
    }
}