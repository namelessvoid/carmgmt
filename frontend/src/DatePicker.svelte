<script>
    import { locale } from 'svelte-i18n';
    import { onMount } from 'svelte';

    import Flatpickr from 'flatpickr';
    import 'flatpickr/dist/flatpickr.css';
    import FlatpickrLanguages from "flatpickr/dist/l10n";

    let fp;
    let dateInput;

    function reCreateFlatpickr(loc) {
        let fpLocale = loc === 'de-DE' ? 'de' : 'en';
        let dateStr = "";

        if(fp) {
            dateStr = fp.input.value;
            fp.destroy();
        }

        fp = Flatpickr(dateInput, {
            enableTime: true,
            locale: FlatpickrLanguages[fpLocale]
        });
        fp.setDate(dateStr);
    }
    
    onMount(() => {
        fp = reCreateFlatpickr($locale);

        const unsubscribe = locale.subscribe(l => {
            reCreateFlatpickr(l);
        });

        return () => {
            fp.destroy();
            unsubscribe();
            console.log("DatePicker destroyed");
        };
    });
</script>

<input name="time" bind:this={dateInput} />