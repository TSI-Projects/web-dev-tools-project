<template>
    <q-drawer
        v-model="drawer"
        elevated
        show-if-above
        no-swipe-open
        side="right"
        class="app-drawer"
    >
        <div class="app-drawer__content">
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
                    <div class="col-2">
                        <q-btn
                            class="full-width"
                            :icon="mdiUpdate"
                            :disable="loading"
                            color="accent"
                            @click.passive="clearFilters"
                        />
                    </div>
                    <div class="col-10">
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
import { Price } from './PriceRange.vue';

export type Props = {
    modelValue: ParsedQuery;
    loading?: boolean;
};

export type Emits = {
    (e: 'update:modelValue', modelValue: ParsedQuery): void;
};

const emits = defineEmits<Emits>();
const props = withDefaults(defineProps<Props>(), {
    loading: false,
});

const price = ref<Price>({
    from: props.modelValue.price.from,
    to: props.modelValue.price.to,
});
const sources = ref<string[]>(props.modelValue.sources);
const categories = ref<string[]>(props.modelValue.categories);

const drawer = ref();

const applyFilters = () => {
    emits('update:modelValue', {
        query: props.modelValue.query,
        sources: sources.value,
        categories: categories.value,
        price: {
            from: price.value.from,
            to: price.value.to,
        },
    });
};

const clearFilters = () => {
    sources.value = [];
    categories.value = [];
    price.value = {
        from: undefined,
        to: undefined,
    };

    emits('update:modelValue', {
        query: props.modelValue.query,
        sources: [],
        categories: [],
        price: {
            from: undefined,
            to: undefined,
        },
    });
};
</script>