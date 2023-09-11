<template>
    <q-page padding>
        <div class="column q-col-gutter-md">
            <div class="col-auto">
                <div class="row justify-center">
                    <div class="col-12 col-xl-5 col-lg-5 col-md-6 col-sm-12">
                        <product-search-input
                            v-model="query"
                            :loading="status === 'pending'"
                        />
                    </div>
                </div>
            </div>
            <template v-if="status === 'success'">
                <div class="col-auto">
                    <product-title
                        :count="0"
                        :has-query="!!query"
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
                                        @click="refresh()"
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
    </q-page>
</template>

<script lang="ts" setup>
import { mdiAlertDecagram, mdiReload } from '@quasar/extras/mdi-v7';

const router = useRouter();
const route = useRoute();
const products = useProducts();

const query = ref(route.query?.query as string | undefined);

const { data: result, status, refresh } = await products.paginate({
    query,
});

const navigateToProduct = (url: string) => {
    router.push({
        name: 'products-url',
        params: {
            url: btoa(url),
        },
    });
};

watch(query, (value) => {
    router.push({
        name: 'index',
        query: {
            query: value,
        },
    });
});
</script>