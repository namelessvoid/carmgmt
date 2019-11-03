import { cleanup, render, fireEvent } from '@testing-library/svelte';

import '../i18n/dictionary';
import '../i18n/locale';

import Login from './Login.svelte';

describe('Login', () => {
    afterEach(cleanup);

    function usernameInput() { return document.querySelector('input[name=username]'); }
    function passwordInput() { return document.querySelector('input[name=password]'); }
    function submitButton() { return document.querySelector('input[type=submit]'); }

    it('should not display fields as invalid initially', () => {
        render(Login);

        expect(usernameInput().classList).not.toContain('invalid');
        expect(passwordInput().classList).not.toContain('invalid');
    });

    it('should validate on form submit', async () => {
        render(Login);

        await fireEvent.click(submitButton());

        expect(usernameInput().classList).toContain('invalid');
        expect(passwordInput().classList).toContain('invalid');
    });

    it('should validate on input', async () => {
        render(Login);

        await fireEvent.input(usernameInput(), { target: { value: '' }});
        expect(usernameInput().classList).toContain('invalid');
        expect(passwordInput().classList).not.toContain('invalid');

        await fireEvent.input(passwordInput(), { target: { value: '' }});
        expect(usernameInput().classList).toContain('invalid');
        expect(passwordInput().classList).toContain('invalid');

        await fireEvent.input(usernameInput(), { target: { value: 'someusername' }});
        expect(usernameInput().classList).not.toContain('invalid');
        expect(passwordInput().classList).toContain('invalid');

        await fireEvent.input(passwordInput(), { target: { value: 'somepassword' }});
        expect(usernameInput().classList).not.toContain('invalid');
        expect(passwordInput().classList).not.toContain('invalid');
    });
});