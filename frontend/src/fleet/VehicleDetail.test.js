import { cleanup, render } from '@testing-library/svelte';

import '../i18n/dictionary';
import '../i18n/locale';

import { getVehicleDetail } from './api';
jest.mock('./api');

import VehicleDetail from './VehicleDetail.svelte';

describe('VehicleDetail', () => {
    afterEach(cleanup);

    it('renders error if api returns error', async () => {
        getVehicleDetail.mockImplementation(() => { throw new Error('some.error') });

        render(VehicleDetail);

        expect(document.querySelectorAll('.error').length).toBe(1);
    });
});