<template>
    <q-input
        :model-value="props.modelValue"
        bottom-slots
        standout="bg-primary text-white"
        counter
        clearable
        :lazy-rules="true"
        :maxlength="32"
        :readonly="props.loading"
        @clear="() => emits('update:modelValue', undefined)"
        @update:model-value="(value) => emits('update:modelValue', value as string | undefined)"
    >
        <template #prepend>
            <q-icon :name="mdiMagnify" />
        </template>
    </q-input>
</template>

<script lang="ts" setup>
import { mdiMagnify } from '@quasar/extras/mdi-v7';

const emits = defineEmits<{
    (e: 'update:modelValue', query?: string): void;
}>();

const props = withDefaults(defineProps<{
    modelValue?: string;
    loading?: boolean;
}>(), {
    modelValue: undefined,
    loading: false,
});
</script>