<template>
    <q-input
        :model-value="props.modelValue"
        bottom-slots
        standout="bg-primary text-white"
        counter
        clearable
        :lazy-rules="true"
        :maxlength="32"
        :readonly="props.readonly"
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

export type Props = {
    modelValue?: string;
    readonly?: boolean;
};

export type Emits = {
    (e: 'update:modelValue', query?: string): void;
};

const emits = defineEmits<Emits>();
const props = withDefaults(defineProps<Props>(), {
    modelValue: undefined,
    readonly: false,
});
</script>