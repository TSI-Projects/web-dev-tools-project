import { useQuasar } from 'quasar';

export default defineNuxtPlugin(() => {
    const quasar = useQuasar();
    const darkModeCookie = useDarkModeCookie();
    
    quasar.dark.set(
        darkModeCookie.get()
    );

    watch(() => quasar.dark.mode, (newValue) => {
        darkModeCookie.set(newValue);
    });
});