import logger from '@hieudoanm/pino';
import { addZero, eurFormatter } from '@hieudoanm/utils';
import { frankfurterClient } from '../libs/forex';

const PERSONAL_CURRENCIES: string[] = [
  'AUD',
  'CAD',
  'CHF',
  'CNH',
  'EUR',
  'GBP',
  'JPY',
  'KRW',
  'SGD',
  'THB',
  'USD',
];

export const getForexMessage = async (): Promise<string> => {
  const response = await frankfurterClient.getLatest();
  logger.info('getForexMessage() base and rates', { response });
  const { base = '', rates = {} } = response;
  logger.info('getForexMessage() base and rates', { base, rates });
  if (base !== '' || Object.keys(rates).length === 0) return '';

  const codes: string[] = Object.keys(rates);
  const personalCodes = codes.filter((key) =>
    PERSONAL_CURRENCIES.includes(key)
  );
  const filterRates: Record<string, number> = {};
  for (const code of personalCodes) {
    filterRates[code] = rates[code];
  }
  return personalCodes
    .map((code) => {
      const codeRate = rates[code];
      const baseRate = rates[base];
      const eurRate = rates['EUR'];
      const rate = (baseRate * eurRate) / codeRate;
      return { code, rate };
    })
    .sort((a, b) => (b.rate > a.rate ? 1 : -1))
    .map(({ code, rate }, index: number) => {
      return `\`${addZero(index + 1)}. ${code} - ${eurFormatter(rate)}\``;
    })
    .join('\n');
};
