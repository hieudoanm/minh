import axios from '@hieudoanm/axios';
import logger from '@hieudoanm/pino';
import { API_KEY_FIXER } from '../../environments';
import { PERSONAL_CURRENCIES, TOP_CURRENCIES } from './forex.constants';
import { ForexResponse } from './forex.types';

export type Category = 'top' | 'personal' | null;

const categoryCurrencies: Record<Exclude<Category, null>, string[]> = {
  top: TOP_CURRENCIES,
  personal: PERSONAL_CURRENCIES,
};

export const getForexRates = async (
  { category = null }: { category: Category } = { category: null }
): Promise<ForexResponse> => {
  try {
    const url = `http://data.fixer.io/api/latest?access_key=${API_KEY_FIXER}`;
    const response = await axios.get<ForexResponse>(url);
    if (category === null) {
      return response;
    }

    const { rates = {} } = response;
    const keys = Object.keys(rates);
    const filterKeys = keys.filter((key) =>
      categoryCurrencies[category].includes(key)
    );
    const filterRates: Record<string, number> = {};
    for (const key of filterKeys) {
      filterRates[key] = rates[key];
    }
    return { ...response, rates: filterRates };
  } catch (error) {
    logger.error(`getForex Error ${error}`);
    throw error;
  }
};
