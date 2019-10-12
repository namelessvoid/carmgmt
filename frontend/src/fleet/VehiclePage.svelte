<script>
    import { Link } from 'svelte-routing';
    import { _ } from 'svelte-i18n';

    export let vehicleId;
    
    import VehicleDetail from './VehicleDetail.svelte';
    import VehicleRefuellingList from './VehicleRefuellingList.svelte';
</script>

<Link to="/fleet">{$_('fleet.backToOverview')}</Link>

<VehicleDetail vehicleId={vehicleId} />

<h3>{$_('fleet.vehicle.refuellings')}</h3>
<VehicleRefuellingList vehicleId={vehicleId} />