// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    devtools: {
        enabled: true
    },
    devServer: {
        port: 3001,
    },
    modules: [
        'nuxt-quasar-ui',
    ],
    quasar: {
        lang: 'en-US',
        iconSet: 'svg-mdi-v7',
        plugins: [
            'Dialog',
            'Dark',
            'Screen',
            'Cookies',
        ],
        sassVariables: '@/assets/css/quasar.variables.scss',
    },
    css: [
        '@/assets/css/app.scss',
    ],
    runtimeConfig: {
        public: {
            api: {
                baseUrl: 'http://localhost:8080',
            },
        }
    },
});
