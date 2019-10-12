<script>
    import { _ } from 'svelte-i18n';

    export let loading = false;
</script>

{#if loading}
<p>{$_('loading')}</p>
{/if}