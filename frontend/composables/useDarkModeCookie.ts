import { useQuasar } from "quasar";

export type Preference = boolean | 'auto';

export const COOKIE_KEY = 'quasar_dark_mode';

export default function () {
    const quasar = useQuasar();

    const get = (defaultValue: Preference = 'auto'): Preference => {
        const preference = quasar.cookies.get(COOKIE_KEY);

        if (!preference) {
            set(defaultValue);
            
            return defaultValue;
        }

        if (['true', 'false', 'auto'].includes(preference)) {
            if (preference === 'auto') {
                return preference;
            }

            return preference === 'true'
                ? true
                : false;
        }

        set(defaultValue);

        return defaultValue;
    };

    const set = (theme: Preference): void => {
        quasar.cookies.set(COOKIE_KEY, String(theme), {
            sameSite: 'Lax',
            secure: true,
            expires: 365,
        });
    };

    return {
        get,
        set,
    }
}