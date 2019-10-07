<script>
    import { onMount } from 'svelte';

    let cars = []

    onMount(() => {
        fetch("http://localhost:8080/cars")
        .then(response => response.json())
        .then(json => cars = json);
    });
</script>
<h2>Cars</h2>
{#each cars as car}
<div>{car.ID}: {car.Name}</div>
{/each}