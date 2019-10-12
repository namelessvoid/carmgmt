<script>
    import { _, locale } from 'svelte-i18n';
    import { onMount } from 'svelte';

    import DatePicker from '../DatePicker.svelte';

    export let vehicleId;

    let dateInput;
    let fp;

    async function handleSubmit() {
        console.log("Adding refuelling to ", vehicleId);
    }
</script>

<style>
.form-row {
    display: flex;
    white-space: nowrap;
}

.form-row > label {
    padding: 0 8px 0 4px;
    flex: 1;
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
        <input type="number" name="kilometers" />
    </div>

    <div class="form-row">
        <label for="amount">{$_('fleet.refuelling.amount')}</label>
        <input type="number" name="amount" />
    </div>
    
    <div class="form-row">
        <label for="price">{$_('fleet.refuelling.price')}</label>
        <input type="number" name="price" />
    </div>

    <div class="form-row">
        <label for="pricePerLiter">{$_('fleet.refuelling.pricePerLiter')}</label>
        <input type="number" name="pricePerLiter" />
    </div>    
</form>
