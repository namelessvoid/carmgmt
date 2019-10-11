<script>
    import { _ } from 'svelte-i18n';

    export let infos = [];
</script>

{#each infos as info}
    <p>{$_(info)}</p>
{/each}