import { writable } from 'svelte/store';

const user = {
    token: null
};

export function setToken(token) {
   user.token = token;
}

export function getToken() {
    return user.token;
}