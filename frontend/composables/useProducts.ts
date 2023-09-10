export type Product = {
    id: string;
    title: string;
    preview_img: string;
    description: string;
    price: string;
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
                'acccept': 'application/json',
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