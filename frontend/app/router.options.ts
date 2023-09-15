import type { RouterConfig } from '@nuxt/schema';
import qs from 'qs';

export default <RouterConfig>{
    stringifyQuery: (query) => {
        return qs.stringify(query, {
            arrayFormat: 'comma',
            allowDots: true,
            skipNulls: true,
        });
    },
    parseQuery: (search) => {
        return qs.parse(search, {
            comma: true,
            allowDots: true,
        });
    },
};