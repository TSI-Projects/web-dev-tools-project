<template>
    <div class="text-body2 q-mb-xs">
        Цена:
    </div>
    <div class="row q-col-gutter-md">
        <div class="col-6">
            <q-input
                v-model.trim.number="min"
                standout="bg-primary text-white"
                type="number"
                clearable
                :readonly="props.readonly"
                label="Мин."
                @clear="() => min = undefined"
            />
        </div>
        <div class="col-6">
            <q-input
                v-model.trim.number="max"
                standout="bg-primary text-white"
                type="number"
                clearable
                :readonly="props.readonly"
                label="Макс."
                @clear="() => max = undefined"
            />
        </div>
    </div>
</template>

<script lang="ts" setup>
export type PriceRange = {
    min: number | undefined;
    max: number | undefined;
};

export type Props = {
    readonly?: boolean;
    modelValue?: PriceRange;
};

export type Emits = {
    (e: 'update:modelValue', price: PriceRange): void;
};

const emits = defineEmits<Emits>();
const props = withDefaults(defineProps<Props>(), {
    modelValue: () => ({
        min: undefined,
        max: undefined,
    }),
    readonly: false,
});

const min = ref<number | undefined>(props.modelValue.min);
const max = ref<number | undefined>(props.modelValue.max);

watch([min, max], ([newMin, newMax]) => {
    if ((newMin && newMax) && newMin > newMax) {
        nextTick(() => {
            min.value = newMax;
        });
    }
    
    emits('update:modelValue', {
        min: newMin,
        max: newMax,
    });
}, {
    immediate: true,
});

watch(() => props.modelValue, (newModelValue) => {
    min.value = newModelValue.min;
    max.value = newModelValue.max;
});
</script>