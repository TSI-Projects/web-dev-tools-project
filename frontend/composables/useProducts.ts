import qs from 'qs';

export type Product = {
    id: string;
    title: string;
    preview_img: string;
    description: string;
    price: string;
    url: string;
};

export type FetchAllParameters = {
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

export type SseFetchResult = {
    products: Ref<Product[] | undefined>;
    error: Ref<boolean>;
    pending: Ref<boolean>;
    execute: (done?: () => void) => void;
    close: () => void;
};

export default function () {
    const { public: publicConfig } = useRuntimeConfig();
    const { wrap } = useArray();

    const sseFetch = (params: FetchAllParameters): SseFetchResult => {
        const pending = ref<boolean>(false);
        const error = ref<boolean>(false);
        const products = ref<Product[]>();
        
        // EventSource object does not exists on server,
        // so just return empty state.
        if (process.server) {
            return {
                products,
                pending,
                error,
                execute: () => { },
                close: () => { },
            };
        }

        const baseUrl = publicConfig.api.baseUrl || 'http://localhost';

        const query = qs.stringify({
            product: params.query.query,
            source: wrap<string, undefined>(params.query.sources, undefined),
            categories: wrap<string, undefined>(params.query.categories, undefined),
            price: {
                min: params.query.price.min,
                max: params.query.price.max,
            },
        }, {
            arrayFormat: 'repeat',
            addQueryPrefix: true,
            skipNulls: true,
        });

        const fullUrl = `${baseUrl}/search${query}`;

        let close: () => void = () => { };

        const execute = (done?: () => void): void  => {
            if (pending.value) {
                close();
            }

            const es = new EventSource(fullUrl);
            
            error.value = false;
            pending.value = true;

            // listeners
            es.addEventListener('message', (e) => {
                const newData = JSON.parse(e.data) as Product[];

                if (products.value) {
                    products.value = [...products.value, ...newData];
                } else {
                    products.value = newData;
                }
            });

            es.addEventListener('close', () => {
                es.close();
                
                pending.value = false;
                error.value = false;

                if (done) {
                    done();
                }
            });

            es.addEventListener('error', () => {
                es.close();
                
                error.value = true;
                pending.value = false;

                if (done) {
                    done();
                }
            });

            // close fn
            close = () => {
                if (es && es.readyState !== es.CLOSED) {
                    es.close();

                    error.value = false;
                    pending.value = false;
                }
            };
        };

        return {
            products,
            pending,
            error,
            execute,
            close,
        };
    };

    const fetchAll = (params: FetchAllParameters): Promise<Product[]> => {
        return $fetch<Product[]>('/search', {
            baseURL: publicConfig.api.baseUrl || 'http://localhost',
            params: {
                product: params.query.query,
                sources: wrap<string, undefined>(params.query.sources, undefined),
                categories: wrap<string, undefined>(params.query.categories, undefined),
                price: {
                    min: params.query.price.min,
                    max: params.query.price.max,
                },
            },
            onRequest: (ctx) => {
                if (ctx.options.params || ctx.options.query) {
                    ctx.request = ctx.request + qs.stringify({
                        ...ctx.options.query,
                        ...ctx.options.params,
                    }, {
                        arrayFormat: 'brackets',
                        addQueryPrefix: true,
                        skipNulls: true,
                    });

                    ctx.options.params = undefined;
                    ctx.options.query = undefined;
                }
            },
            headers: {
                'accept': 'application/json',
            },
            // TODO: delete this when backend is ready
            parseResponse(string) {
                return JSON.parse(string);
            },
        });
    };
    
    const fetchOne = () => {
        // TODO: implement when the backend is ready
    };

    return {
        sseFetch,
        fetchAll,
        fetchOne,
    };
}