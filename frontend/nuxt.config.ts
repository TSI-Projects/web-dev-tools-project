// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
    devtools: {
        enabled: true
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
        ],
    },
});
