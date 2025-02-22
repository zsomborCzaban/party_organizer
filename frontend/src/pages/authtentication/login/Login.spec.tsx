import {expect, it, test} from 'vitest';
import { render } from '@testing-library/react';
import { Login } from './Login';

test('Login page',()=>{
    it('should render',()=>{
        const component = render(<Login/>);
        expect(component.asFragment()).toMatchSnapshot();
    });
});