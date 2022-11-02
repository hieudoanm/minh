import { render } from '@testing-library/react';
import Coins from '.';

describe('Coins', () => {
  test('render default', () => {
    const wrapper = render(<Coins coins={[]} />);
    expect(wrapper.container).toMatchSnapshot();
  });
});
