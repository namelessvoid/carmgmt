<script>
    import { _ } from 'svelte-i18n';

    import Info from '../Info.svelte';

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

        console.log(credentials.username, credentials.password);
    }
</script>

<h2>{$_('auth.login')}</h2>
<Info infos={[]} error={null} />
<form class="form" on:submit|preventDefault={login}>
    <input class:invalid={validation.username === false} type="text" name="username" placeholder={$_('auth.username')} on:input={(e) => setUsername(e.target.value)} />
    <input class:invalid={validation.password == false} type="password" name="password" placeholder={$_('auth.password')} on:input={(e) => setPassword(e.target.value)} />
    <input type="submit" value={$_('auth.login')} />
</form>