<script>
    import { _ } from 'svelte-i18n';

    import VehicleList from './VehicleList.svelte';
    import AddVehicle from './AddVehicle.svelte';
</script>

<h2>{$_('fleet.overview')}</h2>
<VehicleList />
<AddVehicle />