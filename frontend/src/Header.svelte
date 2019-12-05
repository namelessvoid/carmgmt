<script>
    import LanguageSwitcher from './i18n/LanguageSwitcher.svelte';
    import LogoutButton from './auth/LogoutButton.svelte';
</script>

<style>
header {
    height: 48px;
    padding: 0 8px 0 8px;
    background-color: lightgreen;
    display: flex;
}

span {
    margin-top: auto;
    margin-bottom: auto;
}

.logo {
    font-weight: bold;
    font-size: 18px;
    flex-grow: 1;
}
</style>

<header>
<span class="logo">FleetMgmt</span>
<span><LanguageSwitcher /></span>
<span><LogoutButton /></span>
</header>