<script>
    import { _ } from 'svelte-i18n';

    export let infos = [];
    export let error = null;
</script>

{#if error }
    <p class="error">{$_(error.message)}</p>
{/if}
{#each infos as info}
    <p class="info">{$_(info)}</p>
{/each}