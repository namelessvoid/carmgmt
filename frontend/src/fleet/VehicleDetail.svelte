<script>
    import { _ } from 'svelte-i18n';
    import { Link } from 'svelte-routing';
    import { onMount } from 'svelte';

    import Info from '../Info.svelte';
    import VehicleRefuellingList from './VehicleRefuellingList.svelte';

    import { getVehicleDetail } from './api';

    export let vehicleId = null;
    export let location = null;

    let vehicle = null;
    let errors = [];
    let loading = false;

    onMount(async () => {
        loading = true;

        try {
            vehicle = await getVehicleDetail(vehicleId);
        } catch(e) {
            errors = e;
        }

        loading = false;
    });
</script>

<Link to="/fleet">{$_('fleet.backToOverview')}</Link>

{#if loading}
<span>{$_('loading')}</span>
{/if}

<Info infos={errors} />

{#if !loading && vehicle}
<h2>{vehicle.name}</h2>
{/if}

<h3>{$_('fleet.vehicle.refuellings')}</h3>
<VehicleRefuellingList vehicleId={vehicleId} />