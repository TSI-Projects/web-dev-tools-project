<template>
    <q-input
        ref="inputRef"
        :model-value="props.modelValue"
        bottom-slots
        :label="props.label"
        standout="bg-primary"
        counter
        clearable
        :lazy-rules="true"
        :maxlength="32"
        :readonly="props.readonly"
        :loading="props.loading"
        :shadow-text="shadowText"
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
    readonly?: boolean;
    placeholder?: string;
    label?: string;
};

export type Emits = {
    (e: 'update:modelValue', query?: string): void;
    (e: 'applyFilters'): void;
};

const emits = defineEmits<Emits>();
const props = withDefaults(defineProps<Props>(), {
    modelValue: undefined,
    loading: false,
    readonly: false,
    placeholder: undefined,
    label: undefined,
});

const inputRef = ref();
const shadowText = ref<string | undefined>(props.placeholder);

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
            shadowText.value = props.placeholder;

            emits('update:modelValue', undefined);
        } else {
            shadowText.value = undefined;

            emits('update:modelValue', value);
        }
    } else {
        shadowText.value = props.placeholder;

        emits('update:modelValue', undefined);
    }
}
</script>