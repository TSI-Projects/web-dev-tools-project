<template>
    <q-btn
        v-show="shouldBeShown"
        fab
        :icon="mdiFilter"
        color="secondary"
        @click="showDrawer"
    />
</template>

<script lang="ts" setup>
import { mdiFilter } from '@quasar/extras/mdi-v7';
import { useQuasar } from 'quasar';

export type Props = {
    breakpoint?: number;
};

const props = withDefaults(defineProps<Props>(), {
    breakpoint: 1023,
});

const $q = useQuasar();
//const $productsFilter = useProductsFilter();

const shouldBeShown = computed(() => $q.screen.width <= props.breakpoint);

const showDrawer = () => {
    //$productsFilter.sharedState.value = true;
};
</script>