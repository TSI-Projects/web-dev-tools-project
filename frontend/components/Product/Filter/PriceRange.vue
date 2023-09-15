<template>
    <div class="text-body2 q-mb-xs">
        Цена:
    </div>
    <div class="row q-col-gutter-md">
        <div class="col-12">
            <q-input
                label="От"
                standout="bg-primary text-white"
                clearable
                type="number"
                :readonly="props.loading"
                :model-value="props.modelValue.from"
                @update:model-value="(value) => updateModelValue(value, 'from')"
            />
        </div>
        <div class="col-12">
            <q-input
                label="До"
                standout="bg-primary text-white"
                clearable
                type="number"
                :readonly="props.loading"
                :model-value="props.modelValue.to"
                @update:model-value="(value) => updateModelValue(value, 'to')"
            />
        </div>
    </div>
</template>


<script lang="ts" setup>
export type Price = {
    from?: number;
    to?: number;
};

export type Props = {
    loading?: boolean;
    modelValue?: Price;
};

export type Emits = {
    (e: 'update:modelValue', price: Price): void;
};

const props = withDefaults(defineProps<Props>(), {
    modelValue: () => ({
        from: undefined,
        to: undefined,
    }),
    loading: false,
});

const emits = defineEmits<Emits>();

const updateModelValue = (value: string | number | null, type: 'from' | 'to'): void => {
    if (value) {
        let possibleNumber: number | undefined = Number(value);

        if (isNaN(possibleNumber)) {
            possibleNumber = undefined;
        }

        emits('update:modelValue', {
            ...props.modelValue,
            [type]: possibleNumber,
        });
    } else {
        emits('update:modelValue', {
            ...props.modelValue,
            [type]: undefined,
        });
    }
};
</script>