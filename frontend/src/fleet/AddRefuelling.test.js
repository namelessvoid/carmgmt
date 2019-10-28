import { cleanup, render, fireEvent } from '@testing-library/svelte';

import '../i18n/dictionary';
import '../i18n/locale';

import { addRefuellingToVehicle } from './api';

import AddRefuelling from './AddRefuelling.svelte';

jest.mock('./api')

describe('AddRefuelling', () => {
    beforeEach(() => {
        addRefuellingToVehicle.mockClear();
    });

    afterEach(cleanup);

    const dateInput = () => document.querySelector('input[name=date]');
    const timeInput = () => document.querySelector('input[name=time]');
    const tripKilometersInput = () => document.querySelector('input[name=tripKilometers]');
    const amountInput = () => document.querySelector('input[name=amount]');
    const priceInput = () => document.querySelector('input[name=price]');
    const pricePerLiterInput = () => document.querySelector('input[name=pricePerLiter]');
    const submitButton = () => document.querySelector('input[type=submit]');

    const fillInForm = async (formData) => {
        await fireEvent.change(dateInput(), {target: {value: formData.date}});
        await fireEvent.change(timeInput(), {target: {value: formData.time}});

        if(formData.tripKilometers !== undefined)
            await fireEvent.input(tripKilometersInput(), {target: {value: formData.tripKilometers}});

        if(formData.amount !== undefined)
            await fireEvent.input(amountInput(), {target: {value: formData.amount}});
        
        if(formData.price !== undefined)
            await fireEvent.input(priceInput(), {target: {value: formData.price}});

        if(formData.pricePerLiter !== undefined)
            await fireEvent.input(pricePerLiterInput(), {target: {value: formData.pricePerLiter}});
    }

    const validFormData = {
        date: '2010-12-24',
        time: '03:04',
        tripKilometers: 687.4,
        amount: 35.6,
        price: 44.46,
        pricePerLiter: 1.249
    };

    const validRefuelling = {
        time: (new Date('2010-12-24T03:04:00')).toISOString(),
        tripKilometers: validFormData.tripKilometers,
        amount: validFormData.amount,
        price: validFormData.price,
        pricePerLiter: validFormData.pricePerLiter
    }

    it('should submit valid refuelling', async () => {
        render(AddRefuelling, { props: { vehicleId: 20 }});

        await fillInForm(validFormData);
        await fireEvent.click(submitButton());

        expect(addRefuellingToVehicle).toHaveBeenCalledWith(20, {...validRefuelling});
    });

    it.each`
        formData                                  | input                 | reason
        ${{...validFormData, date: ''}}           | ${dateInput}          | ${'date is empty'}
        ${{...validFormData, time: ''}}           | ${timeInput}          | ${'time is empty'}
        ${{...validFormData, price: ''}}          | ${priceInput}         | ${'price is empty'}
        ${{...validFormData, price: 0}}           | ${priceInput}         | ${'price is zero'}
        ${{...validFormData, price: -1}}          | ${priceInput}         | ${'price is negative'}
        ${{...validFormData, amount: ''}}         | ${amountInput}        | ${'amount is empty'}
        ${{...validFormData, amount: 0}}          | ${amountInput}        | ${'amount is zero'}
        ${{...validFormData, amount: -1}}         | ${amountInput}        | ${'amount is negative'}
        ${{...validFormData, pricePerLiter: ''}}  | ${pricePerLiterInput} | ${'pricePerLiter is empty'}        
        ${{...validFormData, pricePerLiter: 0}}   | ${pricePerLiterInput} | ${'pricePerLiter is zero'}
        ${{...validFormData, pricePerLiter: -1}}  | ${pricePerLiterInput} | ${'pricePerLiter is negative'}  
        ${{...validFormData, tripKilometers: ''}} | ${tripKilometersInput} | ${'tripKilometers is empty'}
        ${{...validFormData, tripKilometers: 0}}  | ${tripKilometersInput} | ${'tripKilometers is zero'}
        ${{...validFormData, tripKilometers: -1}} | ${tripKilometersInput} | ${'tripKilometers is negative'}
    `('should have invalid form when $reason', async ({ formData, input }) => {
        render(AddRefuelling, { props: { vehicleId: 20 }});

        await fillInForm(formData);

        // Input field is marked invalid
        expect(input().classList).toContain('invalid');

        // Form does not submit
        await fireEvent.click(submitButton());
        expect(addRefuellingToVehicle).not.toHaveBeenCalled();
    });

    it.each`
        formData | input | inputName
        ${{...validFormData, price: undefined}}   | ${priceInput}         | ${'price has not been entered'}
        ${{...validFormData, amount: undefined}}  | ${amountInput}        | ${'amount has not been entered'}
        ${{...validFormData, pricePerLiter: undefined}}  | ${pricePerLiterInput}  | ${'pricePerLiter has not been entered'}
        ${{...validFormData, tripKilometers: undefined}} | ${tripKilometersInput} | ${'tripKilometers has not been entered'}
    `('should have invalid form when $inputName has not been entered', async ({ formData, input }) => {
        render(AddRefuelling, { props: { vehicleId: 20 }});

        await fillInForm(formData);

        // Form does not submit
        await fireEvent.click(submitButton());
        expect(addRefuellingToVehicle).not.toHaveBeenCalled();

        // Input field is marked invalid
        expect(input().classList).toContain('invalid');
    });
})