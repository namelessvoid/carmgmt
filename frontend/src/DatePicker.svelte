<script>
    import { createEventDispatcher } from 'svelte';

    export let value = new Date();
    value.setSeconds(0);
    value.setMilliseconds(0);

    let validation = {
        date: true,
        time: true
    };

    const dispatch = createEventDispatcher();

    function setDate(dateString) {
        const date = new Date(dateString);
        validation.date = date.toString() !== 'Invalid Date';

        value.setFullYear(date.getFullYear());
        value.setMonth(date.getMonth());
        value.setDate(date.getDate());

        dispatch('changed', value);
    }

    function setTime(timeString) {
        validation.time = timeString !== "";

        const [hours, minutes] = timeString.split(':');
        value.setHours(parseInt(hours));
        value.setMinutes(parseInt(minutes));

        dispatch('changed', value);
    }

    function getDateString(date) {
        const month = `${date.getMonth() + 1}`.padStart(2, '0');
        const day = `${date.getDate()}`.padStart(2, '0');
        return `${date.getFullYear()}-${month}-${day}`;
    }

    function getTimeString(date) {
        const hours = `${date.getHours()}`.padStart(2, '0');
        const minutes = `${date.getMinutes()}`.padStart(2, '0');
        return `${hours}:${minutes}`;
    }
</script>

<style>
input {
    flex: 2;
    min-width: 0;
}
</style>

<input name="date" type="date" value={getDateString(value)} on:change={(e) => setDate(e.target.value)} class:invalid={validation.date === false} />
<input name="time" type="time" value={getTimeString(value)} on:change={(e) => setTime(e.target.value)} class:invalid={validation.time === false} />