<template>
    <div class="text-body2 q-mb-xs">
        Источники:
    </div>
    <q-select
        :options="sources"
        standout="bg-primary text-white"
        clearable
        use-input
        multiple
        emit-value
        map-options
        :model-value="props.modelValue"
        :readonly="props.loading"
        @clear="() => emits('update:modelValue', [])"
        @filter="filterOptions"
        @update:model-value="(selected) => emits('update:modelValue', selected)"
    />
</template>

<script lang="ts" setup>
import { QSelect } from 'quasar';

export type Option = {
    label: string;
    value: string;
};

export type Props = {
    modelValue: string[] | string | undefined;
    loading?: boolean;
};

export type Emits = {
    (e: 'update:modelValue', modelValue: string[]): void;
};

const emits = defineEmits<Emits>();
const props = withDefaults(defineProps<Props>(), {
    modelValue: undefined,
    loading: false,
});

// TODO: add more options
const options: Option[] = [
    {
        label: 'SS.LV',
        value: 'sslv',
    },
];

const sources = ref<Option[]>(options);

const filterOptions: QSelect['onFilter'] = (value, update) => {
    update(() => {
        if (value.length === 0) {
            sources.value = options;
        } else {
            const needle = value.toLowerCase();

            sources.value = options.filter((v) => {
                return v.label.toLowerCase().indexOf(needle) > -1
                    || v.value.toLowerCase().indexOf(needle) > -1;
            });
        }
    });
};
</script>