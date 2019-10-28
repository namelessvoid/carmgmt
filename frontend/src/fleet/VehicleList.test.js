import { cleanup, render } from '@testing-library/svelte';

import '../i18n/dictionary';
import '../i18n/locale';

import { getAllVehicles } from './api';
jest.mock('./api');

import VehicleList from './VehicleList.svelte';

describe('VehicleList', () => {
    afterEach(cleanup);

    it('renders error if api returns error', async () => {
        getAllVehicles.mockImplementation(() => { throw new Error('some.error') });

        render(VehicleList);

        expect(document.querySelectorAll('.error').length).toBe(1);
    });
});