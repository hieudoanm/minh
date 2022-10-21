import axios from '@hieudoanm/axios';
import logger from '@hieudoanm/pino';
import { API_KEY_FIXER } from '../../environments';
import { ForexResponse } from './forex.types';

export type Category = 'top' | 'personal' | null;

const topCurrencies: string[] = [
  'AUD',
  'CAD',
  'CHF',
  'CNH',
  'EUR',
  'GBP',
  'HKD',
  'JPY',
  'NZD',
  'USD',
  'VND',
];

const personalCurrencies: string[] = [
  'AUD',
  'CAD',
  'CNH',
  'EUR',
  'GBP',
  'JPY',
  'KRW',
  'SGD',
  'THB',
  'USD',
  'VND',
];

const categoryCurrencies: Record<Exclude<Category, null>, string[]> = {
  top: topCurrencies,
  personal: personalCurrencies,
};

export const getForex = async (
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
