<script>
    import { _ } from 'svelte-i18n';

    import { addVehicle } from './api';

    let errors = [];

    async function handleSubmit(event) {
        let vehicle = {}
        const formData = new FormData(event.target);
        formData.forEach((value, key) => {
            vehicle[key] = value;
        });

        try {
            vehicle = await addVehicle(vehicle);
            errors = []
        } catch(err) {
            errors = err;
        }
    }
</script>

<style>
label {
    display: inline;
}
</style>

<form on:submit|preventDefault={handleSubmit}>
    <label for="name">{$_('fleet.vehicle.name')}:</label>
    <input name="name" type="text" placeholder={$_('fleet.vehicle.name')} />
    <input type="submit" value={$_('fleet.addVehicle')} />
</form>

{#if errors}
    {#each errors as error}
        <p>{$_(error)}</p>
    {/each}
{/if}