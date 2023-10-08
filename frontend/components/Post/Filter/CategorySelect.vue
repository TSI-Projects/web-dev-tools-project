<template>
    <div class="text-body2 q-mb-xs">
        Категории:
    </div>
    <q-select
        :options="categories"
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
const options = [
    {
        label: 'Компьютеры',
        value: 'computers',
    },
    {
        label: 'Телефоны',
        value: 'smartphones',
    },
    {
        label: 'Машины',
        value: 'cars',
    },
];

const categories = ref<Option[]>(options);

const filterOptions: QSelect['onFilter'] = (value, update) => {
    update(() => {
        if (value.length === 0) {
            categories.value = options;
        } else {
            const needle = value.toLowerCase();

            categories.value = options.filter((v) => {
                return v.label.toLowerCase().indexOf(needle) > -1
                    || v.value.toLowerCase().indexOf(needle) > -1;
            });
        }
    });
};
</script>