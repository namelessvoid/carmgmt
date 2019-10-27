import { cleanup, render, fireEvent } from '@testing-library/svelte';

import './i18n/dictionary';
import './i18n/locale';

import DatePicker from './DatePicker.svelte';

describe('DatePicker', () => {
    afterAll(cleanup);

    const dateInput = () => document.querySelector('input[name=date]');
    const timeInput = () => document.querySelector('input[name=time]');

    it('should display value of date provided as prop', () => {
        const date = new Date('2035-05-22T14:30:00+02:00');

        render(DatePicker, { props: { value: date }});

        expect(dateInput().value).toBe('2035-05-22');
        expect(timeInput().value).toBe('14:30');
    });

    it('should set seconds and milliseconds to zero', () => {
        const date = new Date('2035-05-22T14:30:59.932+02:00');

        const { component } = render(DatePicker, { props: { value: date }});

        const actualDate = component.$$.ctx.value;
        expect(actualDate.getSeconds()).toBe(0);
        expect(actualDate.getMilliseconds()).toBe(0);
    })

    it('should trigger changed event when date is changed', async () => {
        // Be aware of DST
        const date = new Date('2035-05-22T14:30:00+02:00');
        const expectedDate = new Date('2000-10-01T14:30:00+02:00');
        const changeCallback = jest.fn();

        const { component } = render(DatePicker, { props: { value: date }});
        component.$on('changed', changeCallback);

        await fireEvent.change(dateInput(), { target: { value: '2000-10-01' }});

        expect(changeCallback).toHaveBeenCalledTimes(1);
        expect(changeCallback.mock.calls[0][0].detail).toEqual(expectedDate);
    });

    it('should trigger changed event when time is changed', async () => {
        const date = new Date('2035-05-22T14:30:00+02:00');
        const expectedDate = new Date('2035-05-22T17:15:00+02:00');
        const changeCallback = jest.fn();

        const { component } = render(DatePicker, { props: { value: date }});
        component.$on('changed', changeCallback);

        await fireEvent.change(timeInput(), { target: { value: '17:15' }});

        expect(changeCallback).toHaveBeenCalledTimes(1);
        expect(changeCallback.mock.calls[0][0].detail).toEqual(expectedDate);
    });

    it('should trigger changed event with "Invalid Date" when invalid date is entered', async () => {
        const changeCallback = jest.fn();

        const { component } = render(DatePicker, { props: { value: new Date() }});
        component.$on('changed', changeCallback);

        await fireEvent.change(dateInput(), { target: { value: 'asdfas' }});

        expect(changeCallback).toHaveBeenCalledTimes(1);
        expect(changeCallback.mock.calls[0][0].detail.toString()).toEqual('Invalid Date');
    });

    it('should trigger changed event with "Invalid Date" when invalid time is entered', async () => {
        const changeCallback = jest.fn();

        const { component } = render(DatePicker, { props: { value: new Date() }});
        component.$on('changed', changeCallback);

        await fireEvent.change(timeInput(), { target: { value: 'asdfas' }});

        expect(changeCallback).toHaveBeenCalledTimes(1);
        expect(changeCallback.mock.calls[0][0].detail.toString()).toEqual('Invalid Date');
    });

    it('should validate input', async () => {
        render(DatePicker, { props: { value: new Date() }});

        expect(dateInput().classList).not.toContain('invalid');
        expect(timeInput().classList).not.toContain('invalid');

        await fireEvent.change(dateInput(), { target: { value: 'asdf' }});
        expect(dateInput().classList).toContain('invalid');
        expect(timeInput().classList).not.toContain('invalid');

        await fireEvent.change(dateInput(), { target: { value: '2000-10-01' }});
        expect(dateInput().classList).not.toContain('invalid');
        expect(timeInput().classList).not.toContain('invalid');

        await fireEvent.change(timeInput(), { target: { value: 'adf' }});
        expect(dateInput().classList).not.toContain('invalid');
        expect(timeInput().classList).toContain('invalid');

        await fireEvent.change(timeInput(), { target: { value: '14:12' }});
        expect(dateInput().classList).not.toContain('invalid');
        expect(timeInput().classList).not.toContain('invalid');

        await fireEvent.change(dateInput(), { target: { value: 'adsf' }});
        await fireEvent.change(timeInput(), { target: { value: 'adf' }});
        expect(dateInput().classList).toContain('invalid');
        expect(timeInput().classList).toContain('invalid');
    });
});