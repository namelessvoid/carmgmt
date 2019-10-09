<script>
    import { onMount } from 'svelte';
    import { _ } from 'svelte-i18n';

    import { getAllVehicles } from './api';

    let vehicles = [];
    let loading = false;

    onMount(async () => {
        loading = true;
        vehicles = await getAllVehicles();
        loading = false;
    });
</script>

<style>
table {
    width: 100%;
}

th {
    text-align: left;
}
</style>

{#if loading}
    <p>{$_('loading')}</p>
{:else}
    <table>
        <tr>
            <th>{$_('fleet.vehicle.id')}</th>
            <th>{$_('fleet.vehicle.name')}</th>
        </tr>
        {#each vehicles as vehicle}
        <tr>
            <td>{vehicle.id}</td>
            <td>{vehicle.name}</td>
        </tr>
        {/each}
    </table>
{/if}

