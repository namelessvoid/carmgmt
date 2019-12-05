import { writable } from 'svelte/store';
import { navigate } from 'svelte-routing';

const user = {
    token: null,
};

let auth0 = null;
export const isAuthenticated = writable(false);

export async function init() {
    if(!auth0) {
        auth0 = await createAuth0Client({
            domain: "dev-fleetmgmt.eu.auth0.com",
            client_id: "GAt51deyHixXHSIEA0DAmxQHJj3tcYxa",
            audience: "Fleet Management - Local"
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
        redirect_uri: "http://localhost:5000/"
    });
}

export async function logout() {
    auth0.logout({
        returnTo: "http://localhost:5000/"
    });
    alert("You've been logged out");
}

export function setToken(token) {
   user.token = token;
}

export async function getToken() {
    return await auth0.getTokenSilently();
}