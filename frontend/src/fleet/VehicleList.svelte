<script>
    import { onMount } from 'svelte';
    import { _ } from 'svelte-i18n';
    import { navigate } from 'svelte-routing';

    import Info from '../Info.svelte';
    import LoadingSpinner from '../LoadingSpinner.svelte';

    import { getAllVehicles } from './api';

    export const update = async () => {
        loading = true;

        try {
            vehicles = await getAllVehicles();
        } catch(e) {
            vehicles = [];
            errors = e;
        }

        loading = false;
    }

    let vehicles = [];
    let loading = true;
    let errors = [];

    onMount(async () => {
        await update();
    });
</script>

<style>
td {
    cursor: pointer;
}
</style>

<table class="table">
    <tr>
        <th>{$_('fleet.vehicle.id')}</th>
        <th>{$_('fleet.vehicle.name')}</th>
    </tr>
    {#each vehicles as vehicle (vehicle.id)}
    <tr on:click={() => navigate(`/fleet/vehicle/${vehicle.id}`)}>
        <td>{vehicle.id}</td>
        <td>{vehicle.name}</td>
    </tr>
    {/each}
</table>

<LoadingSpinner loading={loading} />
<Info infos={errors} />
