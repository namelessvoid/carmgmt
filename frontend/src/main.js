import './i18n/dictionary';
import './i18n/locale';

import App from './App.svelte';
import { init } from './auth/auth';

const app = init().then(() => {
	return new App({
		target: document.body,
	});
})


export default app;