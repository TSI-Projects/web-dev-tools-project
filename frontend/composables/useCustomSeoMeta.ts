export type OpenGraphCommon = {
    ogTitle?: string;
    ogDescription?: string;
    ogImage?: string;
};

export type OpenGraphWebsite = OpenGraphCommon & {
    ogType: 'website';
};

export type TwitterCardCommon = {
    twitterSite?: string;
    twitterTitle?: string;
    twitterDescription?: string;
    twitterImage?: string;
};

export type TwitterCardSummary = TwitterCardCommon & {
    twitterCard: 'summary' | 'summary_large_image';
};

export type SeoMeta = {
    title: Ref<string | undefined> | string | undefined;
    description?: string;
    openGraph?: OpenGraphWebsite;
    twitterCard?: TwitterCardSummary;
};

const OPENGRAPH_NAMESPACES = {
    'website': 'og: https://ogp.me/ns#',
};

export default function (options?: SeoMeta) {
    const quasar = useQuasar();
    const appConfig = useAppConfig();
    const runtimeConfig = useRuntimeConfig();
    const route = useRoute();

    const openGraph = options?.openGraph;
    const twitterCard = options?.twitterCard;

    const rootUrl = runtimeConfig.public.app.rootUrl;
    const canonicalUrl = new URL(route.path, rootUrl).toString();

    const title = computed<string>(() => {
        const newTitle = unref(options?.title);

        return `${newTitle} - ${appConfig.appName}`;
    });

    const description: string = options?.description
        ? options.description
        : appConfig.footerAboutUs.replace('{{ siteName }}', appConfig.appName);

    // OpenGraph
    if (openGraph) {
        const ogHtmlPrefix = OPENGRAPH_NAMESPACES[openGraph.ogType];

        if (!ogHtmlPrefix) {
            console.error('Unsupported opengraph html namespace prefix.');
        }

        useHead({
            htmlAttrs: {
                prefix: ogHtmlPrefix,
            },
        });

        useServerSeoMeta({
            ...{
                ogType: openGraph.ogType,
                ogTitle: openGraph?.ogTitle || title,
                ogDescription: openGraph?.ogDescription || description,
            }, ...openGraph, ... {
                ogUrl: canonicalUrl,
                ogLocale: quasar.lang.isoName,
                ogSiteName: appConfig.appName,
                ogImage: new URL(openGraph?.ogImage || 'favicon.ico', rootUrl).toString(),
            }
        });
    }
    
    // TwitterCard
    if (twitterCard) {
        useServerSeoMeta({
            ...{
                twitterCard: twitterCard.twitterCard,
                twitterTitle: twitterCard?.twitterTitle || title,
                twitterDescription: twitterCard?.twitterDescription || description,
            }, ...twitterCard, ...{
                twitterImage: new URL(twitterCard?.twitterImage || 'favicon.ico', rootUrl).toString()
            }
        });
    }


    useHead({
        title,
        htmlAttrs: {
            lang: quasar.lang.isoName,
            dir: quasar.lang.rtl ? 'rtl' : 'ltr',
        },
        link: [
            { rel: 'canonical', href: canonicalUrl },
        ],
    });

    useServerSeoMeta({
        description,
    });
}