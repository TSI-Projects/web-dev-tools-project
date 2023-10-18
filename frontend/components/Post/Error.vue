<template>
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
                        @click="() => emits('refresh')"
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

<script lang="ts" setup>
import { mdiAlertDecagram } from '@quasar/extras/mdi-v7';
import { mdiReload } from '@quasar/extras/mdi-v7';

export type Emits = {
    (e: 'refresh'): void;
};

const emits = defineEmits<Emits>();
</script>