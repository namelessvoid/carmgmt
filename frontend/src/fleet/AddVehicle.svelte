<script>
    import { _ } from 'svelte-i18n';
    import { createEventDispatcher } from 'svelte';

    import Info from '../Info.svelte';

    import { addVehicle } from './api';

    let infos = [];
    let error = null;

    let vehicle = null;
    let validation = null;
    let form = null;

    const dispatchVehicleAdded = createEventDispatcher();

    function resetForm() {
        vehicle = {name: ""};
        validation = {name: ""};

        form && form.reset();
    }
    resetForm();

    function validateForm() {
        return validateVehicleName();
    }

    function validateVehicleName() {
        validation.name = vehicle.name.length > 0;
        return validation.name;
    }

    function setVehicleName(name) {
        vehicle.name = name;
        validateVehicleName();
    }

    async function handleSubmit() {
        if(!validateForm()) {
            return;
        }

        try {
            await addVehicle(vehicle);
            infos = ['fleet.vehicleAdded']
            dispatchVehicleAdded('vehicleAdded');
            resetForm();
        } catch(e) {
            error = e;
        }
    }
</script>

<form class="form" on:submit|preventDefault={handleSubmit} bind:this={form}>
    <div class="form-row">
        <label for="name">{$_('fleet.vehicle.name')}:</label>
        <input name="name" class:invalid={validation.name === false} type="text" placeholder={$_('fleet.vehicle.name')} on:input={(e) => setVehicleName(e.target.value)}/>
    </div>

    <div class="form-row">
        <input type="submit" value={$_('fleet.addVehicle')} />
    </div>
</form>

<Info infos={infos} error={error} />