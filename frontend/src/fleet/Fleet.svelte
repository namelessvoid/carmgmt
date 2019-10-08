<script>
    import { _ } from 'svelte-i18n';

    import CarList from './CarList.svelte';
</script>

<h2>{$_('fleet.overview')}</h2>
<CarList />