import qs from 'qs';

export type Post = {
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
    posts: Ref<Post[]>;
    error: Ref<boolean>;
    pending: Ref<boolean>;
    eof: Ref<boolean>;
    execute: (done?: () => void) => void;
    close: () => void;
};

export default function () {
    const { public: publicConfig } = useRuntimeConfig();
    const { wrap } = useArray();

    const sseFetch = (params: FetchAllParameters): SseFetchResult => {
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
            eof.value = false;

            // listeners
            es.addEventListener('message', (e) => {
                const newData = JSON.parse(e.data) as Post[];

                if (newData.length === 0) {
                    es.close();

                    eof.value = true;
                    pending.value = false;

                    return;
                }
                
                posts.value = [...posts.value, ...newData];
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
        };
    };
    
    const fetchOne = () => {
        // TODO: implement when the backend is ready
    };

    return {
        sseFetch,
        fetchOne,
    };
}