declare module 'nuxt/schema' {
    // App config
    interface AppConfigInput {
        appName: string;
        sourceCodeLink: string;
        footerAboutUs: string;
    }

    interface AppConfig {
        appName: string;
        sourceCodeLink: string;
        footerAboutUs: string;
    }

    // Runtime config
    interface PublicRuntimeConfig {
        app: {
            rootUrl: string;
        }
        api: {
            baseUrl: string;
        }
    }
}

export { };