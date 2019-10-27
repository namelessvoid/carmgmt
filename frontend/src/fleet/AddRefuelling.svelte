<script>
    import { _, locale } from 'svelte-i18n';
    import { onMount } from 'svelte';

    import { addRefuellingToVehicle } from './api';

    import Info from '../Info.svelte';
    import DatePicker from '../DatePicker.svelte';

    export let vehicleId;

    let refuelling = {
        time: new Date(),
        tripKilometers: null,
        amount: null,
        price: null,
        pricePerLiter: null
    };

    let validation = {
        time: null,
        tripKilometers: null,
        amount: null,
        price: null,
        pricePerLiter: null
    };

    let errors = [];

    async function handleSubmit(event) {
        if(!validateForm()) {
            return;
        }
    
        const json = {
            ...refuelling,
            time: refuelling.time.toISOString()
        };

        try {
            await addRefuellingToVehicle(vehicleId, json);
        } catch(e) {
            errors = e;
        }
    }

    function parseStringToFloat(s) {
        return parseFloat(s.replace(',', '.'));
    }

    function validateForm() {
        validateTime();
        validateTripKilometers();
        validateAmount();
        validatePrice();
        validatePricePerLiter();

        return validation.time
            && validation.tripKilometers
            && validation.amount
            && validation.price
            && validation.pricePerLiter;
    }

    function validateTime() {
        validation.time = refuelling.time.toString() != 'Invalid Date';
    }

    function validateTripKilometers() {
        validation.tripKilometers = refuelling.tripKilometers != null && !isNaN(refuelling.tripKilometers);
    }

    function validateAmount() {
        validation.amount = refuelling.amount != null && !isNaN(refuelling.amount);
    }

    function validatePrice() {
        validation.price = refuelling.price != null && !isNaN(refuelling.price);
    }

    function validatePricePerLiter() {
        validation.pricePerLiter = refuelling.pricePerLiter !== null && !isNaN(refuelling.pricePerLiter);
    }

    function setTime(date) {
        refuelling.time = date;
        validateTime();
    }

    function setTripKilometers(tripKilometers) {
        refuelling.tripKilometers = parseStringToFloat(tripKilometers);
        validateTripKilometers();
    }

    function setAmount(amount) {
        refuelling.amount = parseStringToFloat(amount);
        validateAmount();
    }

    function setPrice(price) {
        refuelling.price = parseStringToFloat(price);
        validatePrice();
    }

    function setPricePerLiter(pricePerLiter) {
        refuelling.pricePerLiter = parseStringToFloat(pricePerLiter);
        validatePricePerLiter();
    }
</script>

<form class="form" on:submit|preventDefault={handleSubmit} novalidate>
    <div class="form-row">
        <label>{$_('fleet.refuelling.time')}</label>
        <DatePicker value={refuelling.time} on:changed={(e) => setTime(e.detail)} />
    </div>

    <div class="form-row">
        <label for="tripKilometers">{$_('fleet.refuelling.tripKilometers')}</label>
        <input type="number" step="any" name="tripKilometers" on:input={(e) => setTripKilometers(e.target.value)} class:invalid={validation.tripKilometers === false}/>
    </div>

    <div class="form-row">
        <label for="amount">{$_('fleet.refuelling.amount')}</label>
        <input type="number" step="any" name="amount" on:input={(e) => setAmount(e.target.value)} class:invalid={validation.amount === false} />
    </div>
    
    <div class="form-row">
        <label for="price">{$_('fleet.refuelling.price')}</label>
        <input type="number" step="any" name="price" on:input={(e) => setPrice(e.target.value)} class:invalid={validation.price === false} />
    </div>

    <div class="form-row">
        <label for="pricePerLiter">{$_('fleet.refuelling.pricePerLiter')}</label>
        <input type="number" step="any" name="pricePerLiter" on:input={(e) => setPricePerLiter(e.target.value)} class:invalid={validation.pricePerLiter === false} />
    </div>

    <div class="form-row">
        <input type="submit" value={$_('fleet.addRefuelling')} />
    </div>
</form>

<Info infos={errors} />
