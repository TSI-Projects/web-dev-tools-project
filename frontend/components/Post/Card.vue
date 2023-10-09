<template>
    <q-card
        v-ripple
        class="q-hoverable cursor-pointer"
        href="http://localhost"
        @click="navigateToPost"
        @click.middle="navigateToPost"
        @mousedown.middle.prevent.stop="null"
    >
        <div
            tabindex="-1"
            class="q-focus-helper"
        />
        <q-img :src="props.previewImg" />
        <q-card-section>
            <div
                class="text-body1 text-weight-bold line-clamp"
                style="--line-clamp: 1;"
            >
                {{ props.title }}
                <q-tooltip max-width="300px">
                    {{ props.title }}
                </q-tooltip>
            </div>
            <div
                class="text-subtitle2 line-clamp"
                style="--line-clamp: 2;"
            >
                {{ props.description }}
            </div>
            <div class="text-h6 text-weight-bold text-red">
                {{ props.price }}
            </div>
        </q-card-section>
        <q-separator />
        <q-card-actions class="align-center">
            <q-icon
                :name="mdiWeb"
                size="xs"
                class="q-mr-sm"
            />
            <span class="text-uppercase">{{ sourceHost }}</span>
        </q-card-actions>
    </q-card>
</template>

<script lang="ts" setup>
import { mdiWeb } from '@quasar/extras/mdi-v7';

export type Props = {
    previewImg: string;
    title: string;
    description: string;
    price: string;
    url: string;
};

export type Emits = {
    (e: 'navigate', url: string): void,
};

const emits = defineEmits<Emits>();
const props = defineProps<Props>();

const navigateToPost = () => {
    emits('navigate', props.url);
};

const sourceHost = new URL(props.url).hostname.replaceAll('www.', '');
</script>
