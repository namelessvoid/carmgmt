import { cleanup, render, fireEvent } from '@testing-library/svelte';

import '../i18n/dictionary';
import '../i18n/locale';

import { addVehicle } from './api';

import AddVehicle from './AddVehicle.svelte';

jest.mock('./api');

describe('AddVehicle', () => {
    function nameInput() {
        return document.querySelector('input[name=name]');
    }

    function submitButton() {
        return document.querySelector('input[type=submit]');
    }

    beforeEach(() => {
        addVehicle.mockClear();
    });

    afterEach(cleanup);

    test('should add vehicle when valid form is submitted', async () => {
        render(AddVehicle);

        await fireEvent.input(nameInput(),  {target: {value: 'Some Name'}});
        await fireEvent.click(submitButton());

        expect(addVehicle).toHaveBeenCalledWith({'name': 'Some Name'})
    });

    test('should emit vehicleAdded event after adding vehicle successfully', async () => {
        const eventListener = jest.fn();

        const { component } = render(AddVehicle);

        component.$on('vehicleAdded', eventListener);

        await fireEvent.input(nameInput(),  {target: {value: 'Some Name'}});
        await fireEvent.click(submitButton());

        expect(eventListener).toHaveBeenCalled();        
    });

    test('should reset form after adding vehicle successfully', async () => {
        render(AddVehicle);

        await fireEvent.input(nameInput(), {target: {value: 'Some Name'}});
        await fireEvent.click(submitButton());

        expect(nameInput().value).toBe("");
    });

    test('should not add vehicle when vehicle name is empty', async () => {
        render(AddVehicle)

        await fireEvent.click(submitButton());

        expect(addVehicle).not.toHaveBeenCalled();
        expect(nameInput().classList).toContain('invalid');
    });

    test('should not display name as invalid initially', () => {
        render(AddVehicle);

        expect(nameInput().classList).not.toContain('invalid');
    });

    test('should not display name as invalid when name is valid', async () => {
        render(AddVehicle);
        
        await fireEvent.input(nameInput(),  {target: {value: 'Some Name'}});

        expect(nameInput().classList).not.toContain('invalid');
    });

    test('should highlight invalid name', async () => {
        render(AddVehicle);

        await fireEvent.input(nameInput(),  {target: {value: ''}});

        expect(nameInput().classList).toContain('invalid');
    });
});