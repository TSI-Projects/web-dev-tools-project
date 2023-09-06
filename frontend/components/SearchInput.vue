<template>
    <q-input
        v-model="query"
        standout
        dense
        @keydown="onKeyDown"
    >
        <template #append>
            <q-icon
                class="cursor-pointer"
                :name="mdiMagnify"
                @click="performSearch"
            /> 
        </template>
    </q-input>
</template>

<script lang="ts" setup>
import { mdiMagnify } from '@quasar/extras/mdi-v7';

interface Emits
{
    (e: 'performSearch', query?: string): void
}

const emits = defineEmits<Emits>();

const query = ref('');

const onKeyDown = (event: KeyboardEvent) => {
    if (event.code === 'Enter') {
        performSearch();
    }
};

const performSearch = () => {
    emits('performSearch', query.value);
};
</script>