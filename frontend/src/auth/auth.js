import { writable } from 'svelte/store';
import config from './config'

const user = {
    token: null,
};

let auth0 = null;
export const isAuthenticated = writable(false);

export async function init() {
    if(!auth0) {
        auth0 = await createAuth0Client({
            domain: config.domain,
            client_id: config.clientId,
            audience: config.audience
        });
    }
}

export async function login() {
    const query = window.location.search;
    if (query.includes("code=") && query.includes("state=")) {
        await auth0.handleRedirectCallback();

        isAuthenticated.update(() => true);
        
        return;
    }

    await auth0.loginWithRedirect({
        redirect_uri: config.loginRedirectUrl
    });
}

export async function logout() {
    auth0.logout();
    alert("You've been logged out");
}

export function setToken(token) {
   user.token = token;
}

export async function getToken() {
    return await auth0.getTokenSilently();
}