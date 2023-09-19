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

const quasar = useQuasar();
const filterDrawerState = useFilterDrawerState();

const shouldBeShown = computed(() => quasar.screen.width <= props.breakpoint);

const showDrawer = () => {
    filterDrawerState.value = true;
};
</script>