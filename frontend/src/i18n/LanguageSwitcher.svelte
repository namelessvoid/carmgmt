<script>
    import { locale } from 'svelte-i18n';

    const languages = ['de-DE', 'en'];
    let currentLanguage = $locale;

    function changeLanguage(language) {
        locale.set(language);
        currentLanguage = language;
    }
</script>

<style>
span {
    margin: 0 5px 0 5px;
}

.active {
    text-decoration: underline;
}
</style>

{#each languages as language}
<span on:click={() => changeLanguage(language)}
      class='{language === currentLanguage ? 'active' : ''}'>{language}</span>
{/each}