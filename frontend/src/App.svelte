<script>
	import { Router, Route, Link } from 'svelte-routing';
	import Login from './auth/Login.svelte';
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
	<Router url={url}>
		<Route path="/" component={Login} />
		<Route path="/fleet" component={FleetPage} />
		<Route path="/fleet/vehicle/:vehicleId" component={VehiclePage} />
	</Router>
</div>