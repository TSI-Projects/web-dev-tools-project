<template>
    <q-page padding>
        <div class="column q-col-gutter-md">
            <div class="col-auto">
                <div class="row justify-center">
                    <div class="col-12 col-xl-5 col-lg-5 col-md-6 col-sm-12">
                        <product-search-input
                            :model-value="parsedQuery.query"
                            :loading="status === 'pending'"
                            @update:model-value="(query) => parsedQuery = { ...parsedQuery, query }"
                        />
                        {{ parsedQuery }}
                    </div>
                </div>
            </div> 
            <template v-if="status === 'success'">
                <div class="col-auto">
                    <product-title
                        :count="0"
                        :has-query="!!parsedQuery.query"
                    />
                </div>
                <div class="col-auto">
                    <div class="row q-col-gutter-md">
                        <div
                            v-for="product in result"
                            class="col-6 col-xl-2 col-lg-2 col-md-3 col-sm-4"
                        >
                            <product-card
                                :title="product.title"
                                :description="product.description"
                                :price="product.price"
                                :preview-img="product.preview_img"
                                :url="product.url"
                                @navigate="navigateToProduct"
                            />
                        </div>
                    </div>
                </div>
            </template>
            <template v-else-if="status === 'error'">
                <div class="col-auto">
                    <div class="row justify-center">
                        <div class="col-12 col-xl-6 col-lg-6 col-md-6 col-sm-12">
                            <q-banner
                                class="text-white bg-red-10 shadow-2"
                                rounded
                            >
                                <template #avatar>
                                    <q-icon :name="mdiAlertDecagram" />
                                </template>
                                <template #action>
                                    <q-btn
                                        flat
                                        @click="() => refresh()"
                                    >
                                        <q-icon
                                            left
                                            :name="mdiReload"
                                        /> Повторить
                                    </q-btn>
                                </template>
                                Ошибка загрузки данных с сервера.
                            </q-banner>
                        </div>
                    </div>
                </div>
            </template>
            <template v-else-if="status === 'pending'">
                <div class="col-auto self-center">
                    <q-spinner-grid
                        color="primary"
                        size="128px"
                    />
                </div>
            </template>
        </div>
        <!-- TELEPORT -->
        <client-only>
            <teleport to="#q-page-container">
                <q-page-sticky
                    id="q-page-sticky"
                    position="bottom-right"
                    :offset="[16, 16]"
                >
                    <product-filter-fab />
                </q-page-sticky>
            </teleport>
            <teleport to="#q-layout">
                <product-filter-drawer
                    v-model="parsedQuery"
                    :loading="status === 'pending'"
                />
            </teleport>
        </client-only>
    </q-page>
</template>

<script lang="ts" setup>
import { mdiAlertDecagram, mdiReload } from '@quasar/extras/mdi-v7';
import { LocationQuery } from '#vue-router';

const route = useRoute();
const router = useRouter();
const products = useProducts();
const array = useArray();

let skipRouteWatcher = false;
let skipParseQueryWatcher = false;

const navigateToProduct = (url: string) => {
    router.push({
        name: 'products-url',
        params: {
            url: btoa(url),
        },
    });
};

const updateQuery = (params: any) => {
    router.push({
        name: 'index',
        query: {
            ...route.query,
            ...params,
        },
    });
};

const parseQuery = (query: LocationQuery): Ref<ParsedQuery> => {
    return ref<ParsedQuery>({
        query: query?.query as string | undefined,
        sources: array.wrap<string>(query?.sources),
        categories: array.wrap<string>(query?.categories),
        price: {
            from: query?.price?.from,
            to: query?.price?.to,
        },
    });
};

const parsedQuery = parseQuery(route.query);

const { data: result, status, refresh } = useAsyncData('products', () => products.fetchAll({
    query: {
        query: parsedQuery.value.query,
        sources: parsedQuery.value.sources,
        categories: parsedQuery.value.categories,
        price: {
            from: parsedQuery.value.price.from,
            to: parsedQuery.value.price.to,
        },
    },
}), {
    watch: [parsedQuery],
});

watch(() => route.query, (newRouteQuery) => {
    if (skipRouteWatcher) {
        skipRouteWatcher = false;
        return;
    }

    skipParseQueryWatcher = true;

    parsedQuery.value = parseQuery(newRouteQuery).value;
});

watch(parsedQuery, () => {
    if (skipParseQueryWatcher) {
        skipParseQueryWatcher = false;
        return;
    }

    skipRouteWatcher = true;
    
    updateQuery(parsedQuery.value);
});
</script>