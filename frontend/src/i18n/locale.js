import { locale, getClientLocale } from 'svelte-i18n';

locale.set('en');

locale.set(
    getClientLocale({
        navigator: false,
        fallback: 'en',
    }),
);
