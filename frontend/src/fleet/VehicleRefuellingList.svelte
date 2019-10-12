<script>
    import { onMount } from 'svelte';
    import { _ } from 'svelte-i18n';

    import { getRefuellingsByVehicle } from './api';

    import Info from '../Info.svelte';
    import LoadingSpinner from '../LoadingSpinner.svelte';

    export let vehicleId = null;

    let loading = false;
    let errors = [];
    let refuellings = [];

    onMount(async () => {
        loading = true;

        try {
            refuellings = await getRefuellingsByVehicle(vehicleId);
        } catch(e) {
            errors = e;
        }
        
        loading = false;
    });
</script>

<table>
    <tr>
        <th>{$_('fleet.refuelling.time')}</th>
        <th>{$_('fleet.refuelling.kilometers')}</th>
        <th>{$_('fleet.refuelling.amount')}</th>
        <th>{$_('fleet.refuelling.price')}</th>
        <th>{$_('fleet.refuelling.pricePerLiter')}</th>
        <th>{$_('fleet.refuelling.consumption')}</th>
    </tr>
    {#each refuellings as refuelling}
    <tr>
        <td>{refuelling.time}</td>
        <td>{refuelling.kilometers}</td>
        <td>{refuelling.amount}</td>
        <td>{refuelling.price}</td>
        <td>{refuelling.pricePerLiter}</td>
        <td>tdb</td>
    {/each}
</table>

<LoadingSpinner loading={loading} />
<Info infos={errors} />
