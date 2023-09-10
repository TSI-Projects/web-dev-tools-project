<template>
    <q-input
        ref="inputRef"
        v-model="query"
        standout
        bottom-slots
        counter
        :lazy-rules="true"
        :maxlength="32"
        :readonly="props.loading"
        @keydown="onEnterKeyDown"
    >
        <template #append>
            <q-btn
                v-if="props.modelValue"
                :disable="props.loading"
                round
                flat
                dense
                @click="clearQuery"
            >
                <q-icon :name="mdiHistory" />
            </q-btn>
            <q-btn
                :loading="props.loading"
                :disable="inputRef?.hasError || false"
                flat
                @click="updateModelValue"
            >
                <q-icon
                    left
                    :name="mdiMagnify"
                />Найти
            </q-btn>
        </template>
    </q-input>
</template>

<script lang="ts" setup>
import { mdiHistory, mdiMagnify } from '@quasar/extras/mdi-v7';

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

const inputRef = ref();
const query = ref(props.modelValue);

const onEnterKeyDown = (event: KeyboardEvent) => {
    if (event.code === 'Enter') {
        updateModelValue();
    }
};

const clearQuery = () => {
    query.value = '';

    updateModelValue();
};

const updateModelValue = () => {
    inputRef.value?.blur();

    const finalQuery = (query.value && query.value.length > 0)
        ? query.value
        : undefined;

    emits('update:modelValue', finalQuery);
};
</script>