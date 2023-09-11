export type Product = {
    id: string;
    title: string;
    preview_img: string;
    description: string;
    price: string;
    url: string;
}

export type FetchAllOptions = {
    query: string | undefined | Ref<string | undefined>;
}

export default function () {
    const { public: publicConfig } = useRuntimeConfig();

    const fetchAll = (options: FetchAllOptions) => {
        return useFetch<Product[]>('/search', {
            baseURL: publicConfig.api.baseUrl || 'http://localhost',
            params: {
                product: options.query || 'rtx 3060',
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
        paginate: fetchAll,
        find: fetchOne,
    };
}