import type { RouterConfig } from '@nuxt/schema';
import qs from 'qs';

export default <RouterConfig>{
    stringifyQuery: (query) => {
        return qs.stringify(query, {
            arrayFormat: 'brackets',
            skipNulls: true,
        });
    },
    parseQuery: (search) => {
        return qs.parse(search);
    },
};