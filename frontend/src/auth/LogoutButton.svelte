<script>
    import { isAuthenticated, logout } from './auth';
</script>

{#if $isAuthenticated }
    <a href="/" on:click={async () => await logout()}>Logout</a>
{/if}
