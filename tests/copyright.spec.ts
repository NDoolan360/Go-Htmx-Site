import { describe, expect, test } from 'vitest';
import { replaceWithCurrentYear } from '../src/copyright';

describe('Copyright', () => {
    test('replaceWithCurrentYear', () => {
        let testInput = 'TestString created in 2000';
        const currentYear = new Date().getFullYear().toString();

        testInput = replaceWithCurrentYear(testInput, '2000');

        expect(testInput).toEqual('TestString created in ' + currentYear);
    });
});
