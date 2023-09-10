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
const props = withDefaults(defineProps<{
    count?: number;
    hasQuery?: boolean;
}>(), {
    count: 0,
    hasQuery: false,
});
</script>