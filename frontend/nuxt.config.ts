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
    postcss: {
        plugins: {
            'postcss-rtlcss': { },
        }
    },
    app: {
        head: {
            link: [
                {
                    rel: 'icon',
                    type: 'image/x-icon',
                    href: '/favicon.ico',
                }
            ],
        },
    },
    quasar: {
        lang: 'en-US',
        iconSet: 'svg-mdi-v7',
        plugins: [
            'Dialog',
            'Dark',
            'Screen',
            'Cookies',
        ],
        config: {
            // @ts-ignore
            lang: {
                noHtmlAttrs: true,
            },
        },
        sassVariables: '@/assets/css/quasar.variables.scss',
    },
    css: [
        '@/assets/css/app.scss',
    ],
    runtimeConfig: {
        public: {
            app: {
                rootUrl: 'http://localhost:3001',
            },
            api: {
                baseUrl: 'http://localhost:8080',
            },
        }
    },
});
