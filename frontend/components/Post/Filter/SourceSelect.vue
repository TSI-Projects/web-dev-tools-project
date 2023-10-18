<template>
    <div class="text-body2 q-mb-xs">
        Источники:
    </div>
    <q-select
        :options="options"
        standout="bg-primary text-white"
        clearable
        use-input
        multiple
        emit-value
        map-options
        :model-value="props.modelValue"
        :loading="props.loading"
        :readonly="props.readonly"
        @clear="() => emits('update:modelValue', [])"
        @filter="filterOptions"
        @update:model-value="(selected) => emits('update:modelValue', selected)"
    >
        <template #option="{ itemProps, selected, opt, toggleOption }">
            <q-item v-bind="itemProps">
                <q-item-section avatar>
                    <q-checkbox
                        :model-value="selected"
                        @click="() => toggleOption(opt)"
                    />
                </q-item-section>
                <q-item-section>
                    {{ opt.name || `N/A (${opt.id})` }}
                </q-item-section>
            </q-item>
        </template>
    </q-select>
</template>

<script lang="ts" setup>
import { QSelect } from 'quasar';

export type Option = {
    id: string;
    name?: string;
};

export type Props = {
    sources: Option[];
    modelValue: string[] | string | undefined;
    loading?: boolean;
    readonly?: boolean;
};

export type Emits = {
    (e: 'update:modelValue', modelValue: string[]): void;
};

const emits = defineEmits<Emits>();
const props = withDefaults(defineProps<Props>(), {
    modelValue: undefined,
    readonly: false,
    loading: false,
});

const filterData = useFilterData();

const options = ref<Option[]>(props.sources);

const filterOptions: QSelect['onFilter'] = (value, update) => {
    update(() => {
        if (value.length === 0) {
            options.value = props.sources;
        } else {
            const needle = value.toLowerCase();

            options.value = props.sources.filter((v) => {
                return v.id.toLowerCase().indexOf(needle) > -1
                    || (v.name && v.name?.toLowerCase().indexOf(needle) > -1);
            });
        }
    });
};
</script>