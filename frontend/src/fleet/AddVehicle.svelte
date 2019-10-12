<script>
    import { _ } from 'svelte-i18n';
    import { createEventDispatcher } from 'svelte';

    import Info from '../Info.svelte';

    import { addVehicle, formToJson } from './api';

    let infos = [];

    const dispatchVehicleAdded = createEventDispatcher();

    async function handleSubmit(event) {
        const vehicle = formToJson(event.target);

        try {
            vehicle = await addVehicle(vehicle);
            infos = ['fleet.vehicleAdded']
            dispatchVehicleAdded('vehicleAdded');
        } catch(err) {
            infos = err;
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

<Info infos={infos} />