<script>
    import { _ } from 'svelte-i18n';
    import { Link } from 'svelte-routing';
    import { onMount } from 'svelte';

    import Info from '../Info.svelte';
    import LoadingSpinner from '../LoadingSpinner.svelte';

    import { getVehicleDetail } from './api';

    export let vehicleId = null;

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

<LoadingSpinner loading={loading} />
<Info infos={errors} />

{#if vehicle}
<h2>{vehicle.name}</h2>
{/if}
