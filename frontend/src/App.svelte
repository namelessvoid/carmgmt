<script>
	import CarList from './CarList.svelte';

	export let name;
</script>

<style>
	h1 {
		color: purple;
	}
</style>

<h1>Hello {name}!</h1>
<CarList />
