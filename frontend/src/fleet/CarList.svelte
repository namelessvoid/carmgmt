<script>
    import { onMount } from 'svelte';
    import { _ } from 'svelte-i18n';

    let cars = [];
    let loading = false;

    onMount(() => {
        loading = true;
        fetch("http://localhost:8080/cars")
        .then(response => response.json())
        .then(json => {
            cars = json;
            loading = false;
        });
    });
</script>

<style>
table {
    width: 100%;
}

th {
    text-align: left;
}
</style>

{#if loading}
    <p>{$_('loading')}</p>
{:else}
    <table>
        <tr>
            <th>{$_('fleet.car.id')}</th>
            <th>{$_('fleet.car.name')}</th>
        </tr>
        {#each cars as car}
        <tr>
            <td>{car.ID}</td>
            <td>{car.Name}</td>
        </tr>
        {/each}
    </table>
{/if}

