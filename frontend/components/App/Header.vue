<template>
    <q-header
        elevated
        class="bg-primary text-white"
    >
        <q-toolbar>
            <q-toolbar-title>
                {{ appName }}
            </q-toolbar-title>
        </q-toolbar>
    </q-header>
</template>

<script lang="ts" setup>
const { appName } = useAppConfig();
</script>