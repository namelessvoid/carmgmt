<script>
    import { _, locale } from 'svelte-i18n';
    import { onMount } from 'svelte';

    import { addRefuellingToVehicle, formToJson } from './api';

    import Info from '../Info.svelte';
    import DatePicker from '../DatePicker.svelte';

    export let vehicleId;

    let errors = [];

    let dateInput;
    let fp;

    async function handleSubmit(event) {
        console.log("Adding refuelling to ", vehicleId);
        const refuelling = formToJson(event.target);
        refuelling.kilometers = parseFloat(refuelling.kilometers);
        refuelling.amount = parseFloat(refuelling.amount);
        refuelling.price = parseFloat(refuelling.price);
        refuelling.pricePerLiter = parseFloat(refuelling.pricePerLiter);

        console.log(refuelling);
        try {
            await addRefuellingToVehicle(vehicleId, refuelling);
        } catch(e) {
            errors = e;
        }
    }
</script>

<style>
form {
    max-width: 412px;
}

.form-row {
    display: flex;
    white-space: nowrap;
}

.form-row > label {
    padding: 0 8px 0 4px;
    flex: 1;
    align-items: center;
    display: flex;
}

.form-row > input {
    flex: 2;
    min-width: 0;
}
</style>

<form on:submit|preventDefault={handleSubmit}>
    <div class="form-row">
        <label for="time">{$_('fleet.refuelling.time')}</label>
        <DatePicker />
    </div>

    <div class="form-row">
        <label for="kilometers">{$_('fleet.refuelling.kilometers')}</label>
        <input type="number" step="any" name="kilometers" />
    </div>

    <div class="form-row">
        <label for="amount">{$_('fleet.refuelling.amount')}</label>
        <input type="number" step="any" name="amount" />
    </div>
    
    <div class="form-row">
        <label for="price">{$_('fleet.refuelling.price')}</label>
        <input type="number" step="any" name="price" />
    </div>

    <div class="form-row">
        <label for="pricePerLiter">{$_('fleet.refuelling.pricePerLiter')}</label>
        <input type="number" step="any" name="pricePerLiter" />
    </div>

    <div class="form-row">
        <input type="submit" value={$_('fleet.addRefuelling')} />
    </div>
</form>

<Info infos={errors} />
