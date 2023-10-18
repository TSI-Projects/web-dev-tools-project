<template>
    <div>
        <q-infinite-scroll
            :offset="250"
            :disable="error || eof"
            @load="onLoad"
        >
            <div
                v-if="result.length > 0"
                class="row q-col-gutter-md"
            >
                <div
                    v-for="post in result"
                    class="col-6 col-xl-2 col-lg-2 col-md-3 col-sm-4"
                >
                    <post-card
                        :title="post.title"
                        :price="post.price"
                        :preview-img="post.preview_img"
                        :url="post.url"
                        @navigate="navigateToPost"
                    />
                </div>
            </div>
            <template #loading>
                <div class="row justify-center q-my-md">
                    <q-spinner-dots
                        color="primary"
                        size="64px"
                    />
                </div>
            </template>
        </q-infinite-scroll>
        <template v-if="error">
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
                                @click="() => refetch()"
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
        </template>
        <client-only>
            <teleport to="#q-page-container">
                <q-page-sticky
                    id="q-page-sticky"
                    position="bottom-right"
                    :offset="[16, 16]"
                >
                    <post-filter-fab />
                </q-page-sticky>
            </teleport>
            <teleport to="#q-layout">
                <post-filter-drawer
                    v-model="parsedQuery"
                    :readonly="pending"
                />
            </teleport>
        </client-only>
    </div>
</template>

<script lang="ts" setup>
import { mdiAlertDecagram, mdiReload } from '@quasar/extras/mdi-v7';
import { QInfiniteScroll } from 'quasar';
import { FilterFields } from '~/components/Post/Filter/Drawer.vue';

useCustomSeoMeta({
    title: 'Главная',
    openGraph: {
        ogType: 'website',
    },
    twitterCard: {
        twitterCard: 'summary',
    },
});

const nuxtApp = useNuxtApp()
const route = useRoute();
const router = useRouter();
const posts = usePosts();

// According to this: https://nuxt.com/docs/migration/component-options#scrolltotop
// `scrollToTop` is not currently supported, so we are using this workaround instead.
nuxtApp.hook('page:finish', () => {
    window.scrollTo(0, 0);
});

const parsedQuery = computed<FilterFields>({
    get: () => {
        return {
            query: route.query.query as string | undefined,
            sources: route.query.sources as string[] | string | undefined,
            categories: route.query.categories as string[] | string | undefined,
            price: {
                /* @ts-ignore */
                min: route.query.price?.min as number | undefined,
                /* @ts-ignore */
                max: route.query.price?.max as number | undefined,
            },
        }
    },
    set: (newValue) => {
        router.push({
            name: 'index',
            /* @ts-ignore */
            query: {
                ...route.query,
                ...newValue,
            },
        });
    },
});

const { posts: result, eof, error, pending, close, execute, resetEofSources } = posts.sseFetch(() => {
    return {
        query: {
            query: parsedQuery.value.query,
            sources: parsedQuery.value.sources,
            categories: parsedQuery.value.categories,
            price: {
                min: parsedQuery.value.price.min,
                max: parsedQuery.value.price.max,
            },
        },
    }
});

const page = ref<number>(1);

watch(parsedQuery, () => {
    result.value = [];

    window.scrollTo(0, 0);

    resetEofSources();

    execute({
        page: page.value = 1,
    });
});

const onLoad: QInfiniteScroll['onLoad'] = (_, done) => {
    execute({
        page: page.value,
        onFinish: () => {
            page.value += 1;

            done();
        },
    })
};

const refetch = () => {
    execute({
        page: page.value,
    });
};

const navigateToPost = (url: string) => {
    window.open(url, '_blank');
};

onBeforeUnmount(() => close());
</script>