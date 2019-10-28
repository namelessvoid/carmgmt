import { cleanup, render } from '@testing-library/svelte';

import './i18n/dictionary';
import './i18n/locale';

import Info from './Info.svelte';

describe('Info', () => {
    afterAll(cleanup);

    it('renders empty if no info and error is provided', () => {
        const { container } = render(Info, { props: { infos: [], error: null }});
        expect(container.innerHTML).toBe('<div> </div>');
    });

    it('renders provided info', () => {
        render(Info, { props: { infos: [ 'foo', 'bar' ]}});
        expect(document.querySelectorAll('.info').length).toBe(2);
    });

    it('renders provided error', () => {
        render(Info, { props: { infos: [], error: new Error('fooError') }});
        expect(document.querySelectorAll('.error').length).toBe(1);
    });
});