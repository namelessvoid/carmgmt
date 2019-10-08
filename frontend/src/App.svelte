<script>
	import { Router, Route } from 'svelte-routing';
	import CarList from './fleet/CarList.svelte';

	let url = "";
	let loading = false;
</script>

<style>
	h1 {
		color: purple;
	}
</style>

<Router url={url}>
	<h1>Hello foo!</h1>

	<Route path="fleet" component={CarList} />
</Router>
