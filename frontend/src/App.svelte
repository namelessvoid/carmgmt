<script>
	import { Router, Route, Link } from 'svelte-routing';
	import FleetPage from './fleet/FleetPage.svelte';
	import VehiclePage from './fleet/VehiclePage.svelte';
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
	<Router url={url} basepath="/fleet">
		<Route path="/" component={FleetPage} />
		<Route path="/vehicle/:vehicleId" component={VehiclePage} />
	</Router>
</div>