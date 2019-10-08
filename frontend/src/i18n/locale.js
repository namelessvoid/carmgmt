import { locale, getClientLocale } from 'svelte-i18n';

locale.set('en');

locale.set(
    getClientLocale({
        navigator: false,
        search: 'loc',
        fallback: 'en',
    }),
);
  
locale.subscribe(l => {
    console.log('locale change', l)
});