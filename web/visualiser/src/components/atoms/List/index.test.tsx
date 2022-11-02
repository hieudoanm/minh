import { render } from '@testing-library/react';
import List from '.';

describe('List', () => {
  test('render default', () => {
    const wrapper = render(<List>List</List>);
    expect(wrapper.container).toMatchSnapshot();
  });
});
