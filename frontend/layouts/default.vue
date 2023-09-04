<template>
    <q-layout view="lHh LpR fff">
        <q-header
            elevated
            class="bg-primary text-white"
        >
            <q-toolbar>
                <div class="test__test">
                    adasd
                </div>
                <q-toolbar-title
                    
                    shrink
                >
                    Title
                </q-toolbar-title>

                <q-btn
                    dense
                    flat
                    round
                    :icon="mdiMenu"
                    @click="toggleRightDrawer"
                />
            </q-toolbar>
        </q-header>

        <!-- <q-drawer
            v-model="leftDrawerOpen"
            show-if-above
            side="left"
            bordered
        >
            <q-item-label header>
                Объявления
            </q-item-label>
            <q-list>
                <template v-for="category in categories">
                    <q-item
                        v-for="item in category"
                        v-ripple
                        clickable
                    >
                        <q-item-section avatar>
                            <q-icon :name="item.icon" />
                        </q-item-section>
                        <q-item-section>
                            <q-item-label>{{ item.label }}</q-item-label>
                        </q-item-section>
                    </q-item>
                    <q-separator />
                </template>
            </q-list>
        </q-drawer> -->

        <q-drawer
            v-model="rightDrawerOpen"
            show-if-above
            side="right"
            bordered
        >
        <!-- drawer content -->
        </q-drawer>

        <q-page-container>
            <slot />
        </q-page-container>

        <q-footer
            elevated
            class="bg-primary text-white"
        >
        <!-- footer content -->
        </q-footer>
    </q-layout>
</template>

<script lang="ts" setup>
import { mdiMenu } from '@quasar/extras/mdi-v7';

const rightDrawerOpen = ref(false);

const toggleRightDrawer = () => {
    rightDrawerOpen.value = !rightDrawerOpen.value;
};
</script>