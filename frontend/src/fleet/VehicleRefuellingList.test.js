import { cleanup, render } from '@testing-library/svelte';

import '../i18n/dictionary';
import '../i18n/locale';

import { getRefuellingsByVehicle } from './api';
jest.mock('./api');

import VehicleRefuellingList from './VehicleRefuellingList.svelte';

describe('VehicleRefuellingList', () => {
    afterEach(cleanup);

    it('renders error if api returns error', async () => {
        getRefuellingsByVehicle.mockImplementation(() => { throw new Error('some.error') });

        render(VehicleRefuellingList);

        expect(document.querySelectorAll('.error').length).toBe(1);
    });
});