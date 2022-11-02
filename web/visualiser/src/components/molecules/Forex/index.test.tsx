import { render } from '@testing-library/react';
import Rates from '.';

describe('Rates', () => {
  test('render default', () => {
    const wrapper = render(
      <Rates rates={[{ code: 'ABC', currency: 'abcdef', rate: 1 }]} />
    );
    expect(wrapper.container).toMatchSnapshot();
  });
});
