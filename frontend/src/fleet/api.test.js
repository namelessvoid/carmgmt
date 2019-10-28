import { handleFetch } from './api';

describe('handleFetch()', () => {
    let originalConsoleError;
    beforeEach(() => {
        originalConsoleError = console.error;
        console.error = jest.fn();
    });

    afterEach(() => {
        console.error = originalConsoleError;
    });

    it.each`
        ok       | json                            | error              | reason
        ${false} | ${() => { return 'error.foo' }} | ${'error.foo'}     | ${'response is not ok and body is json'}
        ${false} | ${() => { throw 'foo' }}        | ${'error.unknown'} | ${'response is not ok and body is no json'}
        ${true}  | ${() => { throw 'foo' }}        | ${'error.unknown'} | ${'response is ok and body is no json'}
    `('should throw error if $reason', async ({ ok, json, error }) => {
        const fetchPromise = Promise.resolve({ ok, json });

        expect(handleFetch(fetchPromise)).rejects.toThrowError(error);
    });

    it('should throw error if promise rejects', async () => {
        const fetchPromise = Promise.reject(new Error("Fetch promise rejected"));

        expect(handleFetch(fetchPromise)).rejects.toThrowError('error.networkFailure');
    });

    it('should return response body', async () => {
        const response = {
            ok: true,
            status: 200,
            json: () => { return {key: 'value'} }
        };

        const fetchPromise = Promise.resolve(response);

        const expectedResult = {key: 'value'};
        const expectedError = undefined;

        let actualResult;
        let actualError;

        try {
            actualResult = await handleFetch(fetchPromise);
        } catch(e) {
            actualError = e;
        }

        expect(actualResult).toEqual(expectedResult);
        expect(actualError).toEqual(expectedError);
    });
});