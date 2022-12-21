import { useAxios } from '@hieudoanm/hooks';
import CircularProgress from '@mui/material/CircularProgress';
import FormControl from '@mui/material/FormControl';
import InputLabel from '@mui/material/InputLabel';
import MenuItem from '@mui/material/MenuItem';
import Select from '@mui/material/Select';
import { useEffect, useState } from 'react';
import { BASE_URL } from '../../../configs';
import { currencies } from '../../../constants';
import { ForexRate, ForexResponse } from '../../../types';
import List from '../../atoms/List';
import Forex from '../../molecules/Forex';

const ForexRates: React.FC = () => {
  const [params, setParams] = useState<{ category: string }>({ category: '' });
  const url = `${BASE_URL}/api/forex/rates`;
  const { data, loading, error, refetch } = useAxios<ForexResponse>(url, {
    params,
  });
  const [eurRates, setRates] = useState<ForexRate[]>([]);

  useEffect(() => {
    const success: boolean = data?.success || false;
    if (data && success) {
      const base: string = data.base;
      const rates: Record<string, number> = data.rates;
      const codes: string[] = Object.keys(rates);
      const newRates: ForexRate[] = codes
        .map((code) => {
          const codeRate = rates[code];
          const baseRate = rates[base];
          const eurRate = rates['EUR'];
          const rate = (baseRate * eurRate) / codeRate;
          const currency: string = currencies[code];
          return { code, currency, rate };
        })
        .sort((a, b) => (b.rate > a.rate ? 1 : -1));
      setRates(newRates);
    } else {
      setRates([]);
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [data]);

  if (loading) {
    return (
      <List>
        <List.Header>
          <h1 className="uppercase text-center text-2xl">Rates</h1>
        </List.Header>
        <List.Item className="flex items-center justify-center">
          <CircularProgress />
        </List.Item>
      </List>
    );
  }

  if (error) {
    const errorMessage: string = error.stack || error.message || 'Error';

    return (
      <List>
        <List.Header>
          <h1 className="uppercase text-center text-2xl">Rates</h1>
        </List.Header>
        <List.Item>
          <p className="p-4 text-center">ERROR: {errorMessage}</p>
        </List.Item>
      </List>
    );
  }

  if (!data) {
    return (
      <List>
        <List.Header>
          <h1 className="text-2xl uppercase">Rates</h1>
        </List.Header>
        <List.Item>
          <p>Response is null</p>
        </List.Item>
      </List>
    );
  }

  return (
    <List>
      <List.Header>
        <div className="flex justify-between items-center">
          <h1 className="uppercase text-center text-2xl">
            Rates ({eurRates.length})
          </h1>
          <div className="flex gap-4">
            <FormControl size="small">
              <InputLabel id="label-category">Category</InputLabel>
              <Select
                labelId="label-category"
                id="category"
                value={params.category}
                label="Category"
                onChange={(event) => {
                  const category: string | null = event.target.value;
                  setParams({ ...params, category });
                  refetch();
                }}
              >
                <MenuItem value="">All</MenuItem>
                <MenuItem value="top">Top</MenuItem>
                <MenuItem value="personal">Personal</MenuItem>
              </Select>
            </FormControl>
          </div>
        </div>
      </List.Header>
      <Forex rates={eurRates} />
    </List>
  );
};

export default ForexRates;
