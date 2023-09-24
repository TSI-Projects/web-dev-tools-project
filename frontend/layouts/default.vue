<template>
    <q-layout
        id="q-layout"
        view="lHh LpR ffr"
    >
        <app-header />

        <q-page-container id="q-page-container">
            <q-page padding>
                <slot />
            </q-page>
        </q-page-container>

        <app-footer />
    </q-layout>
</template>
