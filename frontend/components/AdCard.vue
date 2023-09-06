<template>
    <q-card
        v-ripple
        class="cursor-pointer"
        @click="goToAd"
    >
        <q-img :src="props.previewImg" />
        <q-card-section>
            <div class="text-body1 ad-title">
                {{ props.title }}
            </div>
            <div class="text-subtitle2 ad-description">
                {{ props.description }}
            </div>
            <div class="text-h6 text-weight-bold text-red">
                {{ props.price }}
            </div>
        </q-card-section>
    </q-card>
</template>

<script lang="ts" setup>
interface Props
{
    id: string;
    title: string;
    previewImg: string;
    description: string;
    price: string;
}

const props = defineProps<Props>();
const router = useRouter();

const goToAd = () => {
    router.push({ name: 'ads-id', params: { id: props.id } })
};
</script>

<style lang="scss" scoped>
@mixin lineClamp($lines) {
    overflow: hidden;
    display: -webkit-box;
    -webkit-line-clamp: $lines;
    -webkit-box-orient: vertical;
}

.ad-title
{
    @include lineClamp(1);
}

.ad-description
{
    @include lineClamp(2);
}
</style>