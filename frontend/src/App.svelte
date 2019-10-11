<script>
	import { Router, Route } from 'svelte-routing';
	import Fleet from './fleet/Fleet.svelte';
	import VehicleDetail from './fleet/VehicleDetail.svelte';
	import Header from './Header.svelte';

	let url = "";
</script>

<style>
.content {
	padding: 8px;
}
</style>

<Header />
<div class="content">
	<Router url={url}>
		<Route path="fleet" component={Fleet} />
		<Route path="fleet/vehicle/:vehicleId" component={VehicleDetail} />
	</Router>
</div>