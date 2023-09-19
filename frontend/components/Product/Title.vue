<template>
    <template v-if="hasQuery">
        <div class="text-h5">
            Найдено {{ props.count }} объявлений
        </div>
    </template>
    <template v-else>
        <div class="text-h5">
            Всего {{ props.count }} объявлений
        </div>
    </template>
</template>

<script lang="ts" setup>
export type Props = {
    count?: number;
    hasQuery?: boolean;
}

const props = withDefaults(defineProps<Props>(), {
    count: 0,
    hasQuery: false,
});
</script>