<template>
    <q-header
        elevated
        class="bg-primary text-white"
    >
        <q-toolbar>
            <q-toolbar-title>
                {{ appConfig.appName }}
            </q-toolbar-title>
            <q-space />
            <app-dark-mode-switcher />
        </q-toolbar>
    </q-header>
</template>

<script lang="ts" setup>
const appConfig = useAppConfig();
</script>