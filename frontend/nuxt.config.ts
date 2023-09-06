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
        '@nuxt/content',
    ],
    quasar: {
        lang: 'en-US',
        iconSet: 'svg-mdi-v7',
        plugins: [
            'Dialog',
            'Dark',
            'Screen',
        ],
    },
});
