<script>
    import { _ } from 'svelte-i18n';

    import VehicleList from './VehicleList.svelte';
    import AddVehicle from './AddVehicle.svelte';

    let updateVehicleList;
</script>

<h2>{$_('fleet.overview')}</h2>
<VehicleList bind:update={updateVehicleList}/>

<h4>{$_('fleet.addVehicle')}</h4>
<AddVehicle on:vehicleAdded={() => updateVehicleList()}/>