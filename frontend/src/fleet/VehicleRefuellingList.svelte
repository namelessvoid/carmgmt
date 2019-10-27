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

{#if refuellings.length > 0 }
<table class="table">
    <tr>
        <th>{$_('fleet.refuelling.time')}</th>
        <th>{$_('fleet.refuelling.tripKilometers')}</th>
        <th>{$_('fleet.refuelling.amount')}</th>
        <th>{$_('fleet.refuelling.price')}</th>
        <th>{$_('fleet.refuelling.pricePerLiter')}</th>
        <th>{$_('fleet.refuelling.consumption')}</th>
    </tr>
    {#each refuellings as refuelling}
    <tr>
        <td>{$_.date(Date.parse(refuelling.time), 'short')}</td>
        <td>{$_.number(refuelling.tripKilometers)} km</td>
        <td>{$_.number(refuelling.amount)} L</td>
        <td>{$_.number(refuelling.price)} €</td>
        <td>{$_.number(refuelling.pricePerLiter)} €/L</td>
        <td>{$_.number(refuelling.consumption)} L/100km</td>
    </tr>
    {/each}
</table>
{:else}
<p>{$_('fleet.refuelling.emptyList')}</p>
{/if}

<LoadingSpinner loading={loading} />
<Info infos={errors} />
