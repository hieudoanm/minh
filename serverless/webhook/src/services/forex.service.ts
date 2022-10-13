import { addZero, vndFormatter } from '@hieudoanm/utils';
import { forexClient } from '../libs/forex';

const PERSONAL_CURRENCIES: string[] = [
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

export const getForexMessage = async (): Promise<string> => {
  const { base, rates } = await forexClient.getLatest();
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
      const vndRate = rates['VND'];
      const rate = (baseRate * vndRate) / codeRate;
      return { code, rate };
    })
    .sort((a, b) => (b.rate > a.rate ? 1 : -1))
    .map(({ code, rate }, index: number) => {
      return `\`${addZero(index + 1)}. ${code} - ${vndFormatter(rate)}\``;
    })
    .join('\n');
};
