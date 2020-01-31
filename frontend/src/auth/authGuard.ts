import { getInstance } from './index';
import { NavigationGuard } from 'vue-router/types';

export const authGuard: NavigationGuard = (to, from, next) => {
    const authService = getInstance();

    const fn = () => {
        if (authService.isAuthenticated) {
            return next();
        }

        authService.loginWithRedirect({ appState: { targetUrl: to.fullPath } });
    }

    if(!authService.loading) {
        return fn();
    }

    authService.$watch("loading", (loading: Boolean) => {
        if (loading === false) {
          return fn();
        }
    });
};