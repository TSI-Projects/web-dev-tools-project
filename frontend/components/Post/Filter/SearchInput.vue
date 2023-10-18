<template>
    <q-input
        ref="inputRef"
        :model-value="props.modelValue"
        bottom-slots
        standout="bg-primary"
        counter
        clearable
        :lazy-rules="true"
        :maxlength="32"
        :readonly="props.loading"
        @clear="() => updateModelValue(undefined)"
        @update:model-value="updateModelValue"
        @keydown="onKeyDown"
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
    loading?: boolean;
};

export type Emits = {
    (e: 'update:modelValue', query?: string): void;
    (e: 'applyFilters'): void;
};

const emits = defineEmits<Emits>();
const props = withDefaults(defineProps<Props>(), {
    modelValue: undefined,
    loading: false,
});

const inputRef = ref();

const onKeyDown = (event: KeyboardEvent) => {
    if (event.code === 'Enter') {
        inputRef.value?.blur();

        emits('applyFilters');
    }
};

const updateModelValue = (value: string | number | null | undefined = undefined) => {
    if (value) {
        value = value.toString();
        
        if (value.length === 0) {

            emits('update:modelValue', undefined);
        } else {

            emits('update:modelValue', value);
        }
    } else {

        emits('update:modelValue', undefined);
    }
}
</script>