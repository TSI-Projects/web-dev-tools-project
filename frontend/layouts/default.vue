<template>
    <q-layout
        id="q-layout"
        view="lHh LpR fff"
    >
        <q-header
            elevated
            class="bg-primary text-white"
        >
            <q-toolbar>
                <q-toolbar-title>
                    Header
                </q-toolbar-title>
            </q-toolbar>
        </q-header>

        <q-page-container id="q-page-container">
            <slot />
        </q-page-container>

        <q-footer
            elevated
            class="bg-primary text-white"
        >
            Footer
        </q-footer>
    </q-layout>
</template>
