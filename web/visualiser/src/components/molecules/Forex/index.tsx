import { vndFormatter } from '@hieudoanm/utils';
import { ForexRate } from '../../../types';
import List from '../../atoms/List';

const Rates: React.FC<{ rates: ForexRate[] }> = ({ rates }) => {
  return (
    <>
      {rates.map((item) => {
        const { code, currency, rate } = item;
        return (
          <List.Item key={`code-${code}`}>
            <div className="flex justify-between items-center">
              <div>
                <span>
                  <b>{code}</b>
                </span>
                <span> - </span>
                <span className="hidden sm:inline ml-1">
                  {currency || 'N/A'}
                </span>
              </div>
              <div>{vndFormatter(rate)}</div>
            </div>
          </List.Item>
        );
      })}
    </>
  );
};

export default Rates;
