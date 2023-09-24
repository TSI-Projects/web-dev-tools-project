declare module 'nuxt/schema' {
    interface AppConfigInput {
        appName: string;
        sourceCodeLink: string;
    }

    interface AppConfig {
        appName: string;
        sourceCodeLink: string;
    }
}

export {};