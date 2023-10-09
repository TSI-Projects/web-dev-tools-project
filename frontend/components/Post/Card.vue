<template>
    <q-card
        v-ripple
        class="q-hoverable cursor-pointer"
        href="http://localhost"
        @click="navigateToPost"
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
            <div class="text-h6 text-weight-bold text-red">
                {{ props.price }}
            </div>
        </q-card-section>
    </q-card>
</template>

<script lang="ts" setup>
export type Props = {
    previewImg: string;
    title: string;
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
</script>
