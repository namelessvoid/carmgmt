<script>
	import { Router, Route } from 'svelte-routing';
	import Fleet from './fleet/Fleet.svelte';
	import VehicleDetail from './fleet/VehicleDetail.svelte';
	import LanguageSwitcher from './i18n/LanguageSwitcher.svelte';

	let url = "";
</script>

<style>
	h1 {
		color: purple;
	}
</style>

<Router url={url}>
	<h1>Hello foo!</h1><LanguageSwitcher />

	<Route path="fleet" component={Fleet} />
	<Route path="fleet/vehicle/:vehicleId" component={VehicleDetail} />
</Router>
