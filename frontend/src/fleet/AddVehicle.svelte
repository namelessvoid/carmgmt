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
            await addVehicle(vehicle);
            infos = ['fleet.vehicleAdded']
            dispatchVehicleAdded('vehicleAdded');
        } catch(err) {
            infos = err;
        }
    }
</script>

<form class="form" on:submit|preventDefault={handleSubmit}>
    <div class="form-row">
        <label for="name">{$_('fleet.vehicle.name')}:</label>
        <input name="name" type="text" placeholder={$_('fleet.vehicle.name')} required on:invalid={(e) => e.target.setCustomValidity($_('fleet.vehicle.validation.nameNotEmpty'))} on:input={(e) => e.target.setCustomValidity('')}/>
    </div>

    <div class="form-row">
        <input type="submit" value={$_('fleet.addVehicle')} />
    </div>
</form>

<Info infos={infos} />