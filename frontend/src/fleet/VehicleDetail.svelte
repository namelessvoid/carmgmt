<script>
    import { _ } from 'svelte-i18n';
    import { Link } from 'svelte-routing';

    export let vehicleId = null;
</script>

<Link to="/fleet">{$_('fleet.backToOverview')}</Link>
<h2>Vehicle Details for {vehicleId}</h2>