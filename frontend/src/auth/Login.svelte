<script>
    import { isAuthenticated, login } from './auth.js';
    
    import { onMount } from 'svelte';
    import { navigate } from 'svelte-routing';

    onMount(async () => {
        await login();

        if($isAuthenticated === true) {
            navigate("/fleet");
        }
    });
</script>

<h1>Login page</h1>
