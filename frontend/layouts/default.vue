<template>
    <q-layout view="lHh LpR fff">
        <q-header
            elevated
            class="bg-primary text-white"
        >
            <q-toolbar>
                <q-toolbar-title>
                    Title
                </q-toolbar-title>
                <search-input
                    style="width: 32vw;"
                    @perform-search="onSearchAds"
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

        <!-- <q-drawer
            v-model="rightDrawerOpen"
            show-if-above
            side="right"
            bordered
        >

        </q-drawer> -->

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
const router = useRouter();

const onSearchAds = (query?: string) => {
    router.push({ name: 'index', query: { query: query } });
}
</script>