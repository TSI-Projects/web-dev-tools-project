import qs from 'qs';

export type Post = {
    id: string;
    title: string;
    preview_img: string;
    description: string;
    price: string;
    url: string;
};

export type Pagination = {
    source: string;
    has_next: boolean;
};

export type SseFetchParameters = () => {
    query: {
        query: string | undefined;
        sources: string[] | string | undefined;
        categories: string[] | string | undefined;
        price: {
            min: number | undefined;
            max: number | undefined;
        };
    };
};

export type SseFetchResultExecuteOptions = {
    page: number;
    onFinish?: () => void;
};

export type SseFetchResult = {
    posts: Ref<Post[]>;
    error: Ref<boolean>;
    pending: Ref<boolean>;
    eof: Ref<boolean>;
    execute: (options: SseFetchResultExecuteOptions) => void;
    close: () => void;
    resetEofSources: () => void;
};

export default function () {
    const { public: publicConfig } = useRuntimeConfig();

    const sseFetch = (parametersResolver: SseFetchParameters): SseFetchResult => {
        const pending = ref<boolean>(false);
        const error = ref<boolean>(false);
        const eof = ref<boolean>(false);
        const posts = ref<Post[]>([]);

        // EventSource object does not exists on server,
        // so just return empty state.
        if (process.server) {
            return {
                posts,
                error,
                pending,
                eof,
                execute: () => { },
                close: () => { },
                resetEofSources: () => { },
            };
        }
        
        let close: () => void = () => { };

        const eofSources = new Set<string>();

        const resetEofSources = (): void => {
            eofSources.clear();
        };

        const execute = (options: SseFetchResultExecuteOptions): void => {
            if (pending.value) {
                close();
            }

            const params = parametersResolver();

            // get sources
            const allSources = Array.from(params.query.sources || []);

            const availableSources = new Set(allSources.filter((source: string) => {
                return ! eofSources.has(source);
            }));

            // build pages
            const pages = Object.fromEntries(
                Array.from(availableSources).map((source) => {
                    return [
                        `${source}_page`,
                        options.page,
                    ];
                }),
            );

            // build url
            const fullUrl = new URL('/posts/search', publicConfig.api.baseUrl);
            fullUrl.search = qs.stringify({
                query: params.query.query,
                sources: Array.from(availableSources),
                categories: params.query.categories,
                price_min: params.query.price.min,
                price_max: params.query.price.max,
                ...pages,
            }, {
                arrayFormat: 'repeat',
                addQueryPrefix: true,
                skipNulls: true,
            });

            const es = new EventSource(fullUrl);
            
            error.value = false;
            pending.value = true;
            eof.value = false;

            // listeners
            es.addEventListener('posts', (e) => {
                const newData = JSON.parse(e.data) as Post[];
                
                posts.value = [...posts.value, ...newData];
            });

            es.addEventListener('pagination', (e) => {
                const newPagination = JSON.parse(e.data) as Pagination;
                
                if (! newPagination.has_next) {
                    eofSources.add(newPagination.source);
                }
            });

            es.addEventListener('close', () => {
                es.close();
                
                pending.value = false;
                error.value = false;
                
                if (eofSources.size === allSources.length) {
                    eof.value = true;
                }

                if (options.onFinish) {
                    options.onFinish();
                }
            });

            es.addEventListener('error', () => {
                es.close();
                
                error.value = true;
                pending.value = false;

                if (options.onFinish) {
                    options.onFinish();
                }
            });

            // close fn
            close = () => {
                if (es && es.readyState !== es.CLOSED) {
                    es.close();

                    eof.value = false;
                    error.value = false;
                    pending.value = false;
                }
            };
        };

        return {
            posts,
            error,
            pending,
            eof,
            execute,
            close,
            resetEofSources,
        };
    };

    return {
        sseFetch,
    };
}