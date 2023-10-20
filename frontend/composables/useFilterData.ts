
export type Source = {
    id: string;
    name?: string;
};

export type DefaultFunctionResult = {
    fetchFilterData: () => Promise<FilterFetchDataResponse>
};

export type FilterFetchDataResponse = {
    sources: Source[];
};

export type FilterFetchSourcesResponse = {
    sources: string[];
};

export const SOURCE_NAMES: Record<string, string> = {
    ss: 'SS.LV',
    pp: 'PP.LV',
    banknote: 'BANKNOTE.LV',
};

export default function (): DefaultFunctionResult {
    const runtimeConfig = useRuntimeConfig();

    const fetchFilterData = async (): Promise<FilterFetchDataResponse> => {
        const sourcesPromise = $fetch<FilterFetchSourcesResponse>('/sources', {
            baseURL: runtimeConfig.public.api.baseUrl,
        });

        // in fututre we could add more filter data
        // that needs to be fetched from backend,
        // so i made this "workpiece" for future changes.
        const [sourcesResponse] = await Promise.all([
            sourcesPromise,
        ]);

        return {
            sources: sourcesResponse.sources.map((v) => ({
                id: v,
                name: SOURCE_NAMES[v] ?? undefined,
            })),
        };
    }

    return {
        fetchFilterData,
    }
}