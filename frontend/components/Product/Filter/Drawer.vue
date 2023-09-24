<template>
    <q-drawer
        v-model="drawerState"
        elevated
        show-if-above
        no-swipe-open
        :width="360"
        side="right"
        class="app-drawer"
    >
        <div class="app-drawer__prepend">
            <q-item>
                <q-item-section avatar>
                    <q-icon :name="mdiFilter" />
                </q-item-section>
                <q-item-section>
                    <q-item-label>Фильтр</q-item-label>
                </q-item-section>
            </q-item>
            <q-separator />
            <div class="q-pa-md">
                <div class="text-body2 q-mb-xs">
                    Поиск:
                </div>
                <product-filter-search-input
                    v-model="query"
                    :loading="loading"
                />
            </div>
            <q-separator />
        </div>
        <div class="app-drawer__content">
            <div class="q-pa-md">
                <div class="column q-col-gutter-md">
                    <div class="col">
                        <product-filter-source-select
                            v-model="sources"
                            :loading="props.loading"
                        />
                    </div>
                    <div class="col">
                        <product-filter-category-select
                            v-model="categories"
                            :loading="props.loading"
                        />
                    </div>
                    <div class="col">
                        <product-filter-price-range
                            v-model="price"
                            :loading="props.loading"
                        />
                    </div>
                </div>
            </div> 
        </div>
        <div class="app-drawer__append">
            <q-separator />
            <div class="q-pa-sm">
                <div class="row q-col-gutter-sm">
                    <div class="col-auto">
                        <q-btn
                            :icon="mdiUpdate"
                            :disable="loading"
                            color="accent"
                            @click.passive="clearFilters"
                        >
                            <q-tooltip
                                anchor="top middle"
                                self="bottom middle"
                            >
                                Сбросить
                            </q-tooltip>
                        </q-btn>
                    </div>
                    <div class="col-grow">
                        <q-btn
                            class="full-width"
                            label="Применить"
                            color="primary"
                            :loading="props.loading"
                            @click.passive="applyFilters"
                        />
                    </div>
                </div>
            </div>
        </div>
    </q-drawer>
</template>

<script lang="ts" setup>
import { mdiFilter, mdiUpdate } from '@quasar/extras/mdi-v7';
import { PriceRange } from './PriceRange.vue';

export type FilterFields = {
    query: string | undefined;
    categories: string[] | string | undefined;
    sources: string[] | string | undefined;
    price: {
        min: number | undefined;
        max: number | undefined;
    },
};

export type Props = {
    modelValue: FilterFields;
    loading?: boolean;
};

export type Emits = {
    (e: 'update:modelValue', modelValue: FilterFields): void;
};

const emits = defineEmits<Emits>();
const props = withDefaults(defineProps<Props>(), {
    loading: false,
});

const query = ref<string | undefined>(props.modelValue.query);
const sources = ref<string[] | string | undefined>(props.modelValue.sources);
const categories = ref<string[] | string | undefined>(props.modelValue.categories);
const price = ref<PriceRange>({ min: props.modelValue.price.min, max: props.modelValue.price.max });

const drawerState = useFilterDrawerState();

const applyFilters = () => {
    emits('update:modelValue', {
        query: query.value,
        categories: categories.value,
        sources: sources.value,
        price: {
            max: price.value.max,
            min: price.value.min,
        },
    });
};

const clearFilters = () => {
    query.value = undefined;
    sources.value = [];
    categories.value = [];
    price.value = { min: undefined, max: undefined };

    applyFilters();
};

watch(() => props.modelValue, (newModelValue) => {
    query.value = newModelValue.query;
    sources.value = newModelValue.sources;
    categories.value = newModelValue.categories;
    price.value = {
        min: newModelValue.price.min,
        max: newModelValue.price.max,
    };
});
</script>