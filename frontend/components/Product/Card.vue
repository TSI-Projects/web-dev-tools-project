<template>
    <q-card
        v-ripple
        class="cursor-pointer"
        @click="goToAd"
    >
        <q-img :src="props.previewImg" />
        <q-card-section>
            <div
                class="text-body1 text-weight-bold line--clamp"
                style="--line-clamp: 1;"
            >
                {{ props.title }}
                <q-tooltip max-width="300px">
                    {{ props.title }}
                </q-tooltip>
            </div>
            <div
                class="text-subtitle2 line--clamp"
                style="--line-clamp: 2;"
            >
                {{ props.description }}
            </div>
            <div class="text-h6 text-weight-bold text-red">
                {{ props.price }}
            </div>
        </q-card-section>
    </q-card>
</template>

<script lang="ts" setup>
const props = defineProps<{
    id: string;
    previewImg: string;
    title: string;
    description: string;
    price: string;
}>();

const router = useRouter();

const goToAd = () => {
    router.push({ name: 'ads-id', params: { id: props.id } })
};
</script>

<style lang="scss" scoped>
.line--clamp {
    overflow: hidden;
    display: -webkit-box;
    -webkit-line-clamp: var(--line-clamp, 2);
    -webkit-box-orient: vertical;
}
</style>