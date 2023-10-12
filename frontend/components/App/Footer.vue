<template>
    <q-footer
        elevated
        class="bg-primary text-white"
    >
        <div class="footer__top q-pa-md">
            <div class="row q-col-gutter-md justify-center">
                <div class="col-12 col-sm-6 col-md-12 col-lg-5 col-xl-4">
                    <div class="text-h6 q-mb-md">
                        О Нас
                    </div>
                    {{ aboutUs }}
                </div>
                <div class="col-12 col-sm-6 col-md-12 col-lg-5 col-xl-4">
                    <div class="text-h6 q-mb-sm">
                        Ссылки
                    </div>
                    <div class="row q-gutter-md">
                        <q-btn
                            :icon="mdiGithub"
                            flat
                            round
                            dense
                            @click="navigateToSourceCode"
                        >
                            <q-tooltip
                                anchor="top middle"
                                self="bottom middle"
                            >
                                Исходный код
                            </q-tooltip>
                        </q-btn>
                    </div>
                </div>
            </div>
        </div>
        <q-separator dark />
        <div class="footer__bottom q-pa-md flex">
            <div>
                © 2023
            </div>
            <q-space />
            <div>
                Created by Vadim and Andrey 
            </div>
        </div>
    </q-footer>
</template>

<script lang="ts" setup>
import { mdiGithub } from '@quasar/extras/mdi-v7';

const appConfig = useAppConfig();

const navigateToSourceCode = () => {
    window.open(appConfig.sourceCodeLink, '_blank');
};

const aboutUs = appConfig.footerAboutUs.replace('{{ siteName }}', appConfig.appName);
</script>