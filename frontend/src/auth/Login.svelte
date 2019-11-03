<script>
    import { onMount } from 'svelte';
    import { _ } from 'svelte-i18n';
    import { navigate } from 'svelte-routing';

    import Info from '../Info.svelte';
    import { setToken } from './user';

    let credentials = {
        username: null,
        password: null
    }

    let validation = {
        username: null,
        password: null
    }

    function setUsername(username) {
        credentials.username = username;
        validateUsername();
    }

    function validateUsername() {
        validation.username = !!credentials.username && credentials.username.length > 0;
    }
    
    function setPassword(password) {
        credentials.password = password;
        validatePassword()
    }

    function validatePassword() {
        validation.password = !!credentials.password && credentials.password.length > 0;
    }

    function validateForm() {
        validateUsername();
        validatePassword();

        return validation.username && validation.password;
    }

    async function login() {
        if(!validateForm()) {
            return;
        }

        const response = await fetch('http://localhost:8080/login', {
            headers: {
                "Authorization": "Basic " + btoa(`${credentials.username}:${credentials.password}`)
            },
            credentials: 'include',
            method: 'POST',
        });

        let body = await response.json();
        setToken(body.token)

        console.log(response);
        navigate('/fleet');
    }

    onMount(async () => {
        const response = await fetch('http://localhost:8080/login', {
            method: 'POST',
            credentials: 'include',
        });
        if(response.ok) {
            let body = await response.json();
            setToken(body.token);
            console.log("Silent login successful")
            navigate('/fleet')
        }
    });
</script>

<h2>{$_('auth.login')}</h2>
<Info infos={[]} error={null} />
<form class="form" on:submit|preventDefault={login}>
    <input class:invalid={validation.username === false} type="text" name="username" placeholder={$_('auth.username')} on:input={(e) => setUsername(e.target.value)} />
    <input class:invalid={validation.password == false} type="password" name="password" placeholder={$_('auth.password')} on:input={(e) => setPassword(e.target.value)} />
    <input type="submit" value={$_('auth.login')} />
</form>