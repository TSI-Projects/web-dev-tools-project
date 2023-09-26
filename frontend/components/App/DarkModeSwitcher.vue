<template>
    <q-btn
        flat
        :icon="mdiThemeLightDark"
    >
        <q-menu>
            <q-item
                v-for="option in options"
                v-close-popup
                clickable
                :active="option.value === quasar.dark.mode"
                @click="switchTheme(option.value)"
            >
                <q-item-section avatar>
                    <q-icon :name="option.icon" />
                </q-item-section>
                <q-item-section>{{ option.name }}</q-item-section>
            </q-item>
        </q-menu>
    </q-btn>
</template>

<script lang="ts" setup>
import { useQuasar } from 'quasar';
import { mdiWeatherNight, mdiLightbulbOnOutline, mdiLaptop, mdiThemeLightDark } from '@quasar/extras/mdi-v7';

export type Theme = {
    value: boolean | 'auto';
    name: string,
    icon: string;
};

const quasar = useQuasar();

const options: Theme[] = [
    {
        value: 'auto',
        name: 'Системная',
        icon: mdiLaptop,
    },
    {
        value: false,
        name: 'Светлая',
        icon: mdiLightbulbOnOutline,
    },
    {
        value: true,
        name: 'Тёмная',
        icon: mdiWeatherNight,
    },
];

const switchTheme = (theme: boolean | 'auto') => {
    quasar.dark.set(theme);
};
</script>