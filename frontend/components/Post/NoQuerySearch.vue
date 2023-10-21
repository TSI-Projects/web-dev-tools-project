<template>
    <div class="row justify-center">
        <div class="col-12 col-xl-6 col-lg-6 col-md-6 col-sm-12">
            <q-card>
                <q-card-section>
                    <post-filter-search-input
                        v-model="query"
                        placeholder="RTX 3060"
                        @apply-filters="applySearch"
                    />
                </q-card-section>
                <q-separator />
                <q-card-actions>
                    <q-space />
                    <q-btn
                        color="primary"
                        @click="applySearch"
                    >
                        <q-icon
                            left
                            :name="mdiLayersSearch"
                        />
                        Применить
                    </q-btn>
                </q-card-actions>
            </q-card>
        </div>
    </div>
</template>

<script lang="ts" setup>
import { mdiLayersSearch } from '@quasar/extras/mdi-v7';

export type Emits = {
    (e: 'applySearch', query: string | undefined): void;
}

const emits = defineEmits<Emits>();

const query = ref<string>();

const applySearch = () => {
    emits('applySearch', query.value);
};
</script>