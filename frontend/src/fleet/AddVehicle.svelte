<script>
    import { _ } from 'svelte-i18n';
    import { createEventDispatcher } from 'svelte';

    import Info from '../Info.svelte';

    import { addVehicle, formToJson } from './api';

    let infos = [];
    let vehicle = {
        name: null
    };

    let validation = {
        name: null
    }

    const dispatchVehicleAdded = createEventDispatcher();

    function vehicleIsValid() {
        return validateVehicleName();
    }

    function validateVehicleName() {
        validation.name = (vehicle.name !== null && vehicle.name.length > 0);
        return validation.name;
    }

    function setVehicleName(name) {
        vehicle.name = name;
        validateVehicleName();
    }

    async function handleSubmit(event) {
        if(!vehicleIsValid()) {
            return;
        }

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
        <input name="name" class:invalid={validation.name === false} type="text" placeholder={$_('fleet.vehicle.name')} on:input={(e) => setVehicleName(e.target.value)}/>
    </div>

    <div class="form-row">
        <input type="submit" value={$_('fleet.addVehicle')} />
    </div>
</form>

<Info infos={infos} />